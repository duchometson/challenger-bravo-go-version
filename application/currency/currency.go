package currency

import (
	"bravo/errorsbravo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type Currency struct {
	service Service
}

func (c *Currency) Get(ctx *gin.Context) {
	name, ok := ctx.GetQuery("name")
	if !ok {
		ctx.JSON(http.StatusBadRequest, errorsbravo.MissingParam)
		return
	}

	name = strings.ToUpper(name)

	value, err := c.service.Get(name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, value)
}

func (c *Currency) Add(ctx *gin.Context) {
	name, ok := ctx.GetQuery("name")
	if !ok {
		ctx.JSON(http.StatusBadRequest, errorsbravo.MissingParam)
		return
	}

	name = strings.ToUpper(name)

	value, ok := ctx.GetQuery("value")
	if !ok {
		ctx.JSON(http.StatusBadRequest, errorsbravo.MissingParam)
		return
	}

	parsedValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorsbravo.InvalidValueParam)
		return
	}

	err = c.service.Set(name, parsedValue)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "currency set")
}

func (c *Currency) Delete(ctx *gin.Context) {
	name, ok := ctx.GetQuery("name")
	if !ok {
		ctx.JSON(http.StatusBadRequest, errorsbravo.MissingParam)
		return
	}

	name = strings.ToUpper(name)

	err := c.service.Delete(name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "currency deleted")
}

func (c *Currency) Convert(ctx *gin.Context) {
	from, ok := ctx.GetQuery("from")
	if !ok {
		ctx.JSON(http.StatusBadRequest, errorsbravo.MissingParam)
		return
	}

	to, ok := ctx.GetQuery("to")
	if !ok {
		ctx.JSON(http.StatusBadRequest, errorsbravo.MissingParam)
		return
	}

	amount, ok := ctx.GetQuery("amount")
	if !ok {
		ctx.JSON(http.StatusBadRequest, errorsbravo.MissingParam.Error())
		return
	}

	parsedAmount, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorsbravo.InvalidValueParam)
		return
	}

	convertedValue, err := c.service.Convert(from, to, parsedAmount)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, convertedValue)
}

func New(service Service) *Currency {
	return &Currency{
		service: service,
	}
}
