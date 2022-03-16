package currency

//go:generate go run -mod=mod github.com/golang/mock/mockgen -destination=./testutil/mock_service.go -package=testutil . Service

type Service interface {
	Delete(string) error
	Get(string) (float64, error)
	GetAllKeys() ([]string, error)
	Set(string, float64) error
	Convert(string, string, float64) (float64, error)
}

//go:generate go run -mod=mod github.com/golang/mock/mockgen -destination=./testutil/mock_external_service.go -package=testutil . ExternalService

type ExternalService interface {
	Get(string) (float64, error)
}
