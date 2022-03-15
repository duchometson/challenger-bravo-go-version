package errorsbravo

import (
	"errors"
)

var CurrencyDoesntExists = errors.New("currency not found")
var MissingParam = errors.New("missing params in request")
var InvalidValueParam = errors.New("invalid value param")
var InternalError = errors.New("internal error")
