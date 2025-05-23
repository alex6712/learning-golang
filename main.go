package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for range 1_000_000_000 {
	}
	elapsed := time.Since(start).Seconds()

	fmt.Printf("Миллиард итераций: %.2f секунд\n", elapsed)
}
