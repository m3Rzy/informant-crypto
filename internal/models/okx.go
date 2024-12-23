package models

import (
	"encoding/json"
	"log"
)

type Okx struct {
	Data struct {
		Last string `json:"last"`
	} `json:"data"`
}

func (o Okx) Transformate(body []byte) string {
	if err := json.Unmarshal(body, &o); err != nil {
		log.Fatal("Ошибка парсинга JSON: ", err)
		return ""
	}
	return o.Data.Last
}
