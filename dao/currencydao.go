package dao

var MOCKED_COINS_DB = map[string]float64{"BTC": 1230.123, "BRL": 0.2, "USD": 1}

func GetCoinValues(from string, to string) (float64, float64) {
	fromValue, ok := MOCKED_COINS_DB[from]
	validateCurrencyExistance(ok)
	toValue, ok := MOCKED_COINS_DB[to]
	validateCurrencyExistance(ok)
	return fromValue, toValue
}

func GetCoinValue(from string) float64 {
	value, ok := MOCKED_COINS_DB[from]
	validateCurrencyExistance(ok)
	return value
}

func validateCurrencyExistance(ok bool) {
	if !ok {
		panic(ok)
	}
}
