package routes

import (
	"github.com/gorilla/mux"
	"hebras-scrapping/controllers"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/scrape/hebras", controllers.GetScrapeHebras).Methods("GET")

	return r
}
