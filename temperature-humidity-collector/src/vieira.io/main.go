package main

import (
	"flag"
	collector "vieira.io/core"
)

func main() {
	listenPortPtr := flag.Int("listen", 5000, "The listen port for incoming WebSocket connections")
	measureInterval := flag.Int("refresh", 60, "The time between two measures, in seconds")
	gpioPort := flag.Int("gpio", 4, "The GPIO pin number where the sensor data cable is plugged")
	flag.Parse()
	collector.Start(*listenPortPtr, *measureInterval, *gpioPort)
}