package redis

import (
	"bravo/errorsbravo"
)

var MOCKED_COINS_DB = map[string]float64{"BTC": 1230.123, "BRL": 0.2, "USD": 1}

type Client struct{}

// TO IMPLEMENT USING REDIS:
//Get(string) (interface{}, error)
//Set(string, interface{}) error
//Delete(string) error
//ErrorNotFound() error

// https://github.com/go-redis/redis

func (m *Client) Get(currency string) (float64, error) {
	value, ok := MOCKED_COINS_DB[currency]
	if !ok {
		return 0, errorsbravo.CURRENCY_DOESNT_EXISTS
	}

	return value, nil
}

func (m *Client) InsertOrUpdate(currency string, value float64) {
	MOCKED_COINS_DB[currency] = value
}

func (m *Client) Delete(currency string) error {
	_, ok := MOCKED_COINS_DB[currency]
	if !ok {
		return errorsbravo.CURRENCY_DOESNT_EXISTS
	}

	delete(MOCKED_COINS_DB, currency)
	return nil
}

func New() *Client {
	return &Client{}
}
