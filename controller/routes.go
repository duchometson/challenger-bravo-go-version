package controller

import (
	"bravo/controller/currency"
	"bravo/dao"
	"bravo/service"

	"github.com/gin-gonic/gin"
)

func InitializeServerRoutes() {
	r := gin.Default()

	database := dao.NewMockedCoins()
	converter := service.NewConverter(database)
	currencyController := currency.New(converter)

	r.GET("/convert", currencyController.ConversionHandler)
	r.Run(":5656")
}
