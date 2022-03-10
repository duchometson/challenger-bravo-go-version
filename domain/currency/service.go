package currency

type Repository interface {
	Get(string) (float64, error)
	GetAllKeys() ([]string, error)
	Set(string, float64) error
	Delete(string) error
}
