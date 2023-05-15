package handler

import (
	"encoding/json"
	"net/http"
	"poker/internal/logic"
)

func (h *Handler) ConnectRandomGameEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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

	_, err = w.Write(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetGameInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	gameID := r.URL.Query().Get("id")
	if gameID == "" || h.Games[gameID] == nil {
		ErrorResponse(w, "game does not exist", http.StatusBadRequest)
		return
	}
	response, err := json.Marshal(h.Games[gameID])
	if err != nil {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err = w.Write(response); err != nil {
		ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GenerateID TODO
func GenerateID() string {
	return "id"
}
