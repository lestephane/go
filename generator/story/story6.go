// story6.go

// +build ignore

package main

import (
	"fmt"
	"time"
)

func xrange(begin, end int) (<-chan int, chan<- struct{}) {
	chOut := make(chan int)
	chStop := make(chan struct{})
	go func() {
		fmt.Println("Goroutine Started")
		defer fmt.Println("Goroutine Terminated")
		defer close(chOut)
		for i := begin; i < end; i++ {
			select {
			case chOut <- i:
				// Nothing to be done
			case <-chStop:
				break
			}
		}
	}()
	return chOut, chStop
}

func main() {
	chOut, chStop := xrange(10, 20)
	for x := range chOut {
		fmt.Println(x)
	}
	close(chStop)
	time.Sleep(1 * time.Second)
}
