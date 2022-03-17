package apilayer

import (
	"math/big"

	"github.com/asvvvad/exchange"
)

type ExternalApiRef struct {
	api *exchange.Exchange
}

func (e *ExternalApiRef) LatestRatesSingle(currency string) (*big.Float, error) {
	return e.api.LatestRatesSingle(currency)
}

func (e *ExternalApiRef) ValidateCode(code string) error {
	return exchange.ValidateCode(code)
}

func New(baseCurrency string) *ExternalApiRef {
	return &ExternalApiRef{
		api: exchange.New(baseCurrency),
	}
}
