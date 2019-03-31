package main

import (
	"bufio"
	"fmt"
	"os"

	"golangLeraning/imooc/pipeline/nodes"
)

func main() {
	// mergeDemo()
	generateFile("small.in", 64)
	// generateFile("large.in", 100000000)
}

func mergeDemo() {
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

/**
 * bufio.NewWriter for file, to increase the write speed
 * Note: to use Flush method to flush the data in the buffer
 */
func generateFile(fileName string, n int) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	// always close file, like Java finally block
	defer file.Close()
	p := nodes.RandomSource(n)

	writer := bufio.NewWriter(file)
	nodes.WriteSink(writer, p)
	writer.Flush()

	// file.Close()

	file, err = os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// p = nodes.ReaderSource(bufio.NewReader(file))
	p = nodes.ReaderSource(bufio.NewReader(file), -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count > 100 {
			break
		}
	}
}
