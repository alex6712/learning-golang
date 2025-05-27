package main

import (
	"fmt"
)

func main() {
	i := 1
	p := &i

	*p++

	fmt.Printf("p = %x\n", p)
	fmt.Printf("i = %d, *p = %d\n", i, *p)
}
