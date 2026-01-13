package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk проходит по бинарному отсортированному
// дереву t и пушит значения поэлементно в канал ch
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same проверяет, совпадают ли два отсортированных
// бинарных дерева t1 и t2
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	// запускаем проход по деревьям в отдельных потоках
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for range 10 {
		// main поток блокируется, пока не получит
		// по элементу из каждого канала, потом их сравнивает
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(3), tree.New(3)))
}
