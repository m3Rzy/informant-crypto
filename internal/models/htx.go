package models

import (
	"encoding/json"
	"fmt"
	"log"
)

type Htx struct {
	Tick struct {
		Ask []float64 `json:"ask"`
	} `json:"tick"`
}

func (h Htx) Transformate(body []byte) string {
	if err := json.Unmarshal(body, &h); err != nil {
		log.Fatal("Ошибка парсинга JSON: ", err)
		return ""
	}
	return fmt.Sprintf("%f", h.Tick.Ask[0])
}