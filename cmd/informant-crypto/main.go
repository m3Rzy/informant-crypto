package main

import (
	"time"

	"informant-crypto/internal/models"
	"informant-crypto/internal/services"
	"informant-crypto/internal/utils"
)

func main() {
	client := utils.BaseResponseBuilder("BTC", "USD", "OKX", "https://www.okx.com/api/v5/market/index-components?index=BTC-USD")
	client2 := utils.BaseResponseBuilder("BTC", "USD", "CoinBase", "https://api.coinbase.com/v2/exchange-rates?currency=BTC")
	okx1 := models.Okx{}
	cb := models.CoinBase{}

	for {
		utils.ToString(client, okx1, services.FetchData(client))
		utils.ToString(client2, cb, services.FetchData(client2))
		time.Sleep(1 * time.Second) // Задержка 1 сек
	}
}


