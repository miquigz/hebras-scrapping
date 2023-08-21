package controllers

import (
	"encoding/json"
	"hebras-scrapping/services"
	"net/http"
)

func GetScrapeHebras(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Access-Control-Allow-Methods", "GET")

	scrapeHebras := services.ScrapeHebras(true)
	if len(scrapeHebras) > 0 {
		w.WriteHeader(http.StatusOK)
		hebrasResponse, e := json.Marshal(scrapeHebras)
		if e != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "Error interno del servidor"}` + e.Error()))
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(hebrasResponse)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
	return
}
