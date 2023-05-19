/*
The game:
	1. Rotate players
	2. Shuffle deck
	3. Distribution
	4. PreFlop: waiting for players' bets
	5. Flop: waiting for players' bets
	6. Turn: waiting for players' bets
	7. River: waiting for players' bets
	8. Congratulations winner
*/

package logic

import (
	"errors"
	"math/rand"
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

type Game struct {
	IsGameLive   bool      `json:"live"`
	SmallBlindID int       `json:"-"`
	Players      []*Player `json:"players"`
	Deck         []Card    `json:"-"`
	Table        [5]Card   `json:"table"`
	CurrentBet   int       `json:"current_bet"`
	Bank         int       `json:"bank"`
	DeckInd      int       `json:"-"`
}

func NewGame() *Game {
	game := &Game{
		IsGameLive:   false,
		SmallBlindID: 0,
		Players:      make([]*Player, MaxPlayers),
		Deck:         GenerateDeck(),
		Table:        [5]Card{},
		DeckInd:      0,
		Bank:         0,
		CurrentBet:   0,
	}
	for i := 0; i < MaxPlayers; i++ {
		game.Players[i] = nil
	}

	return game
}

// JoinGame Return Free position and error in case there are max players in the game
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

func (g *Game) RotateRoles() {
	if g.GetRealLength() != 1 {
		for i := g.SmallBlindID + 1; i < g.SmallBlindID+MaxPlayers; i++ {
			if g.Players[i%MaxPlayers] != nil {
				g.SmallBlindID = i % MaxPlayers
				break
			}
		}
		g.Players = append(g.Players[g.SmallBlindID:], g.Players[:g.SmallBlindID]...)
		g.Players[0].Role = SmallBlind
		bbID := 0
		for i := 1; i < MaxPlayers; i++ {
			if g.Players[i] != nil {
				g.Players[i].Role = BigBlind
				bbID = i
				break
			}
		}
		for i := bbID + 1; i < MaxPlayers; i++ {
			if g.Players[i] != nil {
				g.Players[i].Role = Regular
			}
		}
		for i := MaxPlayers - 1; i > bbID; i++ {
			if g.Players[i] != nil && g.Players[i].Role == Regular {
				g.Players[i].Role = Dealer
				break
			}
		}
	}
}

func (g *Game) DefineWinners() []*Player {
	winners := make([]*Player, 0)
	maxCode := 0
	var bestKickerValue int = 0
	var bestFirstValue int = 0
	var bestSecondValue int = 0

	for _, player := range g.Players {
		currentCode := 0
		var firstValue int
		var secondValue int

		player.GetCombination(g.Table)

		switch player.Combination.(type) {
		case FlushRoyal:
			currentCode = 10
			comb := player.Combination.(FlushRoyal)
			firstValue = comb.HighCard.Value
		case StraightFlush:
			currentCode = 9
			comb := player.Combination.(StraightFlush)
			firstValue = comb.HighCard.Value
		case FourOfAKind:
			currentCode = 8
			comb := player.Combination.(FourOfAKind)
			firstValue = comb.HighCard.Value
		case FullHouse:
			currentCode = 7
			comb := player.Combination.(FullHouse)
			firstValue = comb.HighTriple.Value
			secondValue = comb.HighPair.Value
		case Flush:
			currentCode = 6
			comb := player.Combination.(Flush)
			firstValue = comb.HighCard.Value
		case Straight:
			currentCode = 5
			comb := player.Combination.(Straight)
			firstValue = comb.HighCard.Value
		case Set:
			currentCode = 4
			comb := player.Combination.(Set)
			firstValue = comb.HighCard.Value
		case TwoPair:
			currentCode = 3
			comb := player.Combination.(TwoPair)
			firstValue = comb.FirsPair.Value
			secondValue = comb.SecondPair.Value
		case Pair:
			currentCode = 2
			comb := player.Combination.(Pair)
			firstValue = comb.HighCard.Value
		case HighCard:
			currentCode = 1
			comb := player.Combination.(HighCard)
			firstValue = comb.HighCard.Value
		}

		if currentCode > maxCode {
			maxCode = currentCode
			winners = []*Player{player}
			bestKickerValue = player.Kicker.Value
			bestFirstValue = firstValue
			bestSecondValue = secondValue
		} else if currentCode == maxCode {
			if currentCode == 3 || currentCode == 7 {
				if (firstValue > bestFirstValue) || (firstValue == bestFirstValue && secondValue > bestSecondValue) {
					//player has better combination
					winners = []*Player{player}
					bestFirstValue = firstValue
					bestSecondValue = secondValue
					bestKickerValue = player.Kicker.Value
				} else if firstValue == bestFirstValue && secondValue == bestSecondValue {
					//checking kicker
					if player.Kicker.Value > bestKickerValue {
						bestKickerValue = player.Kicker.Value
					} else if player.Kicker.Value == bestKickerValue {
						winners = append(winners, player)
					}
				}
			} else {
				if firstValue > bestFirstValue {
					//player has better combination
					winners = []*Player{player}
					bestFirstValue = firstValue
					bestKickerValue = player.Kicker.Value
				} else if firstValue == bestFirstValue {
					//checking kicker
					if player.Kicker.Value > bestKickerValue {
						bestKickerValue = player.Kicker.Value
					} else if player.Kicker.Value == bestKickerValue {
						winners = append(winners, player)
					}
				}
			}
		}
	}

	return winners
}

func (g *Game) Distribution() {
	for i := range g.Players {
		if g.Players[i] != nil {
			g.Players[i].Cards[0] = g.Deck[g.DeckInd]
			g.DeckInd++
			g.Players[i].Cards[1] = g.Deck[g.DeckInd]
			g.DeckInd++
		}
	}
}

func (g *Game) ShuffleDeck() {
	rand.Shuffle(len(g.Deck), func(i, j int) {
		g.Deck[i], g.Deck[j] = g.Deck[j], g.Deck[i]
	})
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
