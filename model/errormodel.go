package model

import "fmt"

type ApplicationError struct {
	Err error
	Msg string
}

func NewApplicationError(err error, msg string) *ApplicationError {
	return &ApplicationError{
		Err: err,
		Msg: msg,
	}
}

func (a *ApplicationError) Error() string {
	return fmt.Sprintf("%s - %s", a.Err, a.Msg)
}
