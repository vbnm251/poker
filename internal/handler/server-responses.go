package handler

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func (h *Handler) SendToAllPlayers(gameID string, data map[string]interface{}) {
	for _, player := range h.Games[gameID].Players {
		if player != nil {
			_ = player.Conn.WriteJSON(data)
		}
	}
}

func StatusResponse(status string, conn *websocket.Conn) {
	log.Println("websocket status:", status)
	data := map[string]interface{}{
		"status": status,
	}
	err := conn.WriteJSON(data)
	if err != nil {
		return
	}
}

func ErrorResponse(w http.ResponseWriter, error string, code int) {
	log.Println("error:", error)
	w.WriteHeader(code)
	data, err := json.Marshal(map[string]interface{}{
		"error": error,
	})
	if err != nil {
		return
	}
	if _, err := w.Write(data); err != nil {
		return
	}
}
