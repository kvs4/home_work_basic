package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSensorDataGenerator(t *testing.T) {
	dataMotionSensorChan := make(chan int)
	go SensorDataGenerator(dataMotionSensorChan)

	timer := time.NewTimer(time.Second * 61)

	for {
		select {
		case <-timer.C:
			t.Fatal("Test timed out!")
		case _, ok := <-dataMotionSensorChan:
			if !ok {
				return
			}
		}
	}
}

func TestDataProcessor(t *testing.T) {
	dataMotionSensorChan := make(chan int)
	processedDataChan := make(chan float64)
	go DataProcessor(dataMotionSensorChan, processedDataChan)

	wantDataSetAvg := []float64{45.0, 55.0}
	dataSet := []int{40, 50, 40, 50, 40, 50, 40, 50, 40, 50, 50, 60, 50, 60, 50, 60, 50, 60, 50, 60}

	go func() {
		defer close(dataMotionSensorChan)
		for _, data := range dataSet {
			dataMotionSensorChan <- data
		}
	}()

	timer := time.NewTimer(time.Second * 1)

	for _, wantData := range wantDataSetAvg {
		select {
		case gotData := <-processedDataChan:
			assert.Equal(t, wantData, gotData)
		case <-timer.C:
			t.Fatal("Test timed out!")
		}
	}
}
