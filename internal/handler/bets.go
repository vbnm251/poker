package handler

import (
	"github.com/gorilla/websocket"
	"poker/internal/logic"
)

func (h *Handler) GetUserBest(gameID string, pos int, conn *websocket.Conn) error {
	var action logic.Action

	game := h.Games[gameID]

	for {
		if game.CurrentStep == pos {
			break
		}
	}

	if err := conn.ReadJSON(&action); err != nil {
		return err
	}

	if action.Action == logic.Fold {
		game.Players[pos].InGame = false
	} else {
		game.Players[pos].CurrentBet += action.Sum
		game.Bank += action.Sum
		game.Players[pos].Balance -= action.Sum
		if action.Action == logic.Raise {
			game.CurrentBet = action.Sum
			game.RaiseID = pos
		}
	}

	action.Next = game.CalculateNextStep()
	h.SendToAllPlayers(gameID, action)

	return nil
}

func (h *Handler) PeriodEnd(gameID string) {
	game := h.Games[gameID]

	for {
		if h.Games[gameID].CurrentStep == -1 && game.CheckBets() {
			h.Games[gameID].CurrentStep = h.Games[gameID].SmallBlindID
			if f, pl := h.Games[gameID].CheckPlayers(gameID); !f {
				_ = pl.SendMessage(map[string]interface{}{
					"event": "winner",
				})
				h.Games[gameID].Live = false
			}
			break
		}
	}
	game.ClearBets()
	h.Games[gameID].RaiseID = h.Games[gameID].SmallBlindID
	h.Games[gameID].CurrentBet = 0
	pos := h.Games[gameID].CalculateFirstStep()

	data := map[string]interface{}{
		"event": "step",
		"pos":   pos,
	}
	h.SendToAllPlayers(gameID, data)
}
