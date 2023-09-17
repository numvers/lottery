package domain

type Lottery struct {
	Round      uint   `json:"round"`
	PickedDate string `json:"date"`
	Numbers    []uint `json:"numbers"`
	Wins       []Win  `json:"wins"`
}

type Win struct {
	NumWinners uint `json:"num_winners"`
	Prize      uint `json:"prize"`
}

type LotteryRepoitoy interface {
	FindAll() ([]Lottery, error)
	FindByRound(round uint) (Lottery, error)
}
