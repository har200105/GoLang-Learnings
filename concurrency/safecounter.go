package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu   sync.Mutex
	data map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	c.data[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) getValue(key string) int {
	defer c.mu.Unlock()
	c.mu.Lock()
	return c.data[key]
}

func safeCounterMain() {
	c := SafeCounter{data: make(map[string]int)}
	for i := 0; i < 100; i++ {
		go c.Inc("akipiD")
	}
	time.Sleep(time.Second * 2)
	fmt.Println(c.getValue("akipiD"))
}
