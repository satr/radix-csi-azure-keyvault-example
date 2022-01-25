package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var envVars = map[string]bool{"CONNECTION_STRING": true, "DB_USER": true, "DB_PASS": true, "DB_QA_USER": true}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Azure Key vault env-vars</h1>")
	fmt.Fprintf(w, "<ul>")
	for _, envVar := range os.Environ() {
		envVarNameValue := strings.Split(envVar, "=")
		if len(envVarNameValue) < 2 {
			continue
		}
		envVarName := envVarNameValue[0]
		if _, ok := envVars[envVarName]; ok {
			fmt.Fprintf(w, "<li>%s</li>", envVar)
		}
	}
	fmt.Fprintf(w, "</ul>")
}

func main() {
	fmt.Println("Start server")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8082", nil))
	fmt.Println("Stopped server")
}
