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
	"math/rand"
)

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

func (g *Game) FlopCards() {
	for i := 0; i < 3; i++ {
		g.Table[i] = g.Deck[g.DeckInd]
		g.DeckInd++
	}
}

func (g *Game) TurnCard() {
	g.Table[3] = g.Deck[g.DeckInd]
	g.DeckInd++
}

func (g *Game) RiverCard() {
	g.Table[4] = g.Deck[g.DeckInd]
	g.DeckInd++
}

func (g *Game) RotateRoles() {
	if g.GetRealLength() != 1 {
		bbID, sbID := 0, 0
		upBorder := g.SmallBlindID + MaxPlayers + 1

		//looking for new small blind
		for i := g.SmallBlindID + 1; i < upBorder; i++ {
			if g.Players[i%7] != nil {
				g.SmallBlindID = i % 7
				sbID = i
				g.Players[i%7].Role = "small_blind"
				break
			}
		}
		// looking for new big blind
		for i := sbID + 1; i < upBorder; i++ {
			if g.Players[i%7] != nil {
				g.Players[i%7].Role = "big_blind"
				bbID = i
				break
			}
		}
		//regular game
		for i := bbID + 1; i < upBorder; i++ {
			if g.Players[i%7] != nil {
				g.Players[i%7].Role = "regular"
			}
		}
		//dealer
		for i := upBorder - 1; i > bbID; i-- {
			if g.Players[i%7] != nil {
				g.Players[i%7].Role = "dealer"
				break
			}
		}

		g.RaiseID = g.SmallBlindID
	}
}

func (g *Game) DefineWinners() []*Player {
	winners := make([]*Player, 0)
	maxCode := 0
	var bestKickerValue int = 0
	var bestFirstValue int = 0
	var bestSecondValue int = 0

	for _, player := range g.Players {
		if player != nil {

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
	}

	return winners
}
