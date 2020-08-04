package main

import (
	"log"
	"net/http"
)

type helloHandler struct{}

func main() {
	//http.HandleFunc("/", hello)
	//log.Println("Starting http server")
	//log.Fatal(http.ListenAndServe(":4000",nil))

	http.Handle("/",&helloHandler{})
	log.Println("Starting http server")
	log.Fatal(http.ListenAndServe(":4000",nil))
}

//func hello(writer http.ResponseWriter, request *http.Request) {
//	writer.Write([]byte("hello world"))
//}

func (_ *helloHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}