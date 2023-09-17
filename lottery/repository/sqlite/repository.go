package sqlite

import (
	"database/sql"
	"errors"

	"github.com/numvers/lottery/domain"
	_ "modernc.org/sqlite"
)

type LotteryRepoitoy struct {
	db *sql.DB
}

func NewLotteryRepository(db *sql.DB) *LotteryRepoitoy {
	return &LotteryRepoitoy{db}
}

func (r *LotteryRepoitoy) FindAll() ([]domain.Lottery, error) {
	rows, err := r.db.Query("SELECT * FROM lotteries")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []domain.Lottery{}
	for rows.Next() {
		var r row
		err := rows.Scan(&r.round, &r.picked_date,
			&r.num_first_winners, &r.first_prize,
			&r.num_second_winners, &r.second_prize,
			&r.num_third_winners, &r.third_prize,
			&r.num_forth_winners, &r.forth_prize,
			&r.num_fifth_winners, &r.fifth_prize,
			&r.first_number, &r.second_number, &r.third_number, &r.forth_number, &r.fifth_number, &r.sixth_number, &r.bonus_number)
		if err != nil {
			return nil, err
		}
		results = append(results, r.toLottery())
	}
	return results, nil
}

func (r *LotteryRepoitoy) FindByRound(round uint) (domain.Lottery, error) {
	return domain.Lottery{}, errors.New("not implemented") // TODO: Implement
}

type row struct {
	round              int
	picked_date        string
	num_first_winners  int
	first_prize        int
	num_second_winners int
	second_prize       int
	num_third_winners  int
	third_prize        int
	num_forth_winners  int
	forth_prize        int
	num_fifth_winners  int
	fifth_prize        int
	first_number       int8
	second_number      int8
	third_number       int8
	forth_number       int8
	fifth_number       int8
	sixth_number       int8
	bonus_number       int8
}

func (r *row) toLottery() domain.Lottery {
	return domain.Lottery{
		Round:      uint(r.round),
		PickedDate: r.picked_date,
		Numbers: domain.LotteryNumbers{
			Numbers: []uint{uint(r.first_number), uint(r.second_number), uint(r.third_number), uint(r.forth_number), uint(r.fifth_number), uint(r.bonus_number)},
		},
		Wins: []domain.Win{
			domain.Win{
				NumWinners: uint(r.num_first_winners),
				Prize:      uint(r.first_prize),
			},
			domain.Win{
				NumWinners: uint(r.num_second_winners),
				Prize:      uint(r.second_prize),
			},
			domain.Win{
				NumWinners: uint(r.num_third_winners),
				Prize:      uint(r.third_prize),
			},
			domain.Win{
				NumWinners: uint(r.num_forth_winners),
				Prize:      uint(r.forth_prize),
			},
			domain.Win{
				NumWinners: uint(r.num_fifth_winners),
				Prize:      uint(r.fifth_prize),
			},
		},
	}
}
