package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	StartWorkers(100, &counter, &mu, &wg)
	wg.Wait()
	fmt.Printf("Counter = %v\n", counter)
}

func StartWorkers(numberOfWorkers int, counter *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	for i := 1; i <= numberOfWorkers; i++ {
		wg.Add(1)
		go Worker(i, counter, mu, wg)
	}
}

func Worker(id int, counter *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	mu.Lock()
	*counter++
	mu.Unlock()
	wg.Done()
	fmt.Printf("Gorutine %v is done \n", id)
}
