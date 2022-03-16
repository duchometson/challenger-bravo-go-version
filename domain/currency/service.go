package currency

//go:generate go run github.com/golang/mock/mockgen -destination=./testutil/mock_service.go -package=testutil . Repository

type Repository interface {
	Get(string) (float64, error)
	GetAllKeys() ([]string, error)
	Set(string, float64) error
	Delete(string) error
}
