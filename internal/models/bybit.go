package models

import (
	"encoding/json"
	"log"
)

type Bybit struct {
	Result struct {
		List []struct {
			LastPrice string `json:"lastPrice"`
		} `json:"list"`
	} `json:"result"`
}

func (b Bybit) Transformate(body []byte) string {
	if err := json.Unmarshal(body, &b); err != nil {
		log.Fatal("Ошибка парсинга JSON: ", err)
		return ""
	}
	return b.Result.List[0].LastPrice
}