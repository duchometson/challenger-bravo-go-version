package repository

import "context"

type Database interface {
	Get(context.Context, string) (interface{}, error)
	GetAllKeys(context.Context) ([]string, error)
	Set(context.Context, string, interface{}) error
	Delete(context.Context, string) error
	ErrorNotFound() error
}
