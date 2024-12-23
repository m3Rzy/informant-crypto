package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type OkxData struct {
	Currency string
	Rate     string
	Data     struct {
		Last string `json:"last"`
	} `json:"data"`
}

func (o OkxData) FetchData() (string, error) {
	url := fmt.Sprintf("https://www.okx.com/api/v5/market/index-components?index=%s-%s", o.Currency, o.Rate)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("ошибка подключения к API OKX: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("неожиданный статус ответа: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ошибка чтения ответа: %w", err)
	}

	if err := json.Unmarshal(body, &o); err != nil {
		return "", fmt.Errorf("ошибка парсинга JSON: %w", err)
	}

	return o.Data.Last, nil
}

func (o OkxData) GetCurrency() string {
	return o.Currency
}

func (o OkxData) GetRate() string {
	return o.Rate
}

func (o OkxData) GetSource() string {
	return "OKX"
}
