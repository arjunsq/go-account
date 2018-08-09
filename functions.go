package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		fmt.Println("Inside signup")
		t, _ := template.ParseFiles("signup.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		responsevalue := "Username: " + r.Form["username"][0] + " Password: " + r.Form["password"][0] + " registered successfully"
		fmt.Fprintf(w, responsevalue) // write data to response
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		fmt.Println("Inside login")
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		responsevalue := "Username: " + r.Form["username"][0] + " Password: " + r.Form["password"][0] + " logged in successfully"
		fmt.Fprintf(w, responsevalue) // write data to response
	}
}
