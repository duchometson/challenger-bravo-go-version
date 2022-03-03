package errorsbravo

import (
	"errors"
)

var CURRENCY_DOESNT_EXISTS = errors.New("currency not found")
var MISSING_PARAM = errors.New("missing params in request")
var INVALID_VALUE_PARAM = errors.New("invalid value param")
var INTERNAL_ERROR = errors.New("internal error")

const EMPTY_MSG string = ""
