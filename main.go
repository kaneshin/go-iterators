package main

import (
	"fmt"
	"github.com/kaneshin/go-iterators/iterators"
)

func main() {
	var arr = []int{1, 2, 3}

	var iter = iterators.NewIterator(arr)
	fmt.Println(iter)

	var arr2 = iter.MapInt(func(i int) int {
		return i * 2
	})
	fmt.Println(arr2)
}
