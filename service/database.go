package service

type Database interface {
	Get(string) (float64, error)
}
