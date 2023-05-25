package handler

import (
	"github.com/gorilla/websocket"
	"log"
	"poker/internal/logic"
)

func (h *Handler) GetUserBest(gameID string, pos int, conn *websocket.Conn) {
	var action logic.Action
	for {
		if h.Games[gameID].CurrentStep == pos {
			break
		}
	}
	if err := conn.ReadJSON(&action); err != nil {
		log.Println(err)
		return
	}

	if action.Action == logic.Fold {
		h.Games[gameID].Players[pos].InGame = false
	} else if action.Action == logic.Raise {
		h.Games[gameID].Players[pos].CurrentBet = action.Sum
		h.Games[gameID].CurrentBet = action.Sum
		h.Games[gameID].Bank += action.Sum
	} else if action.Action == logic.Call {
		h.Games[gameID].Players[pos].CurrentBet = action.Sum
		h.Games[gameID].Bank += action.Sum
	}

	action.Next = h.Games[gameID].CalculateNextStep()
	h.SendToAllPlayers(gameID, action)
}

func (h *Handler) PeriodEnd(gameID string) {
	for {
		if h.Games[gameID].CurrentStep == -1 {
			h.Games[gameID].CurrentStep = h.Games[gameID].SmallBlindID
			if f, pl := h.Games[gameID].CheckPlayers(); !f {
				_ = pl.SendMessage(map[string]interface{}{
					"event":  "status",
					"status": "WINNER",
				})
				h.Games[gameID].Live = false
			}
			break
		}
	}
	pos := h.Games[gameID].CalculateFirstStep()
	data := map[string]interface{}{
		"event": "step",
		"pos":   pos,
	}
	h.SendToAllPlayers(gameID, data)
}
