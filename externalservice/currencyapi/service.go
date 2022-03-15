package currencyapi

type Repository interface {
	Get(currency string) (float64, error)
	Set(currency string, value float64) error
}
