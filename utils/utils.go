package utils

import (
	"fmt"
	"io"
	"log"
)
func Fprintf(w io.Writer, format string, a ...interface{})  {
	_, err := fmt.Fprintf(w, format)
	if err != nil {
		log.Fatal("log error happened:", err)
	}

}