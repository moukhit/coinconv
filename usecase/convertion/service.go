package convertion

import (
	"github.com/moukhit/crypto-currency-converter/entity"
)

// Service for getting quotes for currency exchange
type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetQuotes(request *entity.Request) (*entity.Quotes, error) {
	if request == nil {
		return nil, ErrInvalidRequest
	}

	from := request.From
	to := request.To

	quotes, err := s.repo.Get(&from, &to)
	if quotes == nil || err != nil {
		return nil, err
	}

	if len(quotes.List) == 0 {
		return nil, ErrGettingResponse
	}

	return quotes, nil
}
