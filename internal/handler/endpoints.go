package handler

import (
	"encoding/json"
	"net/http"
	"poker/internal/logic"
)

func (h *Handler) ConnectRandomGameEndpoint(w http.ResponseWriter, r *http.Request) {
	gameID := "хуй"
	for id, game := range h.Games {
		if len(game.Players) != logic.MaxPlayers {
			gameID = id
		}
	}
	if gameID == "хуй" {
		gameID = GenerateID()
		h.Games[gameID] = logic.NewGame()
	}

	response, err := json.Marshal(map[string]interface{}{
		"id": gameID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GenerateID TODO
func GenerateID() string {
	return "id"
}
