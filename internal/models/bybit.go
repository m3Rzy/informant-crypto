package models

import (
	"encoding/json"
	"log"
)

type Bybit struct {
	Result struct {
		List []struct {
			IndexPrice string `json:"indexPrice"`
		} `json:"list"`
	} `json:"result"`
}

func (b Bybit) Transformate(body []byte) string {
	if err := json.Unmarshal(body, &b); err != nil {
		log.Fatal("Ошибка парсинга JSON: ", err)
		return ""
	}
	return b.Result.List[0].IndexPrice
}