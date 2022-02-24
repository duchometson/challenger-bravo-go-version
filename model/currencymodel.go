package model

type CurrencyRequestModel struct {
	CurrencyName string
	Value        float64
}

type CurrencyRequestError struct {
	Arg int
	Msg string
}
