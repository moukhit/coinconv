package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/moukhit/crypto-currency-converter/entity"
)

// CoinMarketCap service as a repository for currency exchange quotes
type QuoteCmc struct {
	baseUrl  string
	endpoint string
	key      string
}

// Create new repository
func NewQuoteCmc(baseUrl string, endpoint string, key string) *QuoteCmc {
	return &QuoteCmc{
		baseUrl:  baseUrl,
		endpoint: endpoint,
		key:      key,
	}
}

func (q *QuoteCmc) Get(convertFrom *entity.ConvertFrom, convertTo *entity.ConvertTo) (*entity.Quotes, error) {
	url := q.buildEntireUrl(convertFrom, convertTo)
	timeout := time.Duration(30 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("X-CMC_PRO_API_KEY", q.key)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	quotes, err := parseResponse(request.Body, convertFrom.Code)
	if err != nil {
		return nil, err
	}

	return quotes, nil
}

func (q *QuoteCmc) buildQueryString(convertFrom *entity.ConvertFrom, convertTo *entity.ConvertTo) string {
	var sb strings.Builder

	sb.WriteString("?amount=")
	sb.WriteString(fmt.Sprintf("%f", convertFrom.Amount))
	sb.WriteString("&symbol=")
	sb.WriteString(convertFrom.Code)
	sb.WriteString("&convert=")

	l := len(convertTo.Codes)
	for i, code := range convertTo.Codes {
		sb.WriteString(code)
		if i < l-1 {
			sb.WriteRune(',')
		}
	}

	return sb.String()
}

func (q *QuoteCmc) buildEntireUrl(convertFrom *entity.ConvertFrom, convertTo *entity.ConvertTo) string {
	queryString := q.buildQueryString(convertFrom, convertTo)
	return q.baseUrl + q.endpoint + queryString
}

func parseResponse(response io.Reader, fromCurrency string) (*entity.Quotes, error) {
	b, err := ioutil.ReadAll(response)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}

	data := result["data"].(map[string]interface{})
	from := data[fromCurrency].(map[string]interface{})
	list := from["quote"].(map[string]interface{})

	quotes := entity.Quotes{}
	for k, v := range list {
		prices := v.(map[string]interface{})
		price := prices["price"]

		quote := entity.Quote{
			Code:  k,
			Price: price.(float64),
		}

		quotes.List = append(quotes.List, quote)
	}

	return &quotes, nil
}
