package convertion

import (
	"github.com/moukhit/crypto-currency-converter/entity"
)

// Reader interface
type Reader interface {
	Get(convertFrom *entity.ConvertFrom, convertTo *entity.ConvertTo) (*entity.Quotes, error)
}

// Repository interface
type Repository interface {
	Reader
}

// UseCase interface
type UseCase interface {
	GetQuotes(request *entity.Request) (*entity.Quotes, error)
}
