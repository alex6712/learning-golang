package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)

	rv := make(map[string]int)
	for _, value := range fields {
		rv[value]++
	}

	return rv
}

func main() {
	wc.Test(WordCount)
}
