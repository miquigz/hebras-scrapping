package controllers

import (
	"encoding/json"
	"hebras-scrapping/constants"
	"hebras-scrapping/services"
	"net/http"
)

func GetScrapeHebras(w http.ResponseWriter, r *http.Request) {
	URLS := []string{constants.TEA_BLENDS_URL}

	scrapeHebras := services.ScrapeHebras(URLS)
	if len(scrapeHebras) > 0 {
		w.WriteHeader(http.StatusOK)
		hebrasResponse, err := json.Marshal(scrapeHebras)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "Error interno del servidor"}` + err.Error()))
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(hebrasResponse)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
	return
}
