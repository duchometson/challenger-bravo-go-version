package currency

type Converter interface {
	Convert(string, string, float64) (float64, error)
}
