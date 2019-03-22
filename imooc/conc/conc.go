package main

import (
	"fmt"
)

func main() {
	// go starts a ro routine
	//for i := 0; i < 500000; i++ {
	//	go printHelloWorld(i)
	//}

	ch := make(chan string)
	for i := 0; i < 50000; i++ {
		go printHelloWorldChan(i, ch)
	}
	for {
		msg := <- ch
		fmt.Println(msg)
	}
}

/**
 *
 */
func printHelloWorld(i int) {
	for {
		fmt.Printf("Hello world %d from goroutine\n", i)
	}
}

func printHelloWorldChan(i int, ch chan string) {
	for {
		ch <- fmt.Sprintf("Hello world from goroutine %d\n", i)
	}
}