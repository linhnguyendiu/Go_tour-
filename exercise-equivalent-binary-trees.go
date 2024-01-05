package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t != nil {
			walker(t.Left)
			ch <- t.Value
			walker(t.Right)
		}
	}
	walker(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for v1 := range ch1 {
		v2, ok := <-ch2
		if !ok || v1 != v2 {
			return false
		}
	}

	_, ok := <-ch2
	return !ok
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println(Same(tree.New(1), tree.New(1))) // Should return true
	fmt.Println(Same(tree.New(1), tree.New(2))) // Should return false
}
