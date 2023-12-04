package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static/"))

	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("Server running at localhost:4040")
	if err := http.ListenAndServe(":4040", nil); err != nil {
		log.Fatal(err)
	}
}
