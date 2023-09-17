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
}
