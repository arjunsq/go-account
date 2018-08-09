package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/sayhello", sayhelloName) // setting router rule
	http.HandleFunc("/login", login)
	http.ListenAndServe(":9000", nil)
}
