package nodes

import (
	"bufio"
	"net"
)

const PROTOCOL = "tcp"

func NetworkSink(addr string, in <- chan int) {
	listener, err := net.Listen(PROTOCOL, addr)
	if err != nil {
		panic(err)
	}
	go func() {
		defer listener.Close()
		// maybe a loop accept
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		defer conn.Close()
		writer := bufio.NewWriter(conn)
		defer writer.Flush()
		WriteSink(writer, in)
	}()
}

func NetworkSource(addr string) <-chan int {
	out := make(chan int)
	go func() {
		conn, err := net.Dial(PROTOCOL, addr)
		if err != nil {
			panic(err)
		}
		r := ReaderSource(bufio.NewReader(conn), -1)
		for v := range r {
			out <- v
		}
		close(out)
	}()
	return out
}
