package core

import (
	"os/signal"
	"syscall"
	"os"
	"sync"
	"log"
	"os/user"
	"github.com/gerp/dht22"
)

const subroutines = 2;

func Start(listenPort int, measureInterval int, gpioPort int) {
	signals := make(chan os.Signal, 1)
	trap := make(chan bool, subroutines)
	measures := make(chan Measure)
	currentUser, currentUserFetchError := user.Current()
	if currentUserFetchError != nil { log.Panicln("Unable to get current user.", currentUserFetchError) }
	if currentUser.Uid != "0" { log.Panicln("Root is required to read temperature sensor") }

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	defer func() {
		log.Println("Exiting.")
	}()
	defer close(signals)
	defer close(trap)

	var wg sync.WaitGroup
	wg.Add(subroutines)

	log.Println("Starting temperature collector subroutine")
	go temperatureCollector(measureInterval, measures, trap, &wg, func() (float32, float32, error) {
		return dht22.Read(gpioPort)
	})();

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
