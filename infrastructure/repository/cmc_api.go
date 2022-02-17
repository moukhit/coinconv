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

// CmcRepository service acts as a repository for getting the currency exchange quotes
type CmcRepository struct {
	baseUrl  string
	endpoint string
	key      string
}

// NewCmcRepository creates new repository
func NewCmcRepository(baseUrl string, endpoint string, key string) *CmcRepository {
	return &CmcRepository{
		baseUrl:  baseUrl,
		endpoint: endpoint,
		key:      key,
	}
}

func (q *CmcRepository) Get(convertFrom *entity.ConvertFrom, convertTo *entity.ConvertTo) (*entity.Quotes, error) {
	// build the url with query parameters
	url := q.buildEntireUrl(convertFrom, convertTo)

	// the timeout for request
	timeout := 30 * time.Second
	client := http.Client{
		Timeout: timeout,
	}

	// preparing request
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("X-CMC_PRO_API_KEY", q.key)
	if err != nil {
		return nil, err
	}

	// send request
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(resp.Body)

	// parsing response
	quotes, err := parseResponse(resp.Body, convertFrom.Code)
	if err != nil {
		return nil, err
	}

	return quotes, nil
}

func (q *CmcRepository) buildQueryString(convertFrom *entity.ConvertFrom, convertTo *entity.ConvertTo) string {
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

func (q *CmcRepository) buildEntireUrl(convertFrom *entity.ConvertFrom, convertTo *entity.ConvertTo) string {
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
