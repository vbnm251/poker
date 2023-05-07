package logic

type FlushRoyal struct {
	HighCard Card
}

type StraightFlush struct {
	HighCard Card
}

type FourOfAKind struct {
	HighCard Card
}

type FullHouse struct {
	HighTriple Card
	HighPair   Card
}

type Flush struct {
	HighCard Card
}

type Straight struct {
	HighCard Card
}

type Set struct {
	HighCard Card
}

type TwoPair struct {
	FirsPair   Card
	SecondPair Card
}

type Pair struct {
	HighCard Card
}

type HighCard struct {
	HighCard Card
}
