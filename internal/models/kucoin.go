package models

import (
	"encoding/json"
	"log"
)

type Kucoin struct {
	Data struct {
		Buy string `json:"buy"`
	} `json:"data"`
}

func (k Kucoin) Transformate(body []byte) string {
	if err := json.Unmarshal(body, &k); err != nil {
		log.Fatal("Ошибка парсинга JSON: ", err)
	}

	if k.Data.Buy == "" {
		log.Fatal("Поле [buy] пустое!")
	}
	
	return k.Data.Buy
}
