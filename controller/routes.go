package controller

import (
	currency "bravo/controller/currency"
	"bravo/repository"
	"bravo/service"

	"github.com/gin-gonic/gin"
)

func InitializeServerRoutes(database repository.Database) {
	r := gin.Default()
	converter := service.NewConverter(database)
	currencyManager := service.NewCurrencyManager(database)
	currencyController := currency.New(converter, currencyManager)

	r.GET("/convert", currencyController.ConversionHandler)
	r.GET("/currency", currencyController.CurrencyGetHandler)
	r.POST("/currency", currencyController.CurrencyPostHandler)
	r.DELETE("/currency", currencyController.CurrencyDeleteHandler)

	r.Run(":5656")
}
