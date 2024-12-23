package services

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type CoinBaseData struct {
	Currency string
	Rate string
	Data struct {
		Rates map [string]string `json:"rates"`
	} `json:"data"`
}

func (b CoinBaseData) FetchData() (string, error) {
	if b.Currency == "" {
		panic("Пользователь не указал значение `Currency`")
	}

	if b.Rate == "" {
		panic("Пользователь не указал значение `Rate`")
	}
	url := "https://api.coinbase.com/v2/exchange-rates?currency=" + b.Currency
	resp, err := http.Get(url)

	if resp.StatusCode == 404 {
		panic("404 ошибка! Возможно, неверно указан URL к которому подключаемся или его не существует!")
	}

	if err != nil {
		log.Fatalln("Ошибка подключения к api.coinbase.com / %v", err)
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Ошибка чтения данных api.coinbase.com / %v", err)
		panic(err)
	}

	// Парсим JSON-ответ
	var ratesResponse CoinBaseData
	if err := json.Unmarshal(body, &ratesResponse); err != nil {
		log.Fatalln("Ошибка парсинга JSON-файла api.coinbase.com / %v", err)
		panic(err)
	}

	// Извлекаем поле USD
	usdRate, exists := ratesResponse.Data.Rates[b.Rate]
	if !exists {
		log.Fatalln("Указанного курса не существует! / %v", err)
		panic(err)
	}

	return usdRate, nil
}