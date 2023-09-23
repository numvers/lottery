package domain

import (
	"math/rand"
	"slices"
)

type WinnerLottery struct {
	Lottery
	Round uint  `json:"round"`
	Wins  []Win `json:"wins"`
}

type Lottery struct {
	PickedDate string `json:"date"`
	Numbers    []uint `json:"numbers"`
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

type LotteryRepository interface {
	FindAll() ([]Lottery, error)
	Save(lottery Lottery) error
}

type LotteryGeneratorService struct {
	repo LotteryRepository
}

func NewLotteryGeneratorService(repo LotteryRepository) LotteryGeneratorService {
	return LotteryGeneratorService{repo: repo}
}

func (s LotteryGeneratorService) NewNumber() {
	rand.Int31n(45)
}
