package repository

import (
	"bravo/errorsbravo"
	"strconv"
)

type CurrencyApi struct {
	database Database
}

func (c *CurrencyApi) Get(currency string) (float64, error) {
	value, err := c.database.Get(currency)
	if err != nil {
		if err == c.database.ErrorNotFound() {
			return 0, errorsbravo.CURRENCY_DOESNT_EXISTS
		}

		return 0, err
	}

	parsedValue, err := strconv.ParseFloat(value.(string), 64)
	if err != nil {
		return 0, errorsbravo.INTERNAL_ERROR
	}

	return parsedValue, nil
}
