package models

import (
	"encoding/json"
	"log"
)

type Binance struct {
	// Price map[string]string `json:"price"`
	Price string `json:"price"`
}

func (b Binance) Transformate(body []byte) string {
	if err := json.Unmarshal(body, &b); err != nil {
		log.Fatal("Ошибка парсинга JSON: ", err)
		return ""
	}
	return b.Price
}