package main

import "fmt"

type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	head := &List[int]{val: 0}
	tail := head

	for i := 1; i < 10; i++ {
		new := &List[int]{val: i}

		tail.next = new
		tail = new
	}

	for current := head; current != nil; current = current.next {
		fmt.Printf("New node found: val=%v\n", current.val)
	}
}
