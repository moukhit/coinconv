package apperrors

import (
	"errors"
	"fmt"
)

const (
	// LimitToConvert is a maximum number of currencies for getting quotes for conversion
	LimitToConvert int = 120
)

var ErrInvalidConvertFrom = errors.New("invalid amount or currency code convert from")
var ErrInvalidConvertTo = errors.New("one or more invalid currency codes convert to")
var ErrLimitExceeded = fmt.Errorf("the maximum number of currencies convert to is %d", LimitToConvert)
var ErrInvalidArguments = errors.New("invalid arguments => \n\t sample usage: 123.45 BTC USD,EUR")
var ErrInvalidRequest = errors.New("invalid request for currency conversion")
var ErrGettingResponse = errors.New("error has occurred while getting response from service")
