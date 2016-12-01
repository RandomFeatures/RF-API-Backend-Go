package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	router := httprouter.New()
	GetGameRoutes(router)
	GetNewsRoutes(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
