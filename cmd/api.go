package main

import (
	"bravo/application/currency"
	domainCurrency "bravo/domain/currency"
	"bravo/externalservice/currencyapi"
	"bravo/httpserver"
	"bravo/infrastructure/redis"
	"bravo/repository"
	"context"
	"log"
)

func main() {
	database := redis.New()

	repo := repository.New(database)

	currencyService := domainCurrency.New(repo)

	currencyApi := currencyapi.New()

	worker := currency.NewWorker(currencyService, currencyApi)

	go worker.Update()

	currency.New(currencyService)

	httpServer := httpserver.New(5656)

	defer httpServer.Shutdown(context.Background())

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
