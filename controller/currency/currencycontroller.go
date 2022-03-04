package currency

import (
	"bravo/errorsbravo"
	models "bravo/model"
	"net/http"

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

		ctx.JSON(http.StatusInternalServerError, errorsbravo.CURRENCY_DOESNT_EXISTS.Error())
		return
	}

	ctx.JSON(http.StatusOK, value)
}
