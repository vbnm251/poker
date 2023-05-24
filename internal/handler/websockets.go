package handler

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"poker/internal/logic"
	"time"
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
		StatusResponse("Already max game", &player)
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
				time.Sleep(200 * time.Millisecond)
				log.Printf("Game %s has been started\n", gameID)
				//h.Games[gameID].RotateRoles()

				// Preflop
				h.Games[gameID].ShuffleDeck()
				h.Games[gameID].Distribution()
				h.Games[gameID].StartGame()
				data := map[string]interface{}{
					"event": "gamePlayers",
					"game":  h.Games[gameID].Players,
				}
				h.SendToAllPlayers(gameID, data)
			}

			//get users bets here
			h.GetUserBest(gameID, pos, conn)

			// the end of preflop
			// flop starts
			if player.Role == logic.SmallBlind {
				h.PeriodEnd(gameID)
				h.Games[gameID].FlopCards()
				data := map[string]interface{}{
					"event": "flop",
					"cards": [3]logic.Card{
						h.Games[gameID].Table[0],
						h.Games[gameID].Table[1],
						h.Games[gameID].Table[2],
					},
				}
				h.SendToAllPlayers(gameID, data)
			}

			h.GetUserBest(gameID, pos, conn)

			//the end of flop -> turn starts
			if player.Role == logic.SmallBlind {
				h.PeriodEnd(gameID)
				h.Games[gameID].TurnCard()
				data := map[string]interface{}{
					"event": "turn",
					"card":  h.Games[gameID].Table[3],
				}
				h.SendToAllPlayers(gameID, data)
			}

			h.GetUserBest(gameID, pos, conn)

			// the end of turn -> river
			if player.Role == logic.SmallBlind {
				h.PeriodEnd(gameID)
				h.Games[gameID].RiverCard()
				data := map[string]interface{}{
					"event": "river",
					"card":  h.Games[gameID].Table[4],
				}
				h.SendToAllPlayers(gameID, data)
			}

			h.GetUserBest(gameID, pos, conn)

			if player.Role == logic.SmallBlind {
				h.PeriodEnd(gameID)
				winners := h.Games[gameID].DefineWinners()
				for _, player := range winners {
					data := map[string]interface{}{
						"event":  "status",
						"status": "WINNER",
					}
					_ = player.SendMessage(data)
				}
			}

		}
	}
}
