package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/numvers/lottery/domain"
)

func getLotteries(repo domain.LotteryRepoitoy) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		strIncludes := r.URL.Query()["include"]
		includes := make([]uint, len(strIncludes))
		for i, s := range strIncludes {
			n, err := strconv.Atoi(s)
			if err != nil {
				writeHttpErrorResponse(w, err)
				return
			}
			includes[i] = uint(n)
		}
		lotteries, err := repo.FindAllByNumbers(includes...)
		if err != nil {
			writeHttpErrorResponse(w, err)
			return
		}
		writeHttpJsonResponse(w, 200, lotteries)
	}
}

func getLotteriesByRound(repo domain.LotteryRepoitoy) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		strRound := chi.URLParam(r, "round")
		round, err := strconv.Atoi(strRound)
		if err != nil {
			writeHttpErrorResponse(w, err)
			return
		}
		lottery, err := repo.FindByRound(uint(round))
		if err != nil {
			writeHttpErrorResponse(w, err)
			return
		}
		writeHttpJsonResponse(w, 200, lottery)
	}
}

func getStatsWinByNumber(service domain.LotteryStatsService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		stats, err := service.StatsWinByNumber()
		if err != nil {
			writeHttpErrorResponse(w, err)
			return
		}
		writeHttpJsonResponse(w, 200, stats)
	}
}

func writeHttpErrorResponse(w http.ResponseWriter, err error) {
	writeHttpJsonResponse(w, 500, domain.ErrorResponse{Code: 500, Message: err.Error()})
}

func writeHttpJsonResponse[R jsonResponse](w http.ResponseWriter, code int, cotent R) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(cotent)
}

type jsonResponse interface {
	[]domain.Lottery | domain.Lottery | []domain.WinByNumber | domain.ErrorResponse
}
