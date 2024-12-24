package models

import (
	"encoding/json"
	"log"
)

type Bitget struct {
	Data []struct {
		BidPr string `json:"bidPr"`
	} `json:"data"`
}

func (b Bitget) Transformate(body []byte) string {
	if err := json.Unmarshal(body, &b); err != nil {
		log.Fatal("Ошибка парсинга JSON: ", err)
		return ""
	}
	return b.Data[0].BidPr
}