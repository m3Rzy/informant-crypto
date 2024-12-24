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

	okx_client := utils.BaseResponseBuilder(cur, rate, "OKX", fmt.Sprintf("https://www.okx.com/api/v5/market/index-components?index=%s-%s", cur, rate))
	coinbase_client := utils.BaseResponseBuilder(cur, rate, "CoinBase", fmt.Sprintf("https://api.coinbase.com/v2/exchange-rates?currency=%s", cur))
	kucoin_client := utils.BaseResponseBuilder(cur, rate, "Kucoin", fmt.Sprintf("https://api.kucoin.com/api/v1/market/stats?symbol=%s-%s", cur, rate))
	bybit_client := utils.BaseResponseBuilder(cur, rate, "Bybit", fmt.Sprintf("https://api-testnet.bybit.com/v5/market/tickers?category=inverse&symbol=%s%s", cur, rate))
	binance_client := utils.BaseResponseBuilder(cur, rate, "Binance", fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s%s", cur, rate))

	okx_model := models.Okx{}
	coin_base_model := models.CoinBase{Rate: rate}
	kucoin_model := models.Kucoin{}
	bybit_model := models.Bybit{}
	binance_model := models.Binance{}

	for {
		fmt.Print("\n")
		utils.ToString(okx_client, okx_model, services.FetchData(okx_client))
		utils.ToString(coinbase_client, coin_base_model, services.FetchData(coinbase_client))
		utils.ToString(kucoin_client, kucoin_model, services.FetchData(kucoin_client))
		utils.ToString(bybit_client, bybit_model, services.FetchData(bybit_client))
		utils.ToString(binance_client, binance_model, services.FetchData(binance_client))

		fmt.Println(services.Sort(services.Spisok))

		// time.Sleep(1 * time.Second) // Задержка 1 сек
		time.Sleep(500 * time.Millisecond) // Задержка 1 сек
	}
}


