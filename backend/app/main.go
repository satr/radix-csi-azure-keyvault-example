package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Output all env-vars</h1>")
	fmt.Fprintf(w, "<ul>")
	for _, envVar := range os.Environ() {
		fmt.Fprintf(w, "<li>%s</li>", envVar)
	}
	fmt.Fprintf(w, "</ul>")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}