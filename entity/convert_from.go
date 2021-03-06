package entity

import (
	"github.com/moukhit/crypto-currency-converter/apperrors"
	"strings"
)

// ConvertFrom is a structure containing details of amount and currency for conversion
type ConvertFrom struct {
	Amount float32
	Code   string
}

func NewConvertFrom(amount float32, code string) (*ConvertFrom, error) {
	code = strings.TrimSpace(code)
	code = strings.ToUpper(code)

	c := &ConvertFrom{
		Amount: amount,
		Code:   code,
	}

	if err := c.Validate(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *ConvertFrom) Validate() error {
	if c.Amount <= 0 || len(c.Code) == 0 {
		return apperrors.ErrInvalidConvertFrom
	}

	return nil
}
