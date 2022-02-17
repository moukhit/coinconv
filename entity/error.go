package entity

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
