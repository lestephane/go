// story7.go

// +build ignore

package main

import (
	"fmt"
	"time"
)

func xrange(begin, end int) func() (<-chan int, chan<- struct{}) {
	return func() (<-chan int, chan<- struct{}) {
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
}

func main() {
	f := xrange(10, 20)
	chOut, chStop := f()
	for x := range chOut {
		if x%8 == 0 {
			break
		}
		fmt.Println(x)
	}
	close(chStop)
	time.Sleep(1 * time.Second)
}
