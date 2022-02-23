package service

import (
	"bravo/dao"
	models "bravo/model"
)

func ConvertFromTo(conversionParams models.ConversionRequestModel) float64 {
	var fromValue, toValue = dao.GetCoinValues(conversionParams.From, conversionParams.To)
	var coefficientOfConversion float64 = fromValue / toValue
	return coefficientOfConversion * conversionParams.Value
}
