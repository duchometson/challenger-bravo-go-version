package service

type Database interface {
	Get(string) (float64, error)
	Insert(string, float64)
	Delete(string) error
}
