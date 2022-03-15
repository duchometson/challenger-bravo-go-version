package currencyapi

import "github.com/asvvvad/exchange"

type CurrencyAPI struct {
	client     *exchange.Exchange
	repository Repository
}

const BaseCurrency string = "USD"

func (c *CurrencyAPI) Get(currency string) (float64, error) {
	err := exchange.ValidateCode(currency)
	if err != nil {
		return 0, err
	}

	value, err := c.client.LatestRatesSingle(currency)
	if err != nil {
		return 0, err
	}

	valueAsFloat, _ := value.Float64()

	return valueAsFloat, nil
}

func New() *CurrencyAPI {
	return &CurrencyAPI{
		client: exchange.New(BaseCurrency),
	}
}
