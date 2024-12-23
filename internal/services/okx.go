package services

import (
	"encoding/json"
	"fmt"
)

type Okx struct {
	Data	struct {
		Last string `json:"last"`
	} `json:"data"`
}

func (o Okx) Fetch(body []byte) string {

	if err := json.Unmarshal(body, &o); err != nil {
		fmt.Errorf("Ошибка парсинга JSON: %w", err)
		return err.Error()
	}
	
	return o.Data.Last
}