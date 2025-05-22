package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for x := 0; x < 1_000_000_000; x++ {
	}
	elapsed := time.Since(start).Seconds()

	fmt.Printf("Миллиард итераций: %.2f секунд\n", elapsed)
}
