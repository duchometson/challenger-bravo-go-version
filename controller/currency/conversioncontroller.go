package currency

import (
	"bravo/errorsbravo"
	models "bravo/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Currency struct {
	converter Converter
}

func (c *Currency) ConversionHandler(ctx *gin.Context) {
	from := ctx.Query("from")
	if from == "" {
		ctx.JSON(http.StatusBadRequest, errorsbravo.MISSING_PARAM.Error())
		return
	}

	to := ctx.Query("to")
	if to == "" {
		ctx.JSON(http.StatusBadRequest, errorsbravo.MISSING_PARAM.Error())
		return
	}

	value := ctx.Query("value")
	if value == "" {
		ctx.JSON(http.StatusBadRequest, errorsbravo.INVALID_VALUE_PARAM.Error())
		return
	}

	parsedValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	result, err := c.converter.Convert(from, to, parsedValue)
	if err != nil {
		err, ok := err.(*models.ApplicationError)
		if ok {
			if err.Err == errorsbravo.CURRENCY_DOESNT_EXISTS {
				ctx.JSON(http.StatusNotFound, err.Error())
				return
			}
		}

		ctx.JSON(http.StatusInternalServerError, errorsbravo.CURRENCY_DOESNT_EXISTS.Error())
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func New(converter Converter) *Currency {
	return &Currency{
		converter: converter,
	}
}
