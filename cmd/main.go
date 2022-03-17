package main

import (
	"bravo/application/currency"
	domainCurrency "bravo/domain/currency"
	"bravo/externalservice/apilayer"
	"bravo/externalservice/currencyapi"
	"bravo/httpserver"
	"bravo/infrastructure/redis"
	"bravo/repository"
	"context"
	"log"
	"time"

	_ "github.com/golang/mock/mockgen/model"
)

func main() {
	database := redis.New("localhost:6379", "", 0)

	repo := repository.New(database)

	currencyService := domainCurrency.New(repo)

	currencyApiLayer := apilayer.New("USD")

	currencyApi := currencyapi.New(currencyApiLayer)

	worker := currency.NewWorker(currencyService, currencyApi, 5*time.Minute, []string{"BTC", "BRL", "EUR", "USD"})

	go worker.Update()

	currency := currency.New(currencyService)

	httpServer := httpserver.New(5656)

	defer httpServer.Shutdown(context.Background())

	if err := httpServer.ListenAndServe(currency); err != nil {
		log.Fatal(err)
	}
}
