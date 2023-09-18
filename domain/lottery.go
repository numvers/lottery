package domain

import "slices"

type Lottery struct {
	Round      uint   `json:"round"`
	PickedDate string `json:"date"`
	Numbers    []uint `json:"numbers"`
	Wins       []Win  `json:"wins"`
}

func (l Lottery) CotainsAll(numbers ...uint) bool {
	for _, num := range numbers {
		_, found := slices.BinarySearch(l.Numbers, num)
		if !found {
			return false
		}
	}
	return true
}

type Win struct {
	NumWinners uint `json:"num_winners"`
	Prize      uint `json:"prize"`
}

type LotteryRepoitoy interface {
	FindAll() ([]Lottery, error)
	FindByRound(round uint) (Lottery, error)
	FindAllByNumbers(numbers ...uint) ([]Lottery, error)
}
