package dao

var MOCKED_COINS_DB = map[string]float64{"BTC": 1230.123, "BRL": 0.2, "USD": 1}

func GetCoinValues(from string, to string) (float64, float64) {
	var fromValue float64 = GetCoinValue(from)
	var toValue float64 = GetCoinValue(to)
	return fromValue, toValue
}

func GetCoinValue(from string) float64 {
	return MOCKED_COINS_DB[from]
}
