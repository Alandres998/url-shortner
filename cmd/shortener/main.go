package main

import (
	"log"
	"net/http"
)

type MyHandler struct{}

func Shorter(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Зашифровать"))
}

func Fuller(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Расшировать"))
}

func main() {
	http.HandleFunc("/", Shorter)
	http.HandleFunc("/{id}", Fuller)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Ахтунг сервер прилег: %s\n", err)
	}
}
