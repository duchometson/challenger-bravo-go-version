package config

type Config interface {
	GetInitialCurrencies() []string
	GetMainServerCurrency() string
}

type Configurator struct {
	mainServerCurrency string
	initialCurrencies  []string
}

func (c *Configurator) GetInitialCurrencies() []string {
	return c.initialCurrencies
}

func (c *Configurator) GetMainServerCurrency() string {
	return c.mainServerCurrency
}

func NewConfigutaror() *Configurator {
	return &Configurator{
		mainServerCurrency: "USD",
		initialCurrencies:  []string{"BTC", "BRL", "USD", "EUR", "ETH"},
	}
}
