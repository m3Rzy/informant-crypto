package main

import (
	"fmt"
	"time"

	"informant-crypto/internal/models"
	"informant-crypto/internal/services"
	"informant-crypto/internal/utils"
)

func scan() (string, string) {
	var cur string
	fmt.Println("Введите Currency: ")
	fmt.Scan(&cur)
	var rate string
	fmt.Println("Введите Rate: ")
	fmt.Scan(&rate)
	return cur, rate
}

func main() {
	cur, rate := scan()

	client := utils.BaseResponseBuilder(cur, rate, "OKX", fmt.Sprintf("https://www.okx.com/api/v5/market/index-components?index=%s-%s", cur, rate))
	client2 := utils.BaseResponseBuilder(cur, rate, "CoinBase", fmt.Sprintf("https://api.coinbase.com/v2/exchange-rates?currency=%s", cur))
	okx1 := models.Okx{}
	cb := models.CoinBase{Rate: rate}

	for {
		utils.ToString(client, okx1, services.FetchData(client))
		utils.ToString(client2, cb, services.FetchData(client2))
		time.Sleep(1 * time.Second) // Задержка 1 сек
	}
}


