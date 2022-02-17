package entity

import (
	"github.com/moukhit/crypto-currency-converter/apperrors"
	"strings"
)

// ConvertTo contains a list of currencies convert to
type ConvertTo struct {
	Codes []string
}

func NewConvertTo(codes []string) (*ConvertTo, error) {
	for _, code := range codes {
		code = strings.TrimSpace(code)
		code = strings.ToUpper(code)
	}

	c := &ConvertTo{
		Codes: codes,
	}

	if err := c.Validate(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *ConvertTo) Validate() error {
	if len(c.Codes) > apperrors.LimitToConvert {
		return apperrors.ErrLimitExceeded
	}

	for _, code := range c.Codes {
		if len(code) == 0 {
			return apperrors.ErrInvalidConvertTo
		}
	}

	return nil
}
