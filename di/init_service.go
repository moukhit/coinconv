package di

import (
	repo "github.com/moukhit/crypto-currency-converter/infrastructure/repository"
	conv "github.com/moukhit/crypto-currency-converter/usecase/convertion"
)

func InitService() *conv.Service {
	r := repo.NewCmcRepository()
	svc := conv.NewService(r)
	return svc
}
