package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/moukhit/crypto-currency-converter/config"
	"github.com/moukhit/crypto-currency-converter/entity"
	"github.com/moukhit/crypto-currency-converter/infrastructure/repository"
	"github.com/moukhit/crypto-currency-converter/usecase/convertion"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatalln(ErrInvalidArguments)
	}

	args := os.Args[1:]
	convertFrom, convertTo, err := parseArgs(args)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	request := entity.Request{
		From: *convertFrom,
		To:   *convertTo,
	}

	baseUrl := config.COINMARKET_API_GATEWAY
	key := config.API_KEY
	endpoint := "/v2/tools/price-conversion"

	repo := repository.NewCmcRepository(baseUrl, endpoint, key)
	svc := convertion.NewService(repo)

	quotes, err := svc.GetQuotes(&request)
	if err != nil {
		log.Fatalf("Error getting quotes: %s", err.Error())
	}

	printResult(convertFrom, quotes)
}

func parseArgs(args []string) (*entity.ConvertFrom, *entity.ConvertTo, error) {
	if len(args) < 3 {
		return nil, nil, ErrInvalidArguments
	}

	amount, err := strconv.ParseFloat(args[0], 32)
	if amount <= 0 || err != nil {
		return nil, nil, ErrInvalidArguments
	}

	code := strings.ToUpper(strings.TrimSpace(args[1]))
	if len(code) == 0 {
		return nil, nil, ErrInvalidArguments
	}

	convertFrom, err := entity.NewConvertFrom(float32(amount), code)

	s := strings.Join(args[2:], ",")
	tail := strings.Split(s, ",")

	convertTo, err := entity.NewConvertTo(tail)
	if err != nil {
		return nil, nil, err
	}

	return convertFrom, convertTo, nil
}

func printResult(from *entity.ConvertFrom, quotes *entity.Quotes) {
	fmt.Println(from.Amount, from.Code, "quotes:")
	for _, v := range quotes.List {
		fmt.Println(v.Code, ": ", v.Price)
	}
}
