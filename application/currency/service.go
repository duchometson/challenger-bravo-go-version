package currency

type Service interface {
	Delete(string) error
	Get(string) (float64, error)
	GetAllKeys() ([]string, error)
	Set(string, float64) error
}

type ExternalService interface {
	Get(string) (float64, error)
}
