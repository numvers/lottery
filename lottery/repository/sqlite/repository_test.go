package sqlite

import (
	"database/sql"
	"testing"

	"github.com/assertg/assert"
)

func TestLotteryRepository(t *testing.T) {

	db, err := sql.Open("sqlite", "./testdata/lottery.db")
	assert.NoError(t, err)

	err = db.Ping()
	assert.NoError(t, err)
	t.Cleanup(func() {
		db.Close()
	})

	repository := NewLotteryRepository(db)

	t.Run("FindAll", func(t *testing.T) {
		lotteries, err := repository.FindAll()
		assert.NoError(t, err)
		assert.Positive(t, len(lotteries))
	})

	t.Run("FindByRound", func(t *testing.T) {
		for _, round := range []uint{1, 99} {
			lottery, err := repository.FindByRound(round)
			assert.NoError(t, err)
			assert.Equals(t, lottery.Round, round)
		}

		lottery, err := repository.FindByRound(9999)
		if err == nil {
			t.Error("expected error but got nil")
		}
		assert.Equals(t, lottery.Round, 0)

		// return latest lottery when round is 0
		lottery, err = repository.FindByRound(0)
		assert.NoError(t, err)
		if lottery.Round < 1080 {
			t.Errorf("expected round larger than 1080 but got %v", lottery)
		}
	})

	t.Run("FindAllByNumbers", func(t *testing.T) {
		for _, number := range []uint{10, 20, 30, 40} {
			lotteries, err := repository.FindAllByNumbers(number)
			assert.NoError(t, err)
			assert.Positive(t, len(lotteries))
		}

		lotteries, err := repository.FindAllByNumbers(10, 20, 30)
		assert.NoError(t, err)
		assert.Positive(t, len(lotteries))

		lotteries, err = repository.FindAllByNumbers(0, 10)
		assert.NoError(t, err)
		assert.Equals(t, len(lotteries), 0)
	})

}
