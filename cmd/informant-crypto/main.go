package main

import (
	"fmt"
	"time"

	"informant-crypto/internal/models"
	"informant-crypto/internal/services"
	"informant-crypto/internal/utils"
)

func main() {
	cur := "BTC"
	rate := "USDT"

	client := utils.BaseResponseBuilder(cur, rate, "OKX", fmt.Sprintf("https://www.okx.com/api/v5/market/index-components?index=%s-%s", cur, rate))
	client2 := utils.BaseResponseBuilder(cur, rate, "CoinBase", fmt.Sprintf("https://api.coinbase.com/v2/exchange-rates?currency=%s", cur))
	client3 := utils.BaseResponseBuilder(cur, rate, "Kucoin", fmt.Sprintf("https://api.kucoin.com/api/v1/market/stats?symbol=%s-%s", cur, rate))
	client4 := utils.BaseResponseBuilder(cur, rate, "Bybit", fmt.Sprintf("https://api-testnet.bybit.com/v5/market/tickers?category=inverse&symbol=%s%s", cur, rate))
	client5 := utils.BaseResponseBuilder(cur, rate, "Binance", fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s%s", cur, rate))
	okx1 := models.Okx{}
	cb := models.CoinBase{Rate: rate}
	k := models.Kucoin{}
	bybit := models.Bybit{}
	binance := models.Binance{}

	for {
		fmt.Print("\n")
		utils.ToString(client, okx1, services.FetchData(client))
		utils.ToString(client2, cb, services.FetchData(client2))
		utils.ToString(client3, k, services.FetchData(client3))
		utils.ToString(client4, bybit, services.FetchData(client4))
		utils.ToString(client5, binance, services.FetchData(client5))

		fmt.Println(services.Sort(services.Spisok))

		time.Sleep(1 * time.Second) // Задержка 1 сек
	}
}


