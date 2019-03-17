package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "<h1>Hello world, your name is:%s</h1>", request.FormValue("name"))
		//fmt.Fprintf(writer, "<h1>Hello world</h1>")
	})

	http.ListenAndServe(":8888", nil)

}
