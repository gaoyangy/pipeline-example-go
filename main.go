package main

import (
	"fmt"
	"net/http"
	"log"
)

const webContent = "Hello World!"

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/a", helloHandler)
	http.HandleFunc("/b", helloHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, webContent)
}
func helloHandler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, webContent)
}
func helloHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, webContent)
}
