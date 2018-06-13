package search

import (
	"fmt"

	resty "gopkg.in/resty.v1"
)

type SearchRequest struct {
	Symbols []string `json:"symbols"`
}

func Search(symbols []string, sessionToken string) (*resty.Response, error) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept-Version", "v1").
		SetHeader("Authorization", sessionToken).
		SetBody(&SearchRequest{
			Symbols: symbols,
		}).
		Post("https://trade.dough.com/api/stocks/search")

	return resp, err
}

func OptionChain(symbol string, sessionToken string) (*resty.Response, error) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept-Version", "v1").
		SetHeader("Authorization", sessionToken).
		Get(fmt.Sprintf("https://api.tastyworks.com/option-chains/%s/nested", symbol))

	return resp, err
}
