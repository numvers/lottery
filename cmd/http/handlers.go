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
		lotteries, err := repo.FindAll()
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

func writeHttpErrorResponse(w http.ResponseWriter, err error) {
	writeHttpJsonResponse(w, 500, domain.ErrorResponse{Code: 500, Message: err.Error()})
}

func writeHttpJsonResponse[R jsonResponse](w http.ResponseWriter, code int, cotent R) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(cotent)
}

type jsonResponse interface {
	[]domain.Lottery | domain.Lottery | domain.ErrorResponse
}
