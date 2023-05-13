package handler

import (
	"fmt"
	"log"
	"net/http"
	"poker/internal/logic"
)

func (h *Handler) ConnectRandomGameEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Println("Congratulations")

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

	_, _ = w.Write([]byte(fmt.Sprintf(`
	<!DOCTYPE html>
		<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title>Redirect Page</title>
				<script>
					window.location.href = "/connect?id=%s";
				</script>
			</head>
		</html> `,
		gameID)))
}

// GenerateID TODO
func GenerateID() string {
	return "id"
}
