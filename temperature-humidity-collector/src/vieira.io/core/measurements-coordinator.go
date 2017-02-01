package core

import (
	"os/signal"
	"syscall"
	"os"
	"sync"
	"log"
	"github.com/d2r2/go-dht"
)

const subroutines = 1;

func Start(listenPort int, measureInterval int, gpioPort int) {
	signals := make(chan os.Signal, 1)
	trap := make(chan bool, subroutines)
	measures := make(chan Measure)

	measureProvider := func() (temperature float32, humidity float32, retried int, err error) { return dht.ReadDHTxxWithRetry(dht.AM2302, gpioPort, true, 10) }

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	defer func() {
		log.Println("Exiting.")
	}()
	defer close(signals)
	defer close(trap)

	var wg sync.WaitGroup
	wg.Add(subroutines)

	log.Println("Starting temperature collector subroutine")
	go temperatureCollector(measureInterval, measures, trap, &wg, measureProvider)();

	log.Println("Starting websocket subroutine")
	go webSocketNotifier(listenPort, measures, trap, &wg)();

	log.Println("Starting trap interception subroutine")
	go func() {
		log.Println(<-signals)
		log.Println("Terminating all subroutines...")
		for i := 0; i < subroutines; i++ {
			trap <- true
		}
	}()

	log.Println("Now waiting for interrupt.")
	wg.Wait()
}
