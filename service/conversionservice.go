package service

import (
	"bravo/dao"
	"bravo/errorsbravo"
	models "bravo/model"
)

func ConvertFromTo(conversionParams models.ConversionRequestModel, requestError *models.RequestError) float64 {
	defer errorsbravo.InvalidOperation(errorsbravo.CURRENCY_DOESNT_EXISTS, requestError)
	var fromValue, toValue = dao.GetCoinValues(conversionParams.From, conversionParams.To)
	var coefficientOfConversion float64 = fromValue / toValue
	return coefficientOfConversion * conversionParams.Value
}
