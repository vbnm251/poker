package logic

import "github.com/gorilla/websocket"

type Player struct {
	Username    string `mapstructure:"username"`
	Conn        *websocket.Conn
	Position    int
	Role        string
	Balance     int
	CurrentBet  int
	InGame      bool
	Cards       []Card
	Combination interface{}
	Kicker      Card
}

func NewPlayer(username string, balance int, conn *websocket.Conn) Player {
	return Player{
		Conn:     conn,
		Username: username,
		Balance:  balance,
		InGame:   false,
		Cards:    make([]Card, 2),
	}
}

func (p *Player) Call(sum int) {
	p.Balance -= sum
	p.CurrentBet += sum
}

func (p *Player) Raise(sum int) {
	p.Balance -= sum
	p.CurrentBet += sum
}

func (p *Player) Fold() {
	p.InGame = false
	p.CurrentBet = 0
}

func (p *Player) GetCombination(table [5]Card) {
	p.FindKicker(table)

	combinationFound := false

	combinations := [10]func([5]Card, *bool){
		p.FlushRoyalCheck,
		p.StraightFlushCheck,
		p.FourOfAKindCheck,
		p.FullHouseCheck,
		p.FlushCheck,
		p.StraightCheck,
		p.SetCheck,
		p.TwoPairCheck,
		p.PairCheck,
		p.HighCardCheck,
	}

	for i := 0; i < len(combinations); i++ {
		combinations[i](table, &combinationFound)

		if combinationFound {
			break
		}
	}
}
