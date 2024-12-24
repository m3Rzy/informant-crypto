package models

import (
	"encoding/json"
	"log"
)

type Mexc struct {
	Price string `json:"price"`
}

func (m Mexc) Transformate(body []byte) string {
	if err := json.Unmarshal(body, &m); err != nil {
		log.Fatal("Ошибка парсинга JSON: ", err)
		return ""
	}
	return m.Price
}