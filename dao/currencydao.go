package dao

import (
	"bravo/errorsbravo"
)

var MOCKED_COINS_DB = map[string]float64{"BTC": 1230.123, "BRL": 0.2, "USD": 1}

type MockedCoins struct{}

func (m *MockedCoins) Get(currency string) (float64, error) {
	value, ok := MOCKED_COINS_DB[currency]
	if !ok {
		return 0, errorsbravo.CURRENCY_DOESNT_EXISTS
	}

	return value, nil
}

func (m *MockedCoins) GetAllCurrencies() []string {
	keys := []string{}
	for k := range MOCKED_COINS_DB {
		keys = append(keys, k)
	}
	return keys
}

func (m *MockedCoins) InsertOrUpdate(currency string, value float64) {
	MOCKED_COINS_DB[currency] = value
}

func (m *MockedCoins) Delete(currency string) error {
	_, ok := MOCKED_COINS_DB[currency]
	if !ok {
		return errorsbravo.CURRENCY_DOESNT_EXISTS
	}

	delete(MOCKED_COINS_DB, currency)
	return nil
}

func NewMockedCoins() *MockedCoins {
	return &MockedCoins{}
}
