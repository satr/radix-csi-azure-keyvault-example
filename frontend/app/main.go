package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var envVars = map[string]bool{"CONNECTION-STRING": true, "DB_USER": true, "DB_PASS": true, "DB_QA_USER": true}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Azure Key vault env-vars</h1>")
	fmt.Fprintf(w, "<ul>")
	for _, envVar := range os.Environ() {
		if _, ok := envVars[envVar]; ok {
			fmt.Fprintf(w, "<li>%s</li>", envVar)
		}
	}
	fmt.Fprintf(w, "</ul>")
}

func main() {
	fmt.Println("Start server")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("Stopped server")
}
