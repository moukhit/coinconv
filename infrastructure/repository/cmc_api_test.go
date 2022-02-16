package repository

import (
	"testing"

	"github.com/moukhit/crypto-currency-converter/config"
	"github.com/moukhit/crypto-currency-converter/entity"
	"github.com/stretchr/testify/assert"
)

func Test_Get(t *testing.T) {
	baseUrl := config.COINMARKET_API_GATEWAY
	key := config.API_KEY
	endpoint := "/v2/tools/price-conversion"

	repo := NewQuoteCmc(baseUrl, endpoint, key)

	convertFrom := entity.ConvertFrom{
		Code:   "BTC",
		Amount: 1.0,
	}

	convertTo := entity.ConvertTo{
		Codes: []string{"USD", "EUR"},
	}

	result, err := repo.Get(&convertFrom, &convertTo)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}
