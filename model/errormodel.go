package model

type RequestError struct {
	Arg int
	Msg string
}

func BuildRequestErrorFrom(arg int, msg string) RequestError {
	return RequestError{
		Arg: arg,
		Msg: msg,
	}
}
