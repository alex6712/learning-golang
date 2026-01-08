package main

import "fmt"

func fibonacci() func() int {
	prev, cur := 1, 0

	return func() int {
		prev, cur = cur, prev+cur

		return cur
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
