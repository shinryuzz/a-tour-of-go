package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		// blocks until one of its cases can run, then it executes that case.
		// If more than one of the cases can run, it chooses one at random.
		select {
		case c <- x: // c keep the value untile it is received at fmt.Println(<-c)
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int) // channel to deliver the sequence of Fibonacci numbers
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) // receive from channel c
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
