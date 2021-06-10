package main

import (
	"fmt"
	"net/http"
)

func myHandler (w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "<h1>Bunty</h1>")
}

func main() {
	http.HandleFunc("/", myHandler)
	http.ListenAndServe(":8080", nil)
}