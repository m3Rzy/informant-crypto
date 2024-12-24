package models

import (
	"encoding/json"
	"log"
)

type Exmo struct {
	Coin struct {
		Price string `json:"buy_price"`
	} `json:"BTC_USDT"`
}

func (e Exmo) Transformate(body []byte) string {
	if err := json.Unmarshal(body, &e); err != nil {
		log.Fatal("Ошибка парсинга JSON: ", err)
		return ""
	}
	return e.Coin.Price
}