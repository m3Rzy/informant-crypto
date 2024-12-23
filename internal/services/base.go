package services

import (
	"fmt"
	"informant-crypto/internal/interfaces"
	"io"
	"net/http"
	"strconv"
)

type BaseResponse struct {
	Currency string
	Rate string
	Source string
	URL string
}


func Parse(f interfaces.Apis, body []byte) (float64) {
	finish, err := strconv.ParseFloat(f.Fetch(body), 64)
	if err != nil {
		fmt.Errorf("ошибка преобразования строки в число: %w", err)
	}
	return finish
}

func Poly4(b BaseResponse) []byte {
	resp, err := http.Get(b.URL)
	if err != nil {
		fmt.Errorf("ошибка подключения к API OKX: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("неожиданный статус ответа: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("ошибка чтения ответа: %w", err)
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