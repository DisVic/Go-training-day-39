package main

import (
	"fmt"
)

func main() {
	ch1, ch2 := make(chan int, 1), make(chan int, 1)
	stop := make(chan struct{}, 0)
	ch1 <- 3
	ch2 <- 3
	r := calculator(ch1, ch2, stop)

	close(stop)
	fmt.Println(<-r)
}

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		var x int
		select {
		case x = <-firstChan:
			output <- x * x
		case x = <-secondChan:
			output <- x * 3
		case <-stopChan:
			return
		default:
		}
	}()
	return output
}
