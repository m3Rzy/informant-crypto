package main

import (
	"fmt"
	"strings"
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
	return strings.ToUpper(cur), strings.ToUpper(rate)
}

func main() {
	cur, rate := scan()

	client := utils.BaseResponseBuilder(cur, rate, "OKX", fmt.Sprintf("https://www.okx.com/api/v5/market/index-components?index=%s-%s", cur, rate))
	client2 := utils.BaseResponseBuilder(cur, rate, "CoinBase", fmt.Sprintf("https://api.coinbase.com/v2/exchange-rates?currency=%s", cur))
	client3 := utils.BaseResponseBuilder(cur, rate, "Kucoin", fmt.Sprintf("https://api.kucoin.com/api/v1/market/stats?symbol=%s-%s", cur, rate))
	okx1 := models.Okx{}
	cb := models.CoinBase{Rate: rate}
	k := models.Kucoin{}

	for {
		utils.ToString(client, okx1, services.FetchData(client))
		utils.ToString(client2, cb, services.FetchData(client2))
		utils.ToString(client3, k, services.FetchData(client3))

		fmt.Println(services.Sort(services.Spisok))

		time.Sleep(1 * time.Second) // Задержка 1 сек
	}
}


