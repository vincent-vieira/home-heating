package core

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestTemperatureCollector(t *testing.T) {
	assertThat := assert.New(t)
	measureProvider := func() (temperature float32, humidity float32, err error) { return 23, 37, nil }
	measures := make(chan Measure, 1)

	now := time.Now()
	measure(measures, now, measureProvider)

	assertThat.Equal(Measure{Date: now.UnixNano(), Humidity: 37, Temperature: 23}, <-measures)
	close(measures)
}
