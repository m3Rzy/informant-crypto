package models

import (
	"encoding/json"
	"log"
)

type CoinBase struct {
	Rate string
	Data     struct {
		Rates map[string]string `json:"rates"`
	} `json:"data"`
}

func (c CoinBase) Transformate(body []byte) string {
	if err := json.Unmarshal(body, &c); err != nil {
		log.Fatal("Ошибка парсинга JSON: ", err)
		return ""
	}
	return c.Data.Rates[c.Rate]
}