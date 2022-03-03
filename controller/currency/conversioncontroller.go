package currency

import (
	"bravo/errorsbravo"
	models "bravo/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Currency struct {
	converter Converter
}

func (c *Currency) ConversionHandler(ctx *gin.Context) {
	from := ctx.Query("from")
	if from == "" {
		ctx.JSON(http.StatusBadRequest, "FAILED")
		return
	}

	to := ctx.Query("to")
	if to == "" {
		ctx.JSON(http.StatusBadRequest, "FAILED")
		return
	}

	value := ctx.Query("value")
	if value == "" {
		ctx.JSON(http.StatusBadRequest, "FAILED")
		return
	}

	parsedValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "ERROR")
		return
	}

	result, err := c.converter.Convert(from, to, parsedValue)
	if err != nil {
		err, ok := err.(*models.ApplicationError)
		if ok {
			if err.Err == errorsbravo.CURRENCY_DOESNT_EXISTS {
				ctx.JSON(http.StatusNotFound, err.Error())
				return
			}
		}

		ctx.JSON(http.StatusInternalServerError, "ERROR")
		return
	}

	ctx.JSON(http.StatusOK, result)
}

//func convert(request *http.Request) (float64, models.RequestError) {
//	requestError := models.RequestError{Msg: errorsbravo.EMPTY_MSG}
//	from, to, value := getParamsFromRequest(request, &requestError)
//	var conversionRequestData models.ConversionRequestModel = models.BuildConversionRequestModelFrom(from, to, value)
//	return service.ConvertFromTo(conversionRequestData, &requestError), requestError
//}
//
//func getParamsFromRequest(request *http.Request, requestError *models.RequestError) (string, string, float64) {
//	var expectedParams []string = controller.getConversionRequestExpectedParams()
//
//	var from, to, valueAsString = tryReadingParams(expectedParams, request, requestError)
//	if requestError.Arg != 0 {
//		return invalidGetRequestReturn()
//	}
//	fmt.Println("From: ", from, " To: ", to, " Value: ", valueAsString)
//	var value = tryConvertValueFromString(valueAsString, requestError)
//	return from, to, value
//}
//
//func tryConvertValueFromString(valueAsString string, requestError *models.RequestError) float64 {
//	defer errorsbravo.InvalidOperation(errorsbravo.INVALID_VALUE_PARAM, requestError)
//	return convertValueFromString(valueAsString)
//}
//
//func convertValueFromString(valueAsString string) float64 {
//	var value, err = strconv.ParseFloat(valueAsString, 64)
//	if err != nil {
//		panic(err)
//	}
//	return value
//}
//
//func tryReadingParams(expectedParams []string, request *http.Request, requestError *models.RequestError) (string, string, string) {
//	defer errorsbravo.InvalidOperation(errorsbravo.MISSING_PARAM, requestError)
//	var from, to, valueAsString = getRequestValues(expectedParams, request)
//	return from, to, valueAsString
//}
//
//func getRequestValues(expectedParams []string, request *http.Request) (string, string, string) {
//	var from string = controller.getRequestParams(expectedParams[0], request)
//	var to string = controller.getRequestParams(expectedParams[1], request)
//	var valueAsString string = controller.getRequestParams(expectedParams[2], request)
//	return from, to, valueAsString
//}
//
//func invalidGetRequestReturn() (string, string, float64) {
//	return "", "", 0.0
//}

func New(converter Converter) *Currency {
	return &Currency{
		converter: converter,
	}
}
