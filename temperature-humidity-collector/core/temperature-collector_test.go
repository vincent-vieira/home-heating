package core

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestTemperatureCollector(t *testing.T) {
	assert := assert.New(t)
	measureProvider := struct { MeasureProvider }{}
	measureProvider.measure = func() (temperature float32, humidity float32, retried int, err error) { return 23, 37, 1, nil }
	measures := make(chan Measure, 1)

	now := time.Now()
	measure(measures, now, measureProvider)

	assert.Equal(Measure{date: now.UnixNano(), humidity: 37, temperature: 23}, <-measures)
	close(measures)
}
