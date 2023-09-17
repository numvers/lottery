package domain

type Lottery struct {
	Round      uint
	PickedDate string
	Numbers    LotteryNumbers
	Wins       []Win
}

type LotteryNumbers struct {
	Numbers []uint
}

type Win struct {
	NumWinners uint
	Prize      uint
}

type LotteryRepoitoy interface {
	FindAll() ([]Lottery, error)
	FindByRound(round uint) (Lottery, error)
}
