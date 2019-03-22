package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"golangLeraning/utils"
)

func printRequest(r *http.Request) {
	fmt.Println("r.Form=", r.Form)         // get the Get params
	fmt.Println("r.PostForm=", r.PostForm) // get the Post params
	fmt.Println("path=", r.URL.Path)       // get url location
	fmt.Println("scheme=", r.URL.Scheme)   // scheme
	fmt.Println("method=", r.Method)       // to get the request Method

	if r.Method == "GET" {
		fmt.Println("Http Get params list begin:")
		for k, v := range r.Form {
			fmt.Println("Http Get["+k+"]=", strings.Join(v, " ; "))
		}
		fmt.Println("Http Get params list end:")
	} else if r.Method == "POST" {
		fmt.Println("Http Post params list begin:")
		for k, v := range r.PostForm {
			fmt.Println("Http Post["+k+"]=", strings.Join(v, " ; "))
		}
		fmt.Println("Http Post params list end:")
	}

	query := r.Form["name"]
	fmt.Println("r.Form['name']=", query)
	if len(query) > 0 {
		fmt.Println("r.Form['name'][0]=", query[0])
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	// Explicit parse params, default not
	r.ParseForm()
	fmt.Println("hello world")
	printRequest(r)

	// w write to browser/requester/client
	utils.Fprintf(w, "<h1>Hello World!</h1>")
}

func showMore(w http.ResponseWriter, r *http.Request) {
	// Explicit parse params, default not
	r.ParseForm()

	fmt.Println("More method")
	printRequest(r)

	utils.Fprintf(w, "<h1>Hello More!</h1>")
}

func showMoreOther(w http.ResponseWriter, r *http.Request) {
	// Explicit parse params, default not
	err := r.ParseForm()
	fmt.Println("More other method", err)
	printRequest(r)

	utils.Fprintf(w, "<h1>Hello More1!</h1>");
}

func main() {
	// setting the request path, like 'http://127.0.0.1:9090' + pattern
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/more", showMore)
	http.HandleFunc("/more/", showMoreOther)
	// listening the port: 9090
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("Error happened ListenAndServe: ", err)
	}
}
