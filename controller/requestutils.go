package controller

import (
	"net/http"
)

func getConversionRequestExpectedParams() []string {
	return []string{"from", "to", "value"}
}

func getCurrencyGetRequestExpectedParams() []string {
	return []string{"currency"}
}

func getRequestParams(param string, request *http.Request) string {
	keys, ok := request.URL.Query()[param]
	if !ok || len(keys[0]) < 1 {
		panic(ok)
	}

	key := keys[0]
	return key
}
