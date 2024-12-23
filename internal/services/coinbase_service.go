package services

import (
	"encoding/json"
	"io"
	"net/http"
)

type CoinBaseData struct {
	Data struct {
		Rates map [string]string `json:"rates"`
	} `json:"data"`
}

func (b CoinBaseData) FetchData() (string, error) {
	url := "https://api.coinbase.com/v2/exchange-rates?currency=BTC"
	resp, err := http.Get(url)

	if err != nil {
		return "Ошибка подключения к Api Coin Base", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Ошибка получения данных от Api Coin Base", err
	}

	// Парсим JSON-ответ
	var ratesResponse CoinBaseData
	if err := json.Unmarshal(body, &ratesResponse); err != nil {
		return "Ошибка при разборе JSON:", err
	}

	// Извлекаем поле USD
	usdRate, exists := ratesResponse.Data.Rates["USD"]
	if !exists {
		return "Курс USD не найден", nil
	}

	return usdRate, nil
}