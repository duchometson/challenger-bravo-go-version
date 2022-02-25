package dao

var MOCKED_COINS_DB = map[string]float64{"BTC": 1230.123, "BRL": 0.2, "USD": 1}

func GetCoinValues(from string, to string) (float64, float64) {
	var fromValue float64 = MOCKED_COINS_DB[from]
	var toValue float64 = MOCKED_COINS_DB[to]
	return fromValue, toValue
}

func GetCoinValue(from string) (float64, bool) {
	value, ok := MOCKED_COINS_DB[from]
	return value, ok
}
