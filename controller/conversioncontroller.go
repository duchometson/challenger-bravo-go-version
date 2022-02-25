package controller

import (
	models "bravo/model"
	services "bravo/service"
	"fmt"
	"net/http"
	"strconv"
)

func ConversionHandler(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		var convertedValue, requestError = convert(request)
		if requestError.Arg != 0 {
			generateErrorResponse(responseWriter, requestError)
		} else {
			fmt.Println("Conversion Value: ", convertedValue)
		}
	default:
		http.Error(responseWriter, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func convert(request *http.Request) (float64, models.RequestError) {
	requestError := models.RequestError{Msg: EMPTY_MSG}
	from, to, value := getParamsFromRequest(request, &requestError)
	var conversionRequestData models.ConversionRequestModel = models.BuildConversionRequestModelFrom(from, to, value)
	return services.ConvertFromTo(conversionRequestData), requestError
}

func getParamsFromRequest(request *http.Request, requestError *models.RequestError) (string, string, float64) {
	var expectedParams []string = getConversionRequestExpectedParams()

	var from, to, valueAsString = tryReadingParams(expectedParams, request, requestError)
	if requestError.Arg != 0 {
		return invalidGetRequestReturn()
	}
	fmt.Println("From: ", from, " To: ", to, " Value: ", valueAsString)
	var value = tryConvertValueFromString(valueAsString, requestError)
	return from, to, value
}

func tryConvertValueFromString(valueAsString string, requestError *models.RequestError) float64 {
	defer InvalidOperation(INVALID_VALUE_PARAM, requestError)
	return convertValueFromString(valueAsString)
}

func convertValueFromString(valueAsString string) float64 {
	var value, err = strconv.ParseFloat(valueAsString, 64)
	if err != nil {
		panic(err)
	}
	return value
}

func tryReadingParams(expectedParams []string, request *http.Request, requestError *models.RequestError) (string, string, string) {
	defer InvalidOperation(MISSING_PARAM, requestError)
	var from, to, valueAsString = getRequestValues(expectedParams, request)
	return from, to, valueAsString
}

func getRequestValues(expectedParams []string, request *http.Request) (string, string, string) {
	var from string = getRequestParams(expectedParams[0], request)
	var to string = getRequestParams(expectedParams[1], request)
	var valueAsString string = getRequestParams(expectedParams[2], request)
	return from, to, valueAsString
}

func invalidGetRequestReturn() (string, string, float64) {
	return "", "", 0.0
}
