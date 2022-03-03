package errorsbravo

import (
	"errors"
)

var CURRENCY_DOESNT_EXISTS = errors.New("currency not found")

const MISSING_PARAM string = "Faltam parametros"
const INVALID_VALUE_PARAM string = "Parametro value invalido"
const EMPTY_MSG string = ""

//const CURRENCY_DOESNT_EXISTS string = "Currency indicada nao existe"

//func InvalidOperation(msg string, err *models.RequestError) {
//	if recover := recover(); recover != nil {
//		*err = model.BuildRequestErrorFrom(400, msg)
//	}
//}
