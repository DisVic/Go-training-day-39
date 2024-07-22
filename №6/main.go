package main

func main() {

}

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			x1, x2 := <-in1, <-in2
			fn(x1)
			fn(x2)
			out <- x1 + x2
		}
	}()
}
