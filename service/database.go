package service

type Database interface {
	Get(string) (float64, error)
	InsertOrUpdate(string, float64)
	Delete(string) error
}
