package main

import (
	"fmt"
	"net/http"
)

func linty(w http.ResponseWriter, r *http.Request) {
	// Print the HTML content of the web page to be displayed
	fmt.Fprintf(w,"<!DOCTYPE html><html lang='en'><head><meta charset='utf-8'><title>Linty Services</title><meta charset='utf-8'><meta name='viewport' content='width=device-width, initial-scale=1'><link rel='stylesheet' href='https://maxcdn.bootstrapcdn.com/bootstrap/4.2.1/css/bootstrap.min.css'><script src='https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js'></script><script src='https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.6/umd/popper.min.js'></script><script src='https://maxcdn.bootstrapcdn.com/bootstrap/4.2.1/js/bootstrap.min.js'></script><style>.image img {width:200px;display: block;margin-left: auto;margin-right: auto;}</style><body style='background-color:#e4c0f4;'><div class='image'><img src='image/Linty8.png'></div></body></html>")
}

func main() {
	http.HandleFunc("/", linty)
	// Specifies a static directory to locate resources
	http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("/image"))))
	http.ListenAndServe(":5000", nil)
}
