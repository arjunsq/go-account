package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/index", index)
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir("./"))))
	http.ListenAndServe(":9000", nil)
}
