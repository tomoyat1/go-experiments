package main

import (
	"fmt"
	"sync"
)

func main() {
	lis1 := make(chan int)
	lis2 := make(chan int)
	send := make(chan int)
	listeners := map[chan<- int]struct{}{
		lis1: {},
		lis2: {},
	}

	wg := sync.WaitGroup{}
	go func() {
		wg.Add(1)
		broker(listeners, send)
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		for i := range lis1 {
			fmt.Printf("func1: %d\n", i)
		}
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		for i := range lis2 {
			fmt.Printf("func2: %d\n", i)
		}
		wg.Done()
	}()

	send <- 42
	send <- 43
	send <- 44
	close(send)
	wg.Wait()
}

func broker(listeners map[chan<- int]struct{}, sender <-chan int) {
	for {
		select {
		case i, ok := <-sender:
			if !ok {
				for ch := range listeners {
					// Must close each listener when broker sender is closed, so that deadlocks don't occur.
					close(ch)
				}
				return
			}
			for ch := range listeners {
				ch <- i
			}
		}
	}
}