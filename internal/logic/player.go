package logic

type Player struct {
	Username    string `json:"username"`
	Position    int
	Role        string
	Balance     int
	CurrentBet  int
	InGame      bool
	Cards       []Card
	Combination interface{}
	Kicker      Card
}

func NewPlayer(username string, balance int) Player {
	return Player{
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

	combinations := [9]func([5]Card, *bool){
		p.FlushRoyalCheck,
		p.StraightFlushCheck,
		p.FourOfAKindCheck,
		p.FullHouseCheck,
		p.FlushCheck,
		p.StraightCheck,
		p.SetCheck,
		p.TwoPairCheck,
		p.PairCheck,
	}

	for i := 0; i < len(combinations); i++ {
		combinations[i](table, &combinationFound)

		if combinationFound {
			break
		}
	}

	if !combinationFound {
		p.Combination = HighCard{HighCard: p.Kicker}
	}
}
