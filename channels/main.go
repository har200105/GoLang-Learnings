package main

import (
	"fmt"
	"sync"
)

func runFirst(ch1 chan string) {
	fmt.Println("akipiD")
	ch1 <- "akipiD"
}

func runSecond(ch2 chan string) {
	fmt.Println("Safarnama")
	ch2 <- "Safarnama"
}

func main() {
	fmt.Println("Trying Channels.")

	wg := &sync.WaitGroup{}

	defer wg.Wait()

	wg.Add(2)

	channel := make(chan int, 1)

	go func(ch <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		// value, isChannelOpen := <-ch
		// fmt.Println(isChannelOpen)
		// fmt.Println(value)

		for x := range ch {
			fmt.Println(x)
		}

	}(channel, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()
		ch <- 1000
		ch <- 1001
		close(ch)
	}(channel, wg)

	strCh1 := make(chan string)
	strCh2 := make(chan string)

	go runFirst(strCh1)
	go runSecond(strCh2)

	select {
	case fromCh1 := <-strCh1:
		fmt.Println("1st", fromCh1)
	case fromCh2 := <-strCh2:
		fmt.Println("2nd ", fromCh2)
	}

}
