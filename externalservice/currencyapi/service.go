package currencyapi

import "math/big"

type Repository interface {
	Get(currency string) (float64, error)
	Set(currency string, value float64) error
}

//go:generate go run github.com/golang/mock/mockgen -destination=./testutil/mock_external_service.go -package=testutil . ExternalApi

type ExternalApi interface {
	ValidateCode(string) error
	LatestRatesSingle(string) (*big.Float, error)
}
