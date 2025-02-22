package controllers

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"hebras-scrapping/constants"
	"hebras-scrapping/services"
	"log"
	"net/http"
	"time"
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

// WsHandler con el fin de actualizar la lista de hebras cada vez que el cache es actualizado
func (hc *HebrasController) WsHandler(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{ //ws config buffers sizes
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return //Si no
	}
	defer conn.Close()

	//Nos suscribimos al topic de NATS "scrape.hebras"
	subTea, err := hc.service.Nc.SubscribeSync("scrape.hebras")
	if err != nil {
		log.Println("Error al suscribirse a NATS:", err)
		return
	}
	defer subTea.Unsubscribe() //al final desuscribimos para evitar memory leaks(este caso puede que nunca se de a menos de q suceda un panic)

	// Escuchamos mensajes de NATS y los retornamos por ws
	for {
		msg, err := subTea.NextMsg(time.Minute)
		if err != nil {
			log.Println("Error al recibir mensaje de NATS:", err) //TODO:Puede spamear mucho para logs, probablemente mejor comentar log o filtrar tipos de errores
			continue
		} else if msg != nil && msg.Data != nil {
			if err := conn.WriteMessage(websocket.TextMessage, msg.Data); err != nil {
				log.Println("Error al escribir mensaje en WebSocket:", err)
				return
			}
		}
	}
}
