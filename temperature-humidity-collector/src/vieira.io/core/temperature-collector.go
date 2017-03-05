package core

import (
	"sync"
	"time"
	"log"
)

type MeasureProvider func() (float32, float32, error)

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
	temperature, humidity, err := measureProvider()
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Temperature = %v°C, Humidity = %v%%\n", temperature, humidity)
		measures <- Measure{Date: measureTime.UnixNano(), Humidity: humidity, Temperature: temperature}
	}
}
