package core

import (
	"sync"
	"time"
	"log"
)

type MeasureProvider func() (temperature float32, humidity float32, retried int, err error)

func temperatureCollector(measureInterval int, measures chan<- Measure, trap <-chan bool, wg *sync.WaitGroup, measureProvider MeasureProvider) func() {
	return func() {
		measureTicker := time.NewTicker(time.Second * time.Duration(measureInterval))
		select {
		case <-trap:
			log.Println("Stopping temperature measurements routine...")
			measureTicker.Stop()
			wg.Done()
			log.Println("Temperature measurements routine stopped.")
		default:
			for range measureTicker.C {
				measure(measures, time.Now(), measureProvider)
			}
		}
	}
}

func measure(measures chan<- Measure, measureTime time.Time, measureProvider MeasureProvider) {
	temperature, humidity, retried, err := measureProvider()
	if err != nil {
		log.Panicln("Error while reading temperature sensor", err)
	}
	log.Printf("Temperature = %v°C, Humidity = %v%% (retried %d times)\n", temperature, humidity, retried)
	measures <- Measure{Date: measureTime.UnixNano(), Humidity: humidity, Temperature: temperature}
}
