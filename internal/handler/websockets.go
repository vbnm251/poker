package handler

import (
	"fmt"
	"github.com/gorilla/websocket"
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

	var player logic.Player
	if err := conn.ReadJSON(&player); err != nil {
		log.Println(err)
		return
	}
	gameID := r.URL.Query().Get("id")
	if h.Games[gameID] == nil {
		StatusResponse(fmt.Sprintf("Game %s does not exists", gameID), &player)
		return
	}

	player = logic.NewPlayer(player.Username, 5000, conn)
	pos, err := h.Games[gameID].JoinGame(&player)
	if err != nil {
		StatusResponse("Already max players", &player)
		return
	} else {
		data := map[string]interface{}{
			"event":  "new_player",
			"player": player,
		}
		h.SendToAllPlayers(gameID, data)
	}
	log.Printf("Player %s connected to game %s\n", player.Username, gameID)
	defer func() {
		if h.Games[gameID] != nil {
			log.Println("DELETED", player.Username)
			h.Games[gameID].QuitGame(pos)
		}
	}()

	for {
		if h.Games[gameID].GetRealLength() >= 2 {
			h.Games[gameID].Live = true
		} else {
			h.Games[gameID].Live = false
		}

		for {
			if !h.Games[gameID].Live || h.Games[gameID].GetRealLength() < 2 {
				break
			}

			// This actions must be done only once
			// So it happens only in small blind player
			if player.Role == logic.SmallBlind {
				//h.Games[gameID].RotateRoles()
				log.Printf("Game %s has been started\n", gameID)
				h.Games[gameID].ShuffleDeck()
				h.Games[gameID].Distribution()
				h.Games[gameID].TableCards()
				h.Games[gameID].StartGame()

				//PREFLOP
				data := map[string]interface{}{
					"event": "preflop",
					"cards": [3]logic.Card{
						h.Games[gameID].Table[0],
						h.Games[gameID].Table[1],
						h.Games[gameID].Table[2],
					},
				}
				h.SendToAllPlayers(gameID, data)

			}
			//todo add queue
			var action logic.Action
			if err := conn.ReadJSON(&action); err != nil {
				log.Println(err)
				return
			}
			if action.Action == logic.Fold {
				player.InGame = false
			} else if action.Action == logic.Raise {
				player.CurrentBet = action.Sum
				h.Games[gameID].CurrentBet = action.Sum
			} else if action.Action == logic.Call {
				player.CurrentBet = action.Sum
			}

			// check if at least one player
			//if f, pl := h.Games[gameID].CheckPlayers(); !f && player.Role == logic.SmallBlind {
			//	_ = pl.Conn.WriteJSON(map[string]interface{}{
			//		"status": "WINNER",
			//	})
			//}

		}
	}
}
