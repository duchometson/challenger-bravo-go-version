package repository

type Database interface {
	Delete(string) error
	ErrorNotFound() error
	Get(string) (interface{}, error)
	GetAllKeys() []string
	Set(string, interface{}) error
}
