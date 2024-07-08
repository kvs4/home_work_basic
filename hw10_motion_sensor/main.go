package main

import (
	"fmt"
	"time"
)

func main() {
	dataMotionSensorChan := make(chan int)
	processedDataChan := make(chan float64)
	go SensorDataGenerator(dataMotionSensorChan)
	go DataProcessor(dataMotionSensorChan, processedDataChan)

	for processedData := range processedDataChan {
		fmt.Println("Average of 10 values:", processedData)
	}
}

func SensorDataGenerator(c chan<- int) {
	defer close(c)

	timer := time.NewTimer(time.Second * 60)
	data := 1

outerfor:
	for {
		select {
		case <-timer.C:
			fmt.Println("Attention: timeout of operation data reader from motion sensor!")
			break outerfor
		default:
			c <- data
			// fmt.Println("Send to chan value", data)
			data++
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func DataProcessor(chanin <-chan int, chanout chan<- float64) {
	defer close(chanout)

	var summ int
	i := 1

	// fmt.Println("DataProcessor is started")

	for elem := range chanin {
		summ += elem
		// fmt.Println("DataProcessor read value", elem)
		if i == 10 {
			chanout <- float64(summ) / float64(i)
			i = 1
			summ = 0
		} else {
			i++
		}
	}
}
