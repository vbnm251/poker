package handler

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"poker/internal/logic"
	"strings"
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
			log.Printf("%s quited the game %s", player.Username, gameID)
			h.Games[gameID].QuitGame(pos)
		}
	}()

	// game logic
	go func() {
		for {
		GameLoop:
			for {
				if h.Games[gameID].GetRealLength() >= 2 {
					h.Games[gameID].Live = true
					h.Games[gameID].Add()
				} else {
					h.Games[gameID].Live = false
					break GameLoop
				}

				time.Sleep(200 * time.Millisecond)
				//fmt.Println(player.Username, player.Role)

				if player.Role == logic.SmallBlind {
					log.Printf("Game %s has been started\n", gameID)

					// Preflop
					h.Games[gameID].CalculateFirstStep()
					h.Games[gameID].ShuffleDeck()
					h.Games[gameID].Distribution()
					h.Games[gameID].StartGame()
					data := map[string]interface{}{
						"event":   "gamePlayers",
						"players": h.Games[gameID].Players,
					}
					h.SendToAllPlayers(gameID, data)
					h.PeriodEnd(gameID)

					if h.Games[gameID].Live {
						// Flop
						h.Games[gameID].FlopCards()
						data = map[string]interface{}{
							"event": "flop",
							"cards": [3]logic.Card{
								h.Games[gameID].Table[0],
								h.Games[gameID].Table[1],
								h.Games[gameID].Table[2],
							},
						}
						h.SendToAllPlayers(gameID, data)
						h.PeriodEnd(gameID)
					}

					if h.Games[gameID].Live {
						// Turn
						h.Games[gameID].TurnCard()
						data = map[string]interface{}{
							"event": "turn",
							"card":  h.Games[gameID].Table[3],
						}
						h.SendToAllPlayers(gameID, data)
						h.PeriodEnd(gameID)
					}

					if h.Games[gameID].Live {
						// River
						h.Games[gameID].RiverCard()
						data = map[string]interface{}{
							"event": "river",
							"card":  h.Games[gameID].Table[4],
						}
						h.SendToAllPlayers(gameID, data)
						for {
							if h.Games[gameID].CurrentStep == -1 && h.Games[gameID].CheckBets() {
								h.Games[gameID].CurrentStep = h.Games[gameID].SmallBlindID
								if f, pl := h.Games[gameID].CheckPlayers(gameID); !f {
									pl.Balance += h.Games[gameID].Bank

									winners := []*logic.Player{pl}
									data := map[string]interface{}{
										"event":   "winners",
										"winners": winners,
										"sum":     h.Games[gameID].Bank,
									}
									h.SendToAllPlayers(gameID, data)
									h.Games[gameID].Live = false
								}
								break
							}
						}
					}

					if h.Games[gameID].Live {
						// The final
						winners := h.Games[gameID].DefineWinners()
						winnersName := make([]string, 0)
						sum := h.Games[gameID].Bank / len(winners)
						for _, player := range winners {
							winnersName = append(winnersName, player.Username)
							_ = player.SendMessage(data)
							h.Games[gameID].Players[player.Position].Balance += sum
						}
						data = map[string]interface{}{
							"event":   "winners",
							"winners": winners,
							"sum":     sum,
						}
						h.SendToAllPlayers(gameID, data)
						winnersStr := strings.Join(winnersName, ",")
						log.Printf("Winners of game %s are: %s", gameID, winnersStr)
					}

					h.Games[gameID].Disable()
				}

				h.Games[gameID].Wait()

				if player.Role == logic.SmallBlind {
					h.Games[gameID].Add()
				}

				time.Sleep(5 * time.Second) //pause for players

				if player.Role == logic.SmallBlind {
					h.Games[gameID].ClearGame()
					h.Games[gameID].RotateRoles()
					h.Games[gameID].Disable()
				}

				h.Games[gameID].Wait()
				//fmt.Println("BREAK", player.Username)
			}
		}

	}()

	// get bets
	for {
		if err := h.GetUserBest(gameID, pos, conn); err != nil {
			return
		}
	}

}
