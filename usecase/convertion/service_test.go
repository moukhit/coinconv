package convertion

import (
	"testing"

	"github.com/moukhit/crypto-currency-converter/entity"
	"github.com/moukhit/crypto-currency-converter/infrastructure/repository"
	"github.com/stretchr/testify/assert"
)

func TestService_GetQuotes(t *testing.T) {
	repo := repository.NewCmcRepository()
	convertFrom := entity.ConvertFrom{
		Code:   "BTC",
		Amount: 1.0,
	}

	convertTo := entity.ConvertTo{
		Codes: []string{"USD", "EUR"},
	}

	request := entity.Request{
		From: convertFrom,
		To:   convertTo,
	}

	service := NewService(repo)
	quotes, err := service.GetQuotes(&request)

	assert.Nil(t, err)
	assert.NotNil(t, quotes)
	assert.Equal(t, 2, len(quotes.List))
}
