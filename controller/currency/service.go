package currency

type Converter interface {
	Convert(string, string, float64) (float64, error)
}

type CurrencyManager interface {
	Get(string) (float64, error)
	InsertOrUpdate(string, float64)
	Delete(string) error
}
