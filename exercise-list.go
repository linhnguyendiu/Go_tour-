package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T comparable] struct {
	next *List[T]
	val  T
}

func (l *List[T]) AddFront(value T) *List[T] {
	newNode := &List[T]{val: value, next: l}
	return newNode
}

func (l *List[T]) Remove(value T) *List[T] {
	if l != nil && l.val == value {
		return l.next
	}
	current := l
	for current != nil && current.next != nil && current.next.val != value {
		current = current.next
	}
	if current != nil && current.next != nil {
		current.next = current.next.next
	}
	return l
}

func (l *List[T]) Print() {
	current := l
	for current != nil {
		fmt.Printf("%v ", current.val)
		current = current.next
	}
	fmt.Println()
}

func main() {
	myList := &List[int]{val: 3}
	myList = myList.AddFront(2)
	myList = myList.AddFront(1)
	fmt.Println("Integer List:")
	myList.Print()

	myList = myList.Remove(2)
	fmt.Println("Integer List After Remove 2:")
	myList.Print()
}
