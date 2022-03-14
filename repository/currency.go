package repository

import (
	"bravo/errorsbravo"
	"strconv"
)

var MOCKED_COINS_DB = map[string]float64{"BTC": 1230.123, "BRL": 0.2, "USD": 1}

type Currency struct {
	database Database
}

func (c *Currency) Get(currency string) (float64, error) {
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

func (c *Currency) GetAllKeys() ([]string, error) {
	allKeys, err := c.database.GetAllKeys()
	if err != nil {
		if err == c.database.ErrorNotFound() {
			return []string{}, errorsbravo.CURRENCY_DOESNT_EXISTS
		}

		return []string{}, err
	}
	return allKeys, nil
}

func (c *Currency) Set(currency string, value float64) error {
	err := c.database.Set(currency, value)
	if err != nil {
		return errorsbravo.INTERNAL_ERROR
	}
	return nil
}

func (c *Currency) Delete(currency string) error {
	err := c.database.Delete(currency)
	if err != nil {
		if err == c.database.ErrorNotFound() {
			return errorsbravo.CURRENCY_DOESNT_EXISTS
		}
		return err
	}
	return nil
}

func New(database Database) *Currency {
	return &Currency{
		database: database,
	}
}
