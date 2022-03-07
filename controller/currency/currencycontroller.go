package currency

import (
	"bravo/errorsbravo"
	models "bravo/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *Currency) CurrencyGetHandler(ctx *gin.Context) {
	name := ctx.Query("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, errorsbravo.MISSING_PARAM.Error())
		return
	}
	value, err := c.currencyManager.Get(name)

	if err != nil {
		err, ok := err.(*models.ApplicationError)
		if ok {
			if err.Err == errorsbravo.CURRENCY_DOESNT_EXISTS {
				ctx.JSON(http.StatusNotFound, err.Error())
				return
			}
		}

		ctx.JSON(http.StatusInternalServerError, errorsbravo.INTERNAL_ERROR.Error())
		return
	}

	ctx.JSON(http.StatusOK, value)
}

func (c *Currency) CurrencyPostHandler(ctx *gin.Context) {
	name := ctx.Query("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, errorsbravo.MISSING_PARAM.Error())
		return
	}

	value := ctx.Query("value")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, errorsbravo.MISSING_PARAM.Error())
		return
	}

	parsedValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.currencyManager.Insert(name, parsedValue)

	if err != nil {
		err, ok := err.(*models.ApplicationError)
		if ok {
			if err.Err == errorsbravo.CURRENCY_DOESNT_EXISTS {
				ctx.JSON(http.StatusNotFound, err.Error())
				return
			}
		}

		ctx.JSON(http.StatusInternalServerError, errorsbravo.INTERNAL_ERROR.Error())
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (c *Currency) CurrencyDeleteHandler(ctx *gin.Context) {
	name := ctx.Query("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, errorsbravo.MISSING_PARAM.Error())
		return
	}

	err := c.currencyManager.Delete(name)

	if err != nil {
		err, ok := err.(*models.ApplicationError)
		if ok {
			if err.Err == errorsbravo.CURRENCY_DOESNT_EXISTS {
				ctx.JSON(http.StatusNotFound, err.Error())
				return
			}
		}

		ctx.JSON(http.StatusInternalServerError, errorsbravo.INTERNAL_ERROR.Error())
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
