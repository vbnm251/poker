package handler

import (
	"github.com/gorilla/websocket"
)

func (h *Handler) SendToAllPlayers(gameID string, data map[string]interface{}) {
	for _, player := range h.Games[gameID].Players {
		if player != nil {
			_ = player.Conn.WriteJSON(data)
		}
	}
}

func StatusResponse(status string, conn *websocket.Conn) {
	data := map[string]interface{}{
		"status": status,
	}

	err := conn.WriteJSON(data)
	if err != nil {
		return
	}
}
