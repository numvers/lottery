package domain

import "slices"

type WinnerLottery struct {
	Round      uint   `json:"round"`
	PickedDate string `json:"date"`
	Numbers    []uint `json:"numbers"`
	Wins       []Win  `json:"wins"`
}

type Win struct {
	NumWinners uint `json:"num_winners"`
	Prize      uint `json:"prize"`
}

func (l WinnerLottery) CotainsAll(numbers ...uint) bool {
	for _, num := range numbers {
		_, found := slices.BinarySearch(l.Numbers, num)
		if !found {
			return false
		}
	}
	return true
}

type WinnerLotteryRepoitoy interface {
	FindAll() ([]WinnerLottery, error)
	FindByRound(round uint) (WinnerLottery, error)
	FindAllByNumbers(numbers ...uint) ([]WinnerLottery, error)
}
