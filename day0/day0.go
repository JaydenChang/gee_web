package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "URL.Path=%q\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	// http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
