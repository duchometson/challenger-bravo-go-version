package currencyapi

import "github.com/asvvvad/exchange"

type CurrencyAPI struct {
	repository Repository
}

func (c *CurrencyAPI) Get(currency string) (float64, error) {
	err := exchange.ValidateCode(currency)
	if err != nil {
		return 0, err
	}

	currencyToUpdate := exchange.New(currency)
	value, err := currencyToUpdate.ConvertTo("USD", 1)
	defer func() {
		if recover() != nil {
			err = exchange.ErrInvalidCode
		}
		err = nil
	}()
	if err != nil {
		return 0, exchange.ErrInvalidCode
	}
	valueAsFloat, _ := value.Float64()
	return valueAsFloat, nil
}

func New() *CurrencyAPI {
	return &CurrencyAPI{}
}
