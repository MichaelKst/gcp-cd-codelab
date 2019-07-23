package main

import (
	"io"
	"net/http"
)

func linty(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<img src='image/Linty6.jpg'>")
}

func main() {
	http.HandleFunc("/", linty)
	http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("./image"))))
	http.ListenAndServe(":5000", nil)
}
