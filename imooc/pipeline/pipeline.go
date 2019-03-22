package main

import (
	"fmt"

	"./nodes"
)

func main() {
	// test mem sort
	p := nodes.InMemSort(
		nodes.ArraySource(3, 5, 6, 2, 1, 0, -1))

	p2 := nodes.Merge(
		nodes.InMemSort(
			nodes.ArraySource(50, 11, 45, 23, 90, -5)),
		nodes.InMemSort(
			nodes.ArraySource(3, 5, 6, 2, 1, 0, -1)))

	// normal style
	// for {
	// 	// the ok is a section for break
	// 	if num, ok := <- p; ok {
	// 		fmt.Println(num)
	// 	} else {
	// 		break
	// 	}
	// }

	// Use this style, you must make sure the source will close.
	for v := range p {
		fmt.Println(v)
	}

	fmt.Println("\nthe merge test")
	for v := range p2 {
		fmt.Println(v)
	}

}
