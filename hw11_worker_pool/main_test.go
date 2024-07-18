package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartWorkers(t *testing.T) {
	wantCounter := 100
	counter := 0
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	StartWorkers(wantCounter, &counter, &mu, &wg)
	wg.Wait()
	assert.Equal(t, wantCounter, counter)
}

func TestStartWorkers1(t *testing.T) {
	wantCounter := 1
	counter := 0
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	StartWorkers(wantCounter, &counter, &mu, &wg)
	wg.Wait()
	assert.Equal(t, wantCounter, counter)
}

func TestStartWorkers0(t *testing.T) {
	wantCounter := 0
	counter := 0
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	StartWorkers(wantCounter, &counter, &mu, &wg)
	wg.Wait()
	assert.Equal(t, wantCounter, counter)
}

func TestWorker(t *testing.T) {
	wantCounter := 1
	counter := 0
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	wg.Add(1)
	go Worker(1, &counter, &mu, &wg)
	wg.Wait()
	assert.Equal(t, wantCounter, counter)
}
