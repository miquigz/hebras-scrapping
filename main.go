package main

import (
	"hebras-scrapping/routes"
	"log"
	"net/http"
)

func main() {
	r := routes.Router()

	log.Fatal(http.ListenAndServe(":8080", r))
}
