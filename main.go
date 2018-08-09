package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/", login)
	http.ListenAndServe(":9000", nil)
}
