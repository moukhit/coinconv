package entity

import "strings"

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
	if len(c.Codes) > LimitToConvert {
		return ErrLimitExceeded
	}

	for _, code := range c.Codes {
		if len(code) == 0 {
			return ErrInvalidConvertTo
		}
	}

	return nil
}
