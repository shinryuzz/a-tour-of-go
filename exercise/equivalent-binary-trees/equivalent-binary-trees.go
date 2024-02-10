package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		// when t is nil, it means the end of the tree
		// but it is called recursively
		return
	}
	// search the left tree first, then the root, then the right tree
	Walk(t.Left, ch)
	// when the left tree is done, then the root is sent to the channel
	ch <- t.Value
	// when the root is sent to the channel, then the right tree is searched
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if v1 != v2 || ok1 != ok2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

/*
type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}
*/

func main() {
	// Test Walk
	ch := make(chan int)
	go func() {
		Walk(tree.New(1), ch)
		close(ch)
	}()
	for v := range ch {
		fmt.Println(v)
	}

	// Test Same
	fmt.Println(Same(tree.New(1), tree.New(1))) // true
	fmt.Println(Same(tree.New(1), tree.New(2))) // false

}
