package model

type ConversionRequestModel struct {
	From  string
	To    string
	Value float64
}

type ConversionRequestError struct {
	Arg int
	Msg string
}

func BuildConversionRequestModelFrom(from string, to string, value float64) ConversionRequestModel {
	return ConversionRequestModel{
		From:  from,
		To:    to,
		Value: value,
	}
}

func BuildConversionRequestErrorFrom(arg int, msg string) ConversionRequestError {
	return ConversionRequestError{
		Arg: arg,
		Msg: msg,
	}
}
