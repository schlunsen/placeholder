package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	log.Println("Starting placeholder service...")
	router := httprouter.New()
	router.GET("/:width/:height", ImageHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
