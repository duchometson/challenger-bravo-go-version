package controller

import (
	"fmt"
	"net/http"
)

func CurrencyHandler(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		var convertedValue, conversionError = convert(request)
		if conversionError.Arg != 0 {
			generateErrorResponse(responseWriter, conversionError)
		} else {
			fmt.Println("Conversion Value: ", convertedValue)
		}
	case http.MethodPost:
		//insert currency
	case http.MethodDelete:
		//delete currency
	default:
		http.Error(responseWriter, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
