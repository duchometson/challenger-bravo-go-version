package repository

type Database interface {
	Get(string) (interface{}, error)
	GetAllKeys() ([]string, error)
	Set(string, interface{}) error
	Delete(string) error
	ErrorNotFound() error
}
