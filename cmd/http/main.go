package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/numvers/lottery/domain"
	"github.com/numvers/lottery/lottery/repository/sqlite"
)

func main() {
	var addr string
	flag.StringVar(&addr, "db_path", "", "sqlite db file path")
	flag.Parse()

	if addr == "" {
		log.Fatal("db path is empty")
	}

	db, err := sql.Open("sqlite", addr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repository := sqlite.NewLotteryRepository(db)
	statService := domain.NewLotteryStatsService(repository)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Get("/lotteries", getLotteries(repository))
	router.Get("/lotteries/{round}", getLotteriesByRound(repository))
	router.Get("/stats/win_by_number", getStatsWinByNumber(statService))
	port := ":8080"
	server := &http.Server{
		Addr:    port,
		Handler: router,
	}
	go func() {
		log.Print("Starting server using port " + port + "\n")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v \n", err)
		}
	}()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-exit
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("Server Shutdown error: %v\n", err)
	}
	log.Print("Server Shutdon")
}
