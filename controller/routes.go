package controller

import (
	currency "bravo/controller/currency"
	"bravo/dao"
	"bravo/service"

	"github.com/gin-gonic/gin"
)

func InitializeServerRoutes() {
	r := gin.Default()

	database := dao.NewMockedCoins()
	converter := service.NewConverter(database)
	currencyManager := service.NewCurrencyManager(database)
	currencyController := currency.New(converter, currencyManager)

	r.GET("/convert", currencyController.ConversionHandler)
	r.GET("/currency", currencyController.CurrencyGetHandler)
	r.POST("/currency", currencyController.CurrencyPostHandler)

	r.Run(":5656")
}
