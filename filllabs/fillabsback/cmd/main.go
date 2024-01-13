package main

import (
	"fillabs_intern_project/pkg/handler"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	handler, err := handler.InitHandler()
	if err != nil {
		panic(err)
	}
	cors := cors.AllowAll().Handler(handler)
	log.Fatal(http.ListenAndServe(":8080", cors))
}
