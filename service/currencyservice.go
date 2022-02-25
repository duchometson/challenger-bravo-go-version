package service

import (
	"bravo/dao"
)

func GetCurrencyValue(currencyName string) float64 {
	currencyValue, ok := dao.GetCoinValue(currencyName)
	validateCurrencyExistance(ok)
	return currencyValue
}

func validateCurrencyExistance(ok bool) {
	if !ok {
		panic(ok)
	}
}
