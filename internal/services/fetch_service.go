package services

import (
	"fmt"
	"informant-crypto/internal/interfaces"
	"io"
	"log"
	"net/http"
	"strconv"
)

type BaseResponse struct {
	Currency string
	Rate     string
	Source   string
	URL      string
}

func ToString(b BaseResponse, f interfaces.Transformator, body []byte) {
	fmt.Printf("[%s]Текущий курс %s/%s: %.2f\n", b.Source,
		b.Currency,
		b.Rate,
		parseToFloat(f, body))
}

func parseToFloat(f interfaces.Transformator, body []byte) float64 {
	finish, err := strconv.ParseFloat(f.Transformate(body), 64)
	if err != nil {
		log.Fatalf("Ошибка преобразования строки в число: %w", err)
	}
	return finish
}

func Poly4(b BaseResponse) []byte {
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
