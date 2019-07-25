package main

import (
	"fmt"
	"net/http"
)

func linty(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"<!DOCTYPE html><html lang='en'><head><meta charset='utf-8'><title>Linty Services</title><meta charset='utf-8'><meta name='viewport' content='width=device-width, initial-scale=1'><link rel='stylesheet' href='https://maxcdn.bootstrapcdn.com/bootstrap/4.2.1/css/bootstrap.min.css'><script src='https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js'></script><script src='https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.6/umd/popper.min.js'></script><script src='https://maxcdn.bootstrapcdn.com/bootstrap/4.2.1/js/bootstrap.min.js'></script><style>.image {background-color: #999;width: 200px;height: 200px%;display: inline-block;}.image img {width:100%;}</style><body><div class='image'><img src='https://www.aboriva.com/img/couv/34/nm_sugar-num178-juillet-aout-2016_1154.jpg'><img src='image/Linty6.jpg'></div></body></html>")
}

func main() {
	http.HandleFunc("/", linty)
	http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("/image"))))
	//http.Handle("/image/Linty6.jpg", http.FileServer(http.Dir("image/")))
	//http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":5000", nil)
}
