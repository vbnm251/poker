package logic_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"poker/internal/logic"
	"testing"
)

func TestTestCombination(t *testing.T) {
	type Tests struct {
		Number   int
		Player   logic.Player
		Table    [5]logic.Card
		Expected interface{}
	}

	cases := []Tests{
		//two pair 5 and 7
		{
			Number: 1,
			Player: logic.Player{
				Cards: []logic.Card{
					{
						Value: 14,
						Suit:  logic.Hearts,
					},
					{
						Value: 5,
						Suit:  logic.Clubs,
					},
				},
			},
			Table: [5]logic.Card{
				{
					Value: 12,
					Suit:  logic.Diamonds,
				},
				{
					Value: 7,
					Suit:  logic.Spades,
				},
				{
					Value: 9,
					Suit:  logic.Diamonds,
				},
				{
					Value: 5,
					Suit:  logic.Diamonds,
				},
				{
					Value: 7,
					Suit:  logic.Clubs,
				},
			},
			Expected: logic.TwoPair{
				FirsPair: logic.Card{
					Value: 7,
					Suit:  logic.Spades,
				},
				SecondPair: logic.Card{
					Value: 5,
					Suit:  logic.Clubs,
				},
			},
		},
		//flush royal hearts
		{
			Number: 2,
			Player: logic.Player{
				Cards: []logic.Card{
					{
						Value: 14,
						Suit:  logic.Hearts,
					},
					{
						Value: 12,
						Suit:  logic.Hearts,
					},
				},
			},
			Table: [5]logic.Card{
				{
					Value: 13,
					Suit:  logic.Hearts,
				},
				{
					Value: 7,
					Suit:  logic.Spades,
				},
				{
					Value: 10,
					Suit:  logic.Hearts,
				},
				{
					Value: 5,
					Suit:  logic.Diamonds,
				},
				{
					Value: 11,
					Suit:  logic.Hearts,
				},
			},
			Expected: logic.FlushRoyal{HighCard: logic.Card{
				Value: 14,
				Suit:  logic.Hearts,
			}},
		},
		// Test case 1: flush royal in spades
		{
			Number: 3,
			Player: logic.Player{
				Cards: []logic.Card{
					{
						Value: 14,
						Suit:  logic.Spades,
					},
					{
						Value: 13,
						Suit:  logic.Spades,
					},
				},
			},
			Table: [5]logic.Card{
				{
					Value: 12,
					Suit:  logic.Spades,
				},
				{
					Value: 11,
					Suit:  logic.Spades,
				},
				{
					Value: 10,
					Suit:  logic.Spades,
				},
				{
					Value: 9,
					Suit:  logic.Spades,
				},
				{
					Value: 8,
					Suit:  logic.Spades,
				},
			},
			Expected: logic.FlushRoyal{HighCard: logic.Card{
				Value: 14,
				Suit:  logic.Spades,
			}},
		},
		// Test case 2: no flush royal, only straight flush 12 hearts
		{
			Number: 4,
			Player: logic.Player{
				Cards: []logic.Card{
					{
						Value: 12,
						Suit:  logic.Hearts,
					},
					{
						Value: 11,
						Suit:  logic.Hearts,
					},
				},
			},
			Table: [5]logic.Card{
				{
					Value: 10,
					Suit:  logic.Hearts,
				},
				{
					Value: 9,
					Suit:  logic.Hearts,
				},
				{
					Value: 8,
					Suit:  logic.Hearts,
				},
				{
					Value: 7,
					Suit:  logic.Hearts,
				},
				{
					Value: 6,
					Suit:  logic.Hearts,
				},
			},
			Expected: logic.StraightFlush{HighCard: logic.Card{
				Value: 12,
				Suit:  logic.Hearts,
			}},
		},
		// high card
		{
			Number: 5,
			Player: logic.Player{
				Cards: []logic.Card{
					{
						Value: 14,
						Suit:  logic.Hearts,
					},
					{
						Value: 12,
						Suit:  logic.Hearts,
					},
				},
			},
			Table: [5]logic.Card{
				{
					Value: 11,
					Suit:  logic.Spades,
				},
				{
					Value: 10,
					Suit:  logic.Clubs,
				},
				{
					Value: 8,
					Suit:  logic.Hearts,
				},
				{
					Value: 5,
					Suit:  logic.Diamonds,
				},
				{
					Value: 2,
					Suit:  logic.Spades,
				},
			},
			Expected: logic.HighCard{HighCard: logic.Card{
				Value: 14,
				Suit:  logic.Hearts,
			}},
		},
		//set 14
		{
			Number: 6,
			Player: logic.Player{
				Cards: []logic.Card{
					{
						Value: 14,
						Suit:  logic.Hearts,
					},
					{
						Value: 14,
						Suit:  logic.Spades,
					},
				},
			},
			Table: [5]logic.Card{
				{
					Value: 13,
					Suit:  logic.Hearts,
				},
				{
					Value: 7,
					Suit:  logic.Spades,
				},
				{
					Value: 10,
					Suit:  logic.Hearts,
				},
				{
					Value: 5,
					Suit:  logic.Diamonds,
				},
				{
					Value: 14,
					Suit:  logic.Clubs,
				},
			},
			Expected: logic.Set{HighCard: logic.Card{
				Value: 14,
				Suit:  logic.Hearts,
			}},
		},
		//care 14
		{
			Number: 7,
			Player: logic.Player{
				Cards: []logic.Card{
					{
						Value: 14,
						Suit:  logic.Hearts,
					},
					{
						Value: 14,
						Suit:  logic.Spades,
					},
				},
			},
			Table: [5]logic.Card{
				{
					Value: 13,
					Suit:  logic.Hearts,
				},
				{
					Value: 7,
					Suit:  logic.Spades,
				},
				{
					Value: 14,
					Suit:  logic.Diamonds,
				},
				{
					Value: 5,
					Suit:  logic.Diamonds,
				},
				{
					Value: 14,
					Suit:  logic.Clubs,
				},
			},
			Expected: logic.FourOfAKind{HighCard: logic.Card{
				Value: 14,
				Suit:  logic.Hearts,
			},
			},
		},
		//straight 10
		{
			Number: 8,
			Player: logic.Player{
				Cards: []logic.Card{
					{
						Value: 9,
						Suit:  logic.Hearts,
					},
					{
						Value: 7,
						Suit:  logic.Spades,
					},
				},
			},
			Table: [5]logic.Card{
				{
					Value: 13,
					Suit:  logic.Hearts,
				},
				{
					Value: 7,
					Suit:  logic.Hearts,
				},
				{
					Value: 10,
					Suit:  logic.Hearts,
				},
				{
					Value: 8,
					Suit:  logic.Diamonds,
				},
				{
					Value: 6,
					Suit:  logic.Clubs,
				},
			},
			Expected: logic.Straight{
				HighCard: logic.Card{
					Value: 10,
					Suit:  logic.Hearts,
				},
			},
		},
		//flush royal hearts
		{
			Number: 9,
			Player: logic.Player{
				Cards: []logic.Card{
					{
						Value: 14,
						Suit:  logic.Hearts,
					},
					{
						Value: 12,
						Suit:  logic.Hearts,
					},
				},
			},
			Table: [5]logic.Card{
				{
					Value: 13,
					Suit:  logic.Hearts,
				},
				{
					Value: 13,
					Suit:  logic.Spades,
				},
				{
					Value: 10,
					Suit:  logic.Hearts,
				},
				{
					Value: 5,
					Suit:  logic.Diamonds,
				},
				{
					Value: 11,
					Suit:  logic.Hearts,
				},
			},
			Expected: logic.FlushRoyal{
				HighCard: logic.Card{
					Value: 14,
					Suit:  logic.Hearts,
				},
			},
		},
		//triple 14
		{
			Number: 10,
			Player: logic.Player{
				Cards: []logic.Card{
					{
						Value: 14,
						Suit:  logic.Hearts,
					},
					{
						Value: 14,
						Suit:  logic.Diamonds,
					},
				},
			},
			Table: [5]logic.Card{
				{
					Value: 13,
					Suit:  logic.Hearts,
				},
				{
					Value: 11,
					Suit:  logic.Spades,
				},
				{
					Value: 10,
					Suit:  logic.Hearts,
				},
				{
					Value: 5,
					Suit:  logic.Diamonds,
				},
				{
					Value: 14,
					Suit:  logic.Clubs,
				},
			},
			Expected: logic.Set{HighCard: logic.Card{
				Value: 14,
				Suit:  logic.Hearts,
			}},
		},
		//pair 10
		{
			Number: 11,
			Player: logic.Player{
				Cards: []logic.Card{
					{
						Value: 10,
						Suit:  logic.Spades,
					},
					{
						Value: 4,
						Suit:  logic.Hearts,
					},
				},
			},
			Table: [5]logic.Card{
				{
					Value: 14,
					Suit:  logic.Hearts,
				},
				{
					Value: 13,
					Suit:  logic.Spades,
				},
				{
					Value: 10,
					Suit:  logic.Clubs,
				},
				{
					Value: 9,
					Suit:  logic.Diamonds,
				},
				{
					Value: 8,
					Suit:  logic.Hearts,
				},
			},
			Expected: logic.Pair{
				HighCard: logic.Card{
					Value: 10,
					Suit:  logic.Spades,
				},
			},
		},
		// Test case: Straight - 5 cards in sequence of different suits
		{
			Number: 12,
			Player: logic.Player{
				Cards: []logic.Card{
					{
						Value: 10,
						Suit:  logic.Hearts,
					},
					{
						Value: 9,
						Suit:  logic.Clubs,
					},
				},
			},
			Table: [5]logic.Card{
				{
					Value: 14,
					Suit:  logic.Spades,
				},
				{
					Value: 13,
					Suit:  logic.Hearts,
				},
				{
					Value: 12,
					Suit:  logic.Clubs,
				},
				{
					Value: 11,
					Suit:  logic.Diamonds,
				},
				{
					Value: 8,
					Suit:  logic.Spades,
				},
			},
			Expected: logic.Straight{HighCard: logic.Card{
				Value: 14,
				Suit:  logic.Spades,
			}},
		},
		// Test case: Straight - 5 cards in sequence of the same suit
		{
			Number: 13,
			Player: logic.Player{
				Cards: []logic.Card{
					{
						Value: 10,
						Suit:  logic.Hearts,
					},
					{
						Value: 9,
						Suit:  logic.Hearts,
					},
				},
			},
			Table: [5]logic.Card{
				{
					Value: 14,
					Suit:  logic.Hearts,
				},
				{
					Value: 13,
					Suit:  logic.Hearts,
				},
				{
					Value: 12,
					Suit:  logic.Hearts,
				},
				{
					Value: 11,
					Suit:  logic.Hearts,
				},
				{
					Value: 8,
					Suit:  logic.Spades,
				},
			},
			Expected: logic.FlushRoyal{HighCard: logic.Card{
				Value: 14,
				Suit:  logic.Hearts,
			}},
		},
		// Test case: Straight - 6 cards in sequence of different suits
		{
			Number: 14,
			Player: logic.Player{
				Cards: []logic.Card{
					{
						Value: 10,
						Suit:  logic.Hearts,
					},
					{
						Value: 9,
						Suit:  logic.Clubs,
					},
				},
			},
			Table: [5]logic.Card{
				{
					Value: 14,
					Suit:  logic.Spades,
				},
				{
					Value: 13,
					Suit:  logic.Hearts,
				},
				{
					Value: 12,
					Suit:  logic.Clubs,
				},
				{
					Value: 11,
					Suit:  logic.Diamonds,
				},
				{
					Value: 8,
					Suit:  logic.Spades,
				},
			},
			Expected: logic.Straight{HighCard: logic.Card{
				Value: 14,
				Suit:  logic.Spades,
			}},
		},
	}

	for _, test := range cases {
		test.Player.GetCombination(test.Table)
		assert.Equal(t, test.Expected, test.Player.Combination, fmt.Sprintf("test number: %d", test.Number))
	}
}
