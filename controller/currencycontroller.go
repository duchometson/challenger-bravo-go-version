package controller

import (
	"bravo/model"
	models "bravo/model"
	"bravo/service"
	"fmt"
	"net/http"
)

func CurrencyHandler(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		var currencyValue, requestError = getCurrencyFromRequest(request)
		if requestError.Arg != 0 {
			generateErrorResponse(responseWriter, requestError)
		} else {
			fmt.Println("Value: ", currencyValue)
		}
	case http.MethodPost:
		//insert currency
	case http.MethodDelete:
		//delete currency
	default:
		http.Error(responseWriter, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getCurrencyFromRequest(request *http.Request) (float64, model.RequestError) {
	requestError := models.RequestError{Msg: EMPTY_MSG}
	currencyName := tryReadingParamsCurrency(request, &requestError)
	currencyValue := tryGettingCurrencyValue(currencyName, &requestError)
	return currencyValue, requestError
}

func tryGettingCurrencyValue(currencyName string, requestError *models.RequestError) float64 {
	defer InvalidOperation(CURRENCY_DOESNT_EXISTS, requestError)
	currencyValue := service.GetCurrencyValue(currencyName)
	return currencyValue
}

func tryReadingParamsCurrency(request *http.Request, requestError *model.RequestError) string {
	var expectedParams = getCurrencyGetRequestExpectedParams()
	defer InvalidOperation(MISSING_PARAM, requestError)
	return getRequestParams(expectedParams[0], request)
}
