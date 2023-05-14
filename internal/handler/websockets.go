package handler

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	"log"
	"net/http"
	"poker/internal/logic"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1 << 10,
	WriteBufferSize: 1 << 10,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WsInput struct {
	Action string                 `json:"action"`
	Data   map[string]interface{} `json:"data"`
}

func (h *Handler) WebsocketsEndpoint(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	gameID := r.URL.Query().Get("id")
	if h.Games[gameID] == nil {
		StatusResponse(fmt.Sprintf("Game %s does not exists", gameID), conn)
		return
	}

	log.Println("Client connected to game", gameID)

	for {
		var input WsInput
		if err := conn.ReadJSON(&input); err != nil {
			log.Println(err)
			return
		}

		switch input.Action {
		case "new_player":
			var pl logic.Player
			if err := mapstructure.Decode(input.Data, &pl); err != nil {
				log.Println(err)
				continue
			}

			log.Printf("Player %s connected to game %s\n", pl.Username, gameID)

			player := logic.NewPlayer(pl.Username, 5000, conn)
			pos, err := h.Games[gameID].JoinGame(&player)
			if err != nil {
				StatusResponse("Already max players", conn)
			} else {
				data := map[string]interface{}{
					"event":    "new_player",
					"position": pos,
				}
				h.SendToAllPlayers(gameID, data)
			}
		default:
			StatusResponse("invalid action", conn)
			continue
		}
	}
}
