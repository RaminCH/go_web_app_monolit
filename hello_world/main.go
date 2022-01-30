package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	})

	_ = http.ListenAndServe(":8080", nil)
}

//In terminal
// ramie@ramie:~/go/src/Web_Part1/hello_world$ go run main.go
// Number of bytes written: 12
// Number of bytes written: 12

//In browsser
// http://localhost:8080/
// Hello World!
// Ctrl+C - cancel
