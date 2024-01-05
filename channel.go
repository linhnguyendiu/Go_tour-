package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{2, 3, -4, 2, 4}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

	//Buffered Channels
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//ch <- 3 deadlock
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
