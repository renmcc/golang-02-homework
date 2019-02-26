package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello %s", r.URL.Path)
}

func user_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "user %s", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/user/", user_handler)
	http.ListenAndServe(":8080", nil)
}
