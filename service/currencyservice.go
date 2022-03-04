package service

import (
	"bravo/dao"
	"bravo/model"
)

type CurrencyManager struct {
	database Database
}

func (cM *CurrencyManager) Get(currency string) (float64, error) {
	currencyValue, err := cM.database.Get(currency)
	if err != nil {
		return 0, model.NewApplicationError(err, currency)
	}
	return currencyValue, nil
}

func GetCurrencyValue(currencyName string) float64 {
	currencyValue := dao.GetCoinValue(currencyName)
	return currencyValue
}

func NewCurrencyManager(database Database) *CurrencyManager {
	return &CurrencyManager{
		database: database,
	}
}
