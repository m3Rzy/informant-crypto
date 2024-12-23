package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CoinBaseData struct {
	Currency string
	Rate     string
	Data     struct {
		Rates map[string]string `json:"rates"`
	} `json:"data"`
}

func (b CoinBaseData) FetchData() (string, error) {
	url := "https://api.coinbase.com/v2/exchange-rates?currency=" + b.Currency
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("ошибка подключения к API Coinbase: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("неожиданный статус ответа: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ошибка чтения ответа: %w", err)
	}

	if err := json.Unmarshal(body, &b); err != nil {
		return "", fmt.Errorf("ошибка парсинга JSON: %w", err)
	}

	rate, exists := b.Data.Rates[b.Rate]
	if !exists {
		return "", fmt.Errorf("курс для %s/%s не найден", b.Currency, b.Rate)
	}

	return rate, nil
}

func (b CoinBaseData) GetCurrency() string {
	return b.Currency
}

func (b CoinBaseData) GetRate() string {
	return b.Rate
}

func (b CoinBaseData) GetSource() string {
	return "CoinBase"
}
