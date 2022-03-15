package currencyapi

import "github.com/asvvvad/exchange"

type CurrencyAPI struct {
	repository Repository
}

const SERVER_CURRENCY string = "USD"
const SINGLE_CONVERTION_AMOUNT int = 1

func (c *CurrencyAPI) Get(currency string) (float64, error) {
	err := exchange.ValidateCode(currency)
	if err != nil {
		return 0, err
	}

	currencyToUpdate := exchange.New(currency)
	value, err := currencyToUpdate.ConvertTo(SERVER_CURRENCY, SINGLE_CONVERTION_AMOUNT)
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
