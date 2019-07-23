package main

import (
	"fmt"
	"net/http"
)

func linty(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body><img src='/image/Linty6.jpg'></body></html>")
}

func main() {
	http.HandleFunc("/", linty)
	http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("image"))))
	//http.Handle("/image/Linty6.jpg", http.FileServer(http.Dir("image/")))
	http.ListenAndServe(":5000", nil)
}
