package main

import (
	"fmt"
	"sync"
)

func deadlock() {

	var score = []int{0}

	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	wg.Add(2)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		mutex.Lock()
		score = append(score, 10)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		mutex.Lock()
		score = append(score, 7)
		mutex.Unlock()
		wg.Done()
	}(wg, mutex)

	wg.Wait()
	fmt.Println(score)
}
