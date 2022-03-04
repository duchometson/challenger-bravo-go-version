package currency

type Currency struct {
	converter       Converter
	currencyManager CurrencyManager
}

func New(converter Converter, currencyManager CurrencyManager) *Currency {
	return &Currency{
		converter:       converter,
		currencyManager: currencyManager,
	}
}
