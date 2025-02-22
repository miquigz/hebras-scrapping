package routes

import (
	"github.com/gorilla/mux"
	"hebras-scrapping/controllers"
)

func Router() *mux.Router {
	hebrasController := controllers.NewHebrasController()

	r := mux.NewRouter()
	const version1 = "/api/1"
	r.HandleFunc(version1+"/scrape/hebras", hebrasController.GetScrapeHebras).Methods("GET")
	r.HandleFunc("/tea/messages", hebrasController.WsHandler)

	return r
}
