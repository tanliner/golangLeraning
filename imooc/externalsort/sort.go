package main

import (
	"bufio"
	"fmt"
	"os"

	"golangLeraning/imooc/pipeline/nodes"
)

func main() {
	sort()
}

func sort()  {
	fileName := "small.in"
	fileNameOut := "small.out"
	fileSize := 512
	chunkCount := 4
	p := createPipeLine(fileName, fileSize, chunkCount)
	writeToFile(p, fileNameOut)
	printFile(fileNameOut)
}

func createPipeLine(
	fileName string,
	fileSize, chunkCount int) <- chan int {
	chunkSize := fileSize / chunkCount
	// sortResults := [] <- chan int{}
	var sortResults [] <-chan int

	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i * chunkSize), 0)
		source := nodes.ReaderSource(bufio.NewReader(file), chunkSize)
		sortResults = append(sortResults, nodes.InMemSort(source))
	}
	return nodes.MergeN(sortResults...)
}

func writeToFile(p <- chan int, fileOut string) {
	file, err := os.Create(fileOut)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	nodes.WriteSink(writer, p)
}

func printFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p := nodes.ReaderSource(file, -1)
	for v := range p {
		fmt.Println(v)
	}
}

