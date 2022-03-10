package currencyapi

type CurrencyAPI struct {
}

func (c *CurrencyAPI) Get(currency string) (float64, error) {
	return 0, nil
}

func New() *CurrencyAPI {
	return &CurrencyAPI{}
}
