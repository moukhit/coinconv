package repository

import (
	"testing"

	"github.com/moukhit/crypto-currency-converter/entity"
	"github.com/stretchr/testify/assert"
)

func Test_Get(t *testing.T) {
	repo := NewCmcRepository()
	convertFrom, err := entity.NewConvertFrom(1.0, "BTC")
	assert.Nil(t, err)

	codes := []string{"USD", "EUR"}
	convertTo, err := entity.NewConvertTo(codes)
	assert.Nil(t, err)

	result, err := repo.Get(convertFrom, convertTo)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, len(result.List))
}
