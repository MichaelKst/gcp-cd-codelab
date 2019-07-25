package main

import (
	"fmt"
	"net/http"
)

func linty(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body><img src='https://www.aboriva.com/img/couv/34/nm_sugar-num178-juillet-aout-2016_1154.jpg'><img src='/Linty6.jpg'></body></html>")
}

func main() {
	http.HandleFunc("/", linty)
	http.Handle("/Linty6.jpg/", http.StripPrefix("/Linty6.jpg/", http.FileServer(http.Dir("http://localhost:5000/go/bin/gcp-cd-codelab")))))
	//http.Handle("/image/Linty6.jpg", http.FileServer(http.Dir("image/")))
	//http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":5000", nil)
}
