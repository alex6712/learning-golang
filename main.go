package main

import (
	"fmt"
)

func main() {
	numbers := [3]int{1, 2, 3}
	numbersPointer := &numbers

	fmt.Println(numbers)

	firstSlice := numbers[:2]
	secondSlice := numbers[1:]

	fmt.Println(firstSlice, secondSlice)

	numbersPointer[1] = 0

	fmt.Println(firstSlice, secondSlice)
	fmt.Println(numbers)
}
