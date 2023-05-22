package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"poker/internal/logic"
)

// SendToAllPlayers sends data to every player
func (h *Handler) SendToAllPlayers(gameID string, data map[string]interface{}) {
	for _, player := range h.Games[gameID].Players {
		if player != nil {
			_ = player.SendMessage(data)
		}
	}
}

func StatusResponse(status string, player *logic.Player) {
	log.Println("websocket status:", status)
	data := map[string]interface{}{
		"status": status,
	}
	err := player.SendMessage(data)
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
