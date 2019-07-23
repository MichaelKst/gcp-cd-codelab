package main

import (
	"fmt"
	"net/http"
)

func linty(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body><img src='http://localhost:5000/Linty6.jpg'></body></html>")
}

func main() {
	http.HandleFunc("/", linty)
	http.Handle("/Linty6.jpg", http.FileServer(http.Dir("./image/Linty6.jpg")))
	http.ListenAndServe(":5000", nil)
}
