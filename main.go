package main

import "fmt"

func fibonacci() func() int {
	prev, cur := 0, 0

	return func() int {
		prev, cur = cur, prev+cur

		if cur == 0 {
			cur = 1
		}

		return prev
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
