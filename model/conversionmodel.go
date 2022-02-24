package model

type ConversionRequestModel struct {
	From  string
	To    string
	Value float64
}

func BuildConversionRequestModelFrom(from string, to string, value float64) ConversionRequestModel {
	return ConversionRequestModel{
		From:  from,
		To:    to,
		Value: value,
	}
}
