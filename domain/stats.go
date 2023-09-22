package domain

type LotteryStatsService struct {
	repo WinnerLotteryRepoitoy
}

type WinByNumber struct {
	Number  uint `json:"number"`
	NumWins uint `json:"num_wins"`
}

func NewLotteryStatsService(repo WinnerLotteryRepoitoy) LotteryStatsService {
	return LotteryStatsService{repo: repo}
}

func (s *LotteryStatsService) StatsWinByNumber() ([]WinByNumber, error) {
	lotteries, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	numWins := [45]uint{}
	for _, lottery := range lotteries {
		for _, num := range lottery.Numbers {
			numWins[num-1] += 1
		}
	}

	stats := make([]WinByNumber, 45)
	for i, win := range numWins {
		stats[i] = WinByNumber{Number: uint(i + 1), NumWins: win}
	}
	return stats, nil
}
