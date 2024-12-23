package services

import (
	"encoding/json"
	"fmt"
)

type CoinBase struct {
	Data     struct {
		Rates map[string]string `json:"rates"`
	} `json:"data"`
}

func (c CoinBase) Fetch(body []byte) string {
	if err := json.Unmarshal(body, &c); err != nil {
		fmt.Errorf("Ошибка парсинга JSON: %w", err)
		return err.Error()
	}
	return c.Data.Rates["USD"]
}