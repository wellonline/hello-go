package main

import (
	"hello-go/handler"
	"log"
	"net/http"
)

func main() {
	h := handler.New()
	err := http.ListenAndServe(":80", h)
	if err != nil {
		log.Fatalln(err)
	}
}
