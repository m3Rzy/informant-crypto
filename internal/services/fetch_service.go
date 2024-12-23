package services

import (
	"io"
	"log"
	"net/http"
)

type BaseResponse struct {
	Currency string
	Rate     string
	Source   string
	URL      string
}

func FetchData(b BaseResponse) []byte {
	resp, err := http.Get(b.URL)
	if err != nil {
		log.Fatalf("Ошибка подключения к API OKX: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Неожиданный статус ответа: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка чтения ответа: %w", err)
	}

	return body
}

func (b BaseResponse) GetCurrency() string {
	return b.Currency
}

func (b BaseResponse) GetRate() string {
	return b.Rate
}

func (b BaseResponse) GetSource() string {
	return b.Source
}
