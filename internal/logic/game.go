/*
This package contains all classic poker logic

The game:
	1. Rotate game
	2. Shuffle deck
	3. Distribution
	4. PreFlop: waiting for game' bets
	5. Flop: waiting for game' bets
	6. Turn: waiting for game' bets
	7. River: waiting for game' bets
	8. Congratulations winner
*/

package logic

import (
	"errors"
)

const MaxPlayers = 7

const Hearts = "hearts"
const Diamonds = "diamonds"
const Clubs = "clubs"
const Spades = "spades"

const BigBlind = "big_blind"
const SmallBlind = "small_blind"
const Dealer = "dealer"
const Regular = "regular"

// Game contains all game information
// It provides all methods for running poker game
type Game struct {
	Live         bool      `json:"live"`
	SmallBlindID int       `json:"-"`
	Players      []*Player `json:"game"`
	CurrentStep  int       `json:"currentStep"`
	Deck         []Card    `json:"-"`
	Table        [5]Card   `json:"table"`
	CurrentBet   int       `json:"current_bet"`
	Bank         int       `json:"bank"`
	DeckInd      int       `json:"-"`

	RaiseID int `json:"raiseID"`

	WaitSmallBlind bool `json:"-"`
}

func NewGame() *Game {
	game := &Game{
		Live:         false,
		SmallBlindID: 0,
		Players:      make([]*Player, MaxPlayers),
		Deck:         GenerateDeck(),
		Table:        [5]Card{},
		DeckInd:      0,
		Bank:         0,
		CurrentBet:   0,

		WaitSmallBlind: true,
	}
	for i := 0; i < MaxPlayers; i++ {
		game.Players[i] = nil
	}

	return game
}

func (g *Game) ClearGame() {
	g.Table = [5]Card{}
	g.Bank = 0
	g.DeckInd = 0
	g.CurrentBet = 0
	g.Live = false
	g.WaitSmallBlind = true
	g.ClearBets()
}

func (g *Game) ClearBets() {
	for _, player := range g.Players {
		if player != nil {
			player.CurrentBet = 0
		}
	}
}

// JoinGame Return Free position and error in case there are max game in the game
func (g *Game) JoinGame(player *Player) (int, error) {
	switch g.GetRealLength() {
	case 0:
		player.Role = SmallBlind
	case 1:
		player.Role = BigBlind
	case 2:
		player.Role = Dealer
	case MaxPlayers:
		return 0, errors.New("the game is full")
	default:
		player.Role = Regular
	}

	var freePosition int
	for i := 0; i < MaxPlayers; i++ {
		if g.Players[i] == nil {
			freePosition = i
			break
		}
	}

	g.Players[freePosition] = player
	player.Position = freePosition

	return freePosition, nil
}

func (g *Game) QuitGame(pos int) {
	g.Players[pos] = nil
}

// CheckPlayers returns true in case game contains at least 2 game
// In other way it returns false and winner
func (g *Game) CheckPlayers() (bool, *Player) {
	inGamePlayers := 0
	var pl *Player
	for _, player := range g.Players {
		if player != nil && player.InGame {
			inGamePlayers++
			pl = player
		}
	}
	if inGamePlayers == 1 {
		return false, pl
	}
	return true, nil
}

func (g *Game) CalculateFirstStep() int {
	for i := g.SmallBlindID; i < g.SmallBlindID+MaxPlayers; i++ {
		j := i % 7
		if g.Players[j] != nil && g.Players[j].InGame {
			g.CurrentStep = j
			return j
		}
	}
	return -1
}

// StartGame sends role and hand to every player
func (g *Game) StartGame() {
	g.CurrentStep = g.SmallBlindID
	for _, player := range g.Players {
		if player != nil {
			player.InGame = true
			data := map[string]interface{}{
				"event": "distribution",
				"role":  player.Role,
				"cards": player.Cards,
			}
			_ = player.SendMessage(data)
		}
	}
}

// CalculateNextStep TODO
func (g *Game) CalculateNextStep() int {
	for i := g.CurrentStep + 1; i < g.CurrentStep+MaxPlayers; i++ {
		if i%7 == g.RaiseID { // replace smallblindid for raise id
			break
		}
		if g.Players[i%MaxPlayers] != nil && g.Players[i%MaxPlayers].InGame {
			g.CurrentStep = i % MaxPlayers
			return i % MaxPlayers
		}
	}
	g.CurrentStep = -1
	return -1
}

func (g *Game) CheckBets() bool {
	if g.CurrentBet == 0 {
		return false
	}
	for _, player := range g.Players {
		if player != nil {
			if player.CurrentBet < g.CurrentBet {
				return false
			}
		}
	}
	return true
}

func (g *Game) GetRealLength() int {
	cnt := 0
	for i := 0; i < MaxPlayers; i++ {
		if g.Players[i] != nil {
			cnt++
		}
	}
	return cnt
}

func (g *Game) Add() {
	g.WaitSmallBlind = false
}

func (g *Game) Disable() {
	g.WaitSmallBlind = true
}

func (g *Game) Wait() {
	for {
		if g.WaitSmallBlind {
			break
		}
	}
}
