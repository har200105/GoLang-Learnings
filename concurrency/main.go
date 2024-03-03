package main

import (
	"fmt"
	"sync"
	"time"
)

func runDispatch(name string, done chan bool) {
	fmt.Println(name)
	done <- true
	close(done)
}

func runStore(name string, done chan bool) {
	fmt.Println(name)
	done <- true
}

func main() {
	done := make(chan bool)
	go runDispatch("akipiD", done)
	go runStore("Test", done)
	for range done {
	}
	mainNew()
	mains()

}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(2 * time.Second) // Simulating some work
	fmt.Printf("Worker %d done\n", id)
}

func mainNew() {
	var wg sync.WaitGroup

	numWorkers := 3

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All workers have completed their tasks")
}

var (
	counter int = 5
	mutex   sync.Mutex
)

func incrementCounter(id int, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 0; i < 5; i++ {
		mutex.Lock() // Lock the mutex to protect shared data
		counter += 1 // Modify the shared data
		fmt.Printf("Goroutine %d: Counter is %d\n", id, counter)
		mutex.Unlock() // Unlock the mutex
		time.Sleep(time.Millisecond * 100)
	}
}

func mains() {
	var wg sync.WaitGroup

	numGoroutines := 3

	for i := 1; i <= numGoroutines; i++ {
		wg.Add(1)
		go incrementCounter(i, &wg)
	}

	wg.Wait()

	fmt.Println("Final Counter Value:", counter)
	safeCounterMain()
}
