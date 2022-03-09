package repository

import (
	"bravo/errorsbravo"
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

	parsedValue := value.(float64)

	return parsedValue, nil
}

// IMPLEMENT DELETE AND SET

//func (c *Currency) InsertOrUpdate(currency string, value float64) {
//	MOCKED_COINS_DB[currency] = value
//}
//
//func (c *Currency) Delete(currency string) error {
//	_, ok := MOCKED_COINS_DB[currency]
//	if !ok {
//		return errorsbravo.CURRENCY_DOESNT_EXISTS
//	}
//
//	delete(MOCKED_COINS_DB, currency)
//	return nil
//}

func New(database Database) *Currency {
	return &Currency{
		database: database,
	}
}
