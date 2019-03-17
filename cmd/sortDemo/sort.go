package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int {3, 5, 6, 2, 1, 0, -1}
	// call library func
	sort.Ints(array)

	for i, v := range array {
		fmt.Println(i, v)
	}
}
