package service

import (
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

func (cM *CurrencyManager) InsertOrUpdate(currency string, value float64) {
	cM.database.InsertOrUpdate(currency, value)
}

func (cM *CurrencyManager) Delete(currency string) error {
	err := cM.database.Delete(currency)
	if err != nil {
		return model.NewApplicationError(err, currency)
	}
	return nil
}

func NewCurrencyManager(database Database) *CurrencyManager {
	return &CurrencyManager{
		database: database,
	}
}
