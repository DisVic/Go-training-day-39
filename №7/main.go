package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
)

func main() {
	squares := make([]int, 0)
	workerPool(&squares)
	fmt.Println(squares)
}

func workerPool(s *[]int) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := new(sync.WaitGroup)
	numbersToMultiply, numbersAfterMultiplying := make(chan int), make(chan int)

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, numbersToMultiply, numbersAfterMultiplying)
		}()
	}

	go func() {
		defer close(numbersToMultiply)
		for i := 0; i < 1000; i++ {
			numbersToMultiply <- i
		}
	}()

	go func() {
		wg.Wait()
		close(numbersAfterMultiplying)
	}()
	for result := range numbersAfterMultiplying {
		*s = append(*s, result)
	}
}

func worker(ctx context.Context, numbersToMultiply <-chan int, numbersAfterMultiplying chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return
		case value, ok := <-numbersToMultiply:
			if !ok {
				return
			}
			numbersAfterMultiplying <- square(value)
		}
	}
}

func square(number int) int {
	return number * number
}
