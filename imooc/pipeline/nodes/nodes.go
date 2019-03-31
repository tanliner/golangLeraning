package nodes

import (
	"encoding/binary"
	"io"
	"math/rand"
	"sort"
)

/**
 * send array data
 */
func ArraySource(a ...int) <- chan int {
	out := make(chan int)
	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

/**
 * read data from in and send to out
 */
func InMemSort(in <- chan int) <- chan int {
	out := make(chan int)
	go func() {
		// Read into memory
		a := []int {}
		for v := range in {
			a = append(a, v)
		}
		// Sort
		sort.Ints(a)

		// Output
		for _, v := range a {
			out <- v
		}

		close(out)

	}()
	return out
}

func Merge(in1, in2 <- chan int) <- chan int  {
	out := make(chan int)
	go func() {
		v1, ok1 := <- in1
		v2, ok2 := <- in2
		for ok1 || ok2  {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <- in1
			} else {
				out <- v2
				v2, ok2 = <- in2
			}
		}
		close(out)
	}()
	return out
}

// to check the external sort Demo
func ReaderSource2(reader io.Reader) <- chan int {
	out := make(chan int)
	go func() {
		buffer := make([]byte, 8)
		for {
			n, error := reader.Read(buffer)
			if n > 0 {
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
			if error != nil {
				break
			}
		}
		close(out)
	}()
	return out
}

func ReaderSource(reader io.Reader, chunkSize int) <- chan int {
	out := make(chan int)
	go func() {
		buffer := make([]byte, 8)
		bytesRead := 0
		for {
			n, error := reader.Read(buffer)
			bytesRead += n
			if n > 0 {
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
			if error != nil ||
				(chunkSize != -1 && bytesRead >= chunkSize) {
				break
			}
		}
		close(out)
	}()
	return out
}

func WriteSink(writer io.Writer, in <- chan int) {
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))
		writer.Write(buffer)
	}
}

func RandomSource(count int) <- chan int {
	out := make(chan int)
	go func() {
		for i:= 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}

// merge external sort
func MergeN(inputs ...<-chan int) <- chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}
	mid := len(inputs) / 2
	// merge inputs[0..m) and inputs[m..end)
	return Merge(MergeN(inputs[:mid]...), MergeN(inputs[mid:]...))
}