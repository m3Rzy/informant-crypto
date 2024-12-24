package services

import (
	"fmt"
	"io"
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
		fmt.Printf("Ошибка подключения к API %s: %v", b.Source, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Ошибка чтения ответа: %v", err)
	}

	// if resp.StatusCode != http.StatusOK {
	// 	fmt.Printf("Неожиданный статус ответа: %d\nПроблема с body: %s", resp.StatusCode, body)
	// }

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
