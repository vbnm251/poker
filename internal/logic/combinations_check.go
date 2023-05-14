package logic

import (
	"sort"
)

func (p *Player) FlushRoyalCheck(table [5]Card, combinationFound *bool) {
	suits := make(map[string][]Card, 0)

	suits[p.Cards[0].Suit] = append(suits[p.Cards[0].Suit], p.Cards[0])
	suits[p.Cards[1].Suit] = append(suits[p.Cards[1].Suit], p.Cards[1])

	for _, card := range table {
		suits[card.Suit] = append(suits[card.Suit], card)
	}

	for suit := range suits {
		sort.Slice(suits[suit], func(i, j int) bool {
			return suits[suit][j].Value < suits[suit][i].Value
		})
		if len(suits[suit]) >= 5 && suits[suit][0].Value == 14 && suits[suit][1].Value == 13 && suits[suit][2].Value == 12 && suits[suit][3].Value == 11 && suits[suit][4].Value == 10 {
			p.Combination = FlushRoyal{HighCard: suits[suit][0]}
			*combinationFound = true
		}
	}
}

func (p *Player) StraightFlushCheck(table [5]Card, combinationFound *bool) {
	suits := make(map[string][]Card, 0)

	suits[p.Cards[0].Suit] = append(suits[p.Cards[0].Suit], p.Cards[0])
	suits[p.Cards[1].Suit] = append(suits[p.Cards[1].Suit], p.Cards[1])

	for _, card := range table {
		suits[card.Suit] = append(suits[card.Suit], card)
	}

	for suit := range suits {
		sort.Slice(suits[suit], func(i, j int) bool {
			return suits[suit][j].Value > suits[suit][i].Value
		})
		if len(suits[suit]) >= 5 {
			n := len(suits[suit])
			i, length, bestLength, m := n-1, 1, 1, 0
			for j := 0; j < n; j++ {
				if suits[suit][j].Value-suits[suit][i].Value != 1 && !(suits[suit][i].Value == 14 && suits[suit][j].Value == 2) {
					if length >= bestLength {
						m = i
						bestLength = length
					}
					i = j
					length = 1
				} else {
					i = (i + 1) % n
					length++
				}
			}
			if length >= bestLength {
				m = n - 1
				bestLength = length
			}
			if bestLength >= 5 {
				p.Combination = StraightFlush{HighCard: suits[suit][m]}
				*combinationFound = true
			}

		}
	}
}

func (p *Player) FourOfAKindCheck(table [5]Card, combinationFound *bool) {
	counter := make(map[int]int, 0)

	counter[p.Cards[0].Value]++
	counter[p.Cards[1].Value]++

	for _, card := range table {
		counter[card.Value]++
	}

	for value, cnt := range counter {
		if cnt == 4 {
			p.Combination = FourOfAKind{HighCard: Card{Value: value, Suit: p.GetSuitByValue(value, table)}}
			*combinationFound = true
		}
	}
}

func (p *Player) FullHouseCheck(table [5]Card, combinationFound *bool) {
	counter := make(map[int]int)
	counter[p.Cards[0].Value]++
	counter[p.Cards[1].Value]++

	for _, card := range table {
		counter[card.Value]++
	}

	var triple Card
	var double Card
	tripleFound := false
	doubleFound := false

	for value, cnt := range counter {
		if cnt == 2 {
			double = Card{Value: value, Suit: p.GetSuitByValue(value, table)}
			doubleFound = true
		} else if cnt == 3 {
			triple = Card{Value: value, Suit: p.GetSuitByValue(value, table)}
			tripleFound = true
		}
	}

	if tripleFound && doubleFound {
		p.Combination = FullHouse{
			HighTriple: triple,
			HighPair:   double,
		}
		*combinationFound = true
	}
}

func (p *Player) FlushCheck(table [5]Card, combinationFound *bool) {
	suits := make(map[string][]Card, 0)

	suits[p.Cards[0].Suit] = append(suits[p.Cards[0].Suit], p.Cards[0])
	suits[p.Cards[1].Suit] = append(suits[p.Cards[1].Suit], p.Cards[1])

	for _, card := range table {
		suits[card.Suit] = append(suits[card.Suit], card)
	}

	for suit, cards := range suits {
		if len(cards) >= 5 {
			m := 0
			for _, card := range cards {
				if card.Value > m {
					m = card.Value
				}
			}
			p.Combination = Flush{HighCard: Card{Value: m, Suit: suit}}
			*combinationFound = true
		}
	}
}

func (p *Player) StraightCheck(table [5]Card, combinationFound *bool) {
	set := make(map[int]Card)
	cards := make([]Card, 0)

	set[p.Cards[0].Value] = p.Cards[0]
	set[p.Cards[1].Value] = p.Cards[1]

	for _, card := range table {
		set[card.Value] = card
	}
	for _, card := range set {
		cards = append(cards, card)
	}

	if len(cards) >= 5 {
		sort.Slice(cards, func(i, j int) bool {
			return cards[j].Value > cards[i].Value
		})
		n := len(cards)
		i, length, bestLength, m := n-1, 1, 1, 0
		for j := 0; j < n; j++ {
			if cards[j].Value-cards[i].Value != 1 && !(cards[i].Value == 14 && cards[j].Value == 2) {
				if length >= bestLength {
					m = i
					bestLength = length
				}
				i = j
				length = 1
			} else {
				i = (i + 1) % n
				length++
			}
		}
		if length >= bestLength {
			m = n - 1
			bestLength = length
		}
		if bestLength >= 5 {
			p.Combination = Straight{HighCard: cards[m]}
			*combinationFound = true
		}
	}
}

func (p *Player) SetCheck(table [5]Card, combinationFound *bool) {
	counter := make(map[int]int, 0)

	counter[p.Cards[0].Value]++
	counter[p.Cards[1].Value]++

	for _, card := range table {
		counter[card.Value]++
	}

	for value, cnt := range counter {
		if cnt == 3 {
			p.Combination = Set{HighCard: Card{Value: value, Suit: p.GetSuitByValue(value, table)}}
			*combinationFound = true
		}
	}
}

func (p *Player) TwoPairCheck(table [5]Card, combinationFound *bool) {
	counter := make([]int, 15)

	counter[p.Cards[0].Value]++
	counter[p.Cards[1].Value]++

	for _, card := range table {
		counter[card.Value]++
	}

	var first Card
	var second Card
	pair1Found := false
	pair2Found := false
	for i := len(counter) - 1; i >= 0; i-- {
		if counter[i] == 2 {
			if !pair1Found {
				first = Card{Value: i, Suit: p.GetSuitByValue(i, table)}
				pair1Found = true
			} else if !pair2Found {
				second = Card{Value: i, Suit: p.GetSuitByValue(i, table)}
				pair2Found = true
			}
		}
		if pair1Found && pair2Found {
			p.Combination = TwoPair{FirsPair: first, SecondPair: second}
			*combinationFound = true
			break
		}
	}

}

func (p *Player) PairCheck(table [5]Card, combinationFound *bool) {
	counter := make(map[int]int, 15)

	counter[p.Cards[0].Value]++
	counter[p.Cards[1].Value]++

	for _, card := range table {
		counter[card.Value]++
	}

	for value, cnt := range counter {
		if cnt == 2 {
			p.Combination = Pair{HighCard: Card{Value: value, Suit: p.GetSuitByValue(value, table)}}
			*combinationFound = true
		}
	}
}

func (p *Player) HighCardCheck(table [5]Card, combinationFound *bool) {
	if p.Cards[0].Value > p.Cards[1].Value {
		p.Combination = HighCard{HighCard: p.Cards[0]}
	} else {
		p.Combination = HighCard{HighCard: p.Cards[1]}
	}
	*combinationFound = true
}

func (p *Player) FindKicker(table [5]Card) {
	kicker := Card{Suit: "", Value: 0}

	for _, card := range p.Cards {
		if card.Value > kicker.Value {
			kicker = card
		}
	}
	for _, card := range table {
		if card.Value > kicker.Value {
			kicker = card
		}
	}

	p.Kicker = kicker
}

func (p *Player) GetSuitByValue(value int, table [5]Card) string {
	for _, card := range p.Cards {
		if card.Value == value {
			return card.Suit
		}
	}

	for _, card := range table {
		if card.Value == value {
			return card.Suit
		}
	}

	return ""
}
