package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Now Server localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, this is a web page! Cool huh? %s", time.Now())
}
