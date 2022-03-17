package currencyapi

type CurrencyAPI struct {
	client     ExternalApi
	repository Repository
}

const BaseCurrency string = "USD"

func (c *CurrencyAPI) Get(currency string) (float64, error) {
	err := c.client.ValidateCode(currency)
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

func New(externalApi ExternalApi) *CurrencyAPI {
	return &CurrencyAPI{
		client: externalApi,
	}
}
