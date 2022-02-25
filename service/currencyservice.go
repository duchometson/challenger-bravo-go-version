package service

import (
	"bravo/dao"
)

func GetCurrencyValue(currencyName string) float64 {
	currencyValue := dao.GetCoinValue(currencyName)
	return currencyValue
}
