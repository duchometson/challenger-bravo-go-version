package controller

import (
	models "bravo/model"
	services "bravo/service"
	"fmt"
	"net/http"
	"strconv"
)

const MISSING_PARAM string = "Faltam parametros"
const INVALID_VALUE_PARAM string = "Parametro value invalido"
const EMPTY_MSG string = ""

func ConversionHandler(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		var convertedValue, conversionError = convert(request)
		if conversionError.Arg != 0 {
			generateErrorResponse(responseWriter, conversionError)
		} else {
			fmt.Println("Conversion Value: ", convertedValue)
		}
	default:
		http.Error(responseWriter, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func convert(request *http.Request) (float64, models.ConversionRequestError) {
	conversionError := models.ConversionRequestError{Msg: EMPTY_MSG}
	from, to, value := getParamsFromRequest(request, &conversionError)
	var conversionRequestData models.ConversionRequestModel = models.BuildConversionRequestModelFrom(from, to, value)
	return services.ConvertFromTo(conversionRequestData), conversionError
}

func getParamsFromRequest(request *http.Request, conversionError *models.ConversionRequestError) (string, string, float64) {
	var expectedParams []string = getConversionRequestExpectedParams()

	var from, to, valueAsString = tryReadingParams(expectedParams, request, conversionError)
	if conversionError.Arg != 0 {
		return invalidGetRequestReturn()
	}
	fmt.Println("From: ", from, " To: ", to, " Value: ", valueAsString)
	var value = tryConvertValueFromString(valueAsString, conversionError)
	return from, to, value
}

func tryConvertValueFromString(valueAsString string, conversionError *models.ConversionRequestError) float64 {
	defer invalidOperation(INVALID_VALUE_PARAM, conversionError)
	return convertValueFromString(valueAsString)
}

func convertValueFromString(valueAsString string) float64 {
	var value, err = strconv.ParseFloat(valueAsString, 64)
	if err != nil {
		panic(err)
	}
	return value
}

func tryReadingParams(expectedParams []string, request *http.Request, conversionError *models.ConversionRequestError) (string, string, string) {
	defer invalidOperation(MISSING_PARAM, conversionError)
	var from, to, valueAsString = getRequestValues(expectedParams, request)
	return from, to, valueAsString
}

func getRequestValues(expectedParams []string, request *http.Request) (string, string, string) {
	var from string = getRequestParams(expectedParams[0], request)
	var to string = getRequestParams(expectedParams[1], request)
	var valueAsString string = getRequestParams(expectedParams[2], request)
	return from, to, valueAsString
}

func invalidOperation(msg string, err *models.ConversionRequestError) {
	if recover := recover(); recover != nil {
		*err = models.BuildConversionRequestErrorFrom(400, msg)
	}
}

func generateErrorResponse(responseWriter http.ResponseWriter, err models.ConversionRequestError) {
	http.Error(responseWriter, err.Msg, err.Arg)
	fmt.Println(err.Msg)
}

func invalidGetRequestReturn() (string, string, float64) {
	return "", "", 0.0
}
