package main

import (
	"fmt"
	"strings"
)

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	//Slices are like references to arrays
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	//Slice literals
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	h := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(h)

	//Slice defaults
	s1 := []int{2, 3, 5, 7, 11, 13}

	s1 = s1[1:4]
	fmt.Println(s1)

	s1 = s1[:2]
	fmt.Println(s1)

	s1 = s1[1:]
	fmt.Println(s1)

	//Slice length and capacity
	s2 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s2)

	// Slice the slice to give it zero length.
	s2 = s2[:0]
	printSlice(s2)

	// Extend its length.
	s2 = s2[:4]
	printSlice(s2)

	// Drop its first two values.
	s2 = s2[2:]
	printSlice(s2)

	// Nil slices
	var s3 []int
	fmt.Println(s3, len(s3), cap(s3))
	if s3 == nil {
		fmt.Println("nil!")
	}

	// Creating a slice with make
	e := make([]int, 5)
	printSlice(e)

	f := make([]int, 0, 5)
	printSlice(f)

	g := f[:2]
	printSlice(g)

	d := g[2:5]
	printSlice(d)

	// Slices of slices
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// Appending to a slice
	var s4 []int
	printSlice(s4)

	// append works on nil slices.
	s4 = append(s4, 0)
	printSlice(s4)

	// The slice grows as needed.
	s4 = append(s4, 1)
	printSlice(s4)

	// We can add more than one element at a time.
	s4 = append(s4, 2, 3, 4)
	printSlice(s4)

	// Range
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// Range continued
	pow1 := make([]int, 10)
	for i := range pow1 {
		pow1[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow1 {
		fmt.Printf("%d\n", value)
	}
}
