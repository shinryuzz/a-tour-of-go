package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int) // make channel

	go sum(s[:len(s)/2], c) // run sum in goroutine
	go sum(s[len(s)/2:], c) // run sum in goroutine
	x, y := <-c, <-c        // receive from c

	fmt.Println(x, y, x+y)
}
