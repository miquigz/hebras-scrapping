package controllers

import (
	"encoding/json"
	"hebras-scrapping/constants"
	"hebras-scrapping/services"
	"net/http"
)

type HebrasController struct {
	service *services.HebrasService
}

func NewHebrasController() *HebrasController {
	return &HebrasController{
		service: services.NewHebrasService(),
	}
}

func (hc *HebrasController) GetScrapeHebras(w http.ResponseWriter, r *http.Request) {
	URLS := []string{constants.TEA_BLENDS_URL, constants.TEA_CONNECTION_URL}

	scrapeHebras := hc.service.ScrapeHebras(URLS)
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
