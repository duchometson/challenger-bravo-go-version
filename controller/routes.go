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
	//r.POST(
	//	r.PUT())
	//http.HandleFunc("/conversion", currency.ConversionHandler)
	//http.HandleFunc("/currency", CurrencyHandler)

	//fmt.Printf("Starting server at port 8080\n")
	//if err := http.ListenAndServe(":8080", nil); err != nil {
	//	log.Fatal(err)
	//}
	r.Run(":5656")
}
