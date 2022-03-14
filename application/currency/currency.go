package currency

import (
	"bravo/errorsbravo"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Currency struct {
	service Service
}

func (c *Currency) Get(ctx *gin.Context) {
	name, ok := ctx.GetQuery("name")
	if !ok {
		ctx.JSON(http.StatusBadRequest, errorsbravo.MISSING_PARAM)
		return
	}

	value, err := c.service.Get(name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, value)
	return
}

func (c *Currency) Add(ctx *gin.Context) {
	name, ok := ctx.GetQuery("name")
	if !ok {
		ctx.JSON(http.StatusBadRequest, errorsbravo.MISSING_PARAM)
		return
	}

	amount, ok := ctx.GetQuery("amount")
	if !ok {
		ctx.JSON(http.StatusBadRequest, errorsbravo.MISSING_PARAM)
		return
	}

	parsedAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorsbravo.INVALID_VALUE_PARAM)
		return
	}

	err = c.service.Set(name, parsedAmount)
	if err != nil {
		fmt.Println("não entendi po")
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("entao foi?")

	ctx.JSON(http.StatusOK, "currency set")
	return
}

func (c *Currency) Delete(ctx *gin.Context) {
	name, ok := ctx.GetQuery("name")
	if !ok {
		ctx.JSON(http.StatusBadRequest, errorsbravo.MISSING_PARAM)
		return
	}

	err := c.service.Delete(name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "currency deleted")
	return
}

func (c *Currency) Convert(ctx *gin.Context) {
	from, ok := ctx.GetQuery("from")
	if !ok {
		ctx.JSON(http.StatusBadRequest, errorsbravo.MISSING_PARAM)
		return
	}

	to, ok := ctx.GetQuery("to")
	if !ok {
		ctx.JSON(http.StatusBadRequest, errorsbravo.MISSING_PARAM)
		return
	}

	amount, ok := ctx.GetQuery("amount")
	if !ok {
		ctx.JSON(http.StatusBadRequest, errorsbravo.MISSING_PARAM)
		return
	}

	parsedAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorsbravo.INVALID_VALUE_PARAM)
		return
	}

	convertedValue, err := c.service.Convert(from, to, parsedAmount)
	ctx.JSON(http.StatusOK, convertedValue)
	return
}

func New(service Service) *Currency {
	return &Currency{
		service: service,
	}
}
