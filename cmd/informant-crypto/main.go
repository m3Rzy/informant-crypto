package main

import (
	"fmt"
	"strings"
	"time"

	"informant-crypto/internal/models"
	"informant-crypto/internal/services"
	"informant-crypto/internal/utils"
)

func main() {
	// Засекаем время запуска программы
	startTime := time.Now()

	// var volume float64 = 0.001032 // соотношение 100 usdt / btc
	var volume float64 = 0.0001032 // соотношение 10 usdt / btc
	cur := "BTC"
	rate := "USDT"

	okx_client := utils.BaseResponseBuilder(cur, rate, "OKX", fmt.Sprintf("https://www.okx.com/api/v5/market/index-components?index=%s-%s", cur, rate))
	coinbase_client := utils.BaseResponseBuilder(cur, rate, "CoinBase", fmt.Sprintf("https://api.coinbase.com/v2/exchange-rates?currency=%s", cur))
	kucoin_client := utils.BaseResponseBuilder(cur, rate, "Kucoin", fmt.Sprintf("https://api.kucoin.com/api/v1/market/stats?symbol=%s-%s", cur, rate))
	bybit_client := utils.BaseResponseBuilder(cur, rate, "Bybit", fmt.Sprintf("https://api-testnet.bybit.com/v5/market/tickers?category=inverse&symbol=%s%s", cur, rate))
	binance_client := utils.BaseResponseBuilder(cur, rate, "Binance", fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s%s", cur, rate))
	bitget_client := utils.BaseResponseBuilder(cur, rate, "Bitget", fmt.Sprintf("https://api.bitget.com/api/v2/spot/market/tickers?symbol=%s%s", cur, rate))
	htx_client := utils.BaseResponseBuilder(cur, rate, "Htx", fmt.Sprintf("https://api.huobi.pro/market/detail/merged?symbol=%s%s", strings.ToLower(cur), strings.ToLower(rate)))
	mexc_client := utils.BaseResponseBuilder(cur, rate, "Mexc", fmt.Sprintf("https://api.mexc.com/api/v3/ticker/price?symbol=%s%s", strings.ToUpper(cur), strings.ToUpper(rate)))
	exmo_client := utils.BaseResponseBuilder(cur, rate, "Exmo", "https://api.exmo.com/v1.1/ticker")

	okx_model := models.Okx{}
	coin_base_model := models.CoinBase{Rate: rate}
	kucoin_model := models.Kucoin{}
	bybit_model := models.Bybit{}
	binance_model := models.Binance{}
	bitget_model := models.Bitget{}
	htx_model := models.Htx{}
	mexc_model := models.Mexc{}
	exmo_model := models.Exmo{}

	for {
		fmt.Print("\n")
		utils.ToString(okx_client, okx_model, services.FetchData(okx_client))
		utils.ToString(coinbase_client, coin_base_model, services.FetchData(coinbase_client))
		utils.ToString(kucoin_client, kucoin_model, services.FetchData(kucoin_client))
		utils.ToString(bybit_client, bybit_model, services.FetchData(bybit_client))
		utils.ToString(binance_client, binance_model, services.FetchData(binance_client))
		utils.ToString(bitget_client, bitget_model, services.FetchData(bitget_client))
		utils.ToString(htx_client, htx_model, services.FetchData(htx_client))
		utils.ToString(mexc_client, mexc_model, services.FetchData(mexc_client))
		utils.ToString(exmo_client, exmo_model, services.FetchData(exmo_client))

		fmt.Println(services.Sort(services.Spisok, volume, 10.00))

		// Вычисляем и выводим время, прошедшее с начала работы программы
		elapsed := time.Since(startTime)
		hours := int(elapsed.Hours())
		minutes := int(elapsed.Minutes()) % 60
		seconds := int(elapsed.Seconds()) % 60
		fmt.Printf("Время с момента запуска программы: %02d:%02d:%02d\n", hours, minutes, seconds)

		time.Sleep(1 * time.Second) // Задержка 1 сек
	}
}
