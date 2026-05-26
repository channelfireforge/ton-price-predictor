package price

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Fetcher struct {
	endpoint string
	client   *http.Client
}

type priceResponse struct {
	Price float64 `json:"price"`
}

func NewFetcher(endpoint string) *Fetcher {
	return &Fetcher{
		endpoint: endpoint,
		client:   &http.Client{Timeout: 10 * time.Second},
	}
}

func (f *Fetcher) Fetch() (float64, error) {
	resp, err := f.client.Get(f.endpoint)
	if err != nil {
		return 0, fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	var pr priceResponse
	if err := json.NewDecoder(resp.Body).Decode(&pr); err != nil {
		return 0, fmt.Errorf("json decode failed: %w", err)
	}

	return pr.Price, nil
}