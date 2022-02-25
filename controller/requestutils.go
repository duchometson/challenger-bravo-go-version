package controller

import (
	models "bravo/model"
	"fmt"
	"net/http"
)

func getConversionRequestExpectedParams() []string {
	return []string{"from", "to", "value"}
}

func getCurrencyGetRequestExpectedParams() []string {
	return []string{"name"}
}

func getRequestParams(param string, request *http.Request) string {
	keys, ok := request.URL.Query()[param]
	if !ok || len(keys[0]) < 1 {
		panic(ok)
	}

	key := keys[0]
	return key
}

func generateErrorResponse(responseWriter http.ResponseWriter, err models.RequestError) {
	http.Error(responseWriter, err.Msg, err.Arg)
	fmt.Println(err.Msg)
}
