package main

import (
	"fmt"
	"time"

	"informant-crypto/internal/interfaces"
	"informant-crypto/internal/services"
)

func main() {
	client := scan("BTC", "USD", "OKX", "https://www.okx.com/api/v5/market/index-components?index=BTC-USD")
	client2 := scan("BTC", "USD", "CoinBase", "https://api.coinbase.com/v2/exchange-rates?currency=BTC")
	okx1 := services.Okx{}
	cb := services.CoinBase{}

	for {
		ToString(client, okx1, services.Poly4(client))
		ToString(client2, cb, services.Poly4(client2))
		time.Sleep(1 * time.Second) // Задержка 1 сек
	}
}

func ToString(b services.BaseResponse, f interfaces.Apis, body []byte) {
	fmt.Printf("[%s]Текущий курс %s/%s: %.2f\n", b.Source,
	b.Currency,
	b.Rate,
	services.Parse(f, body))
}

func scan(cur, rate, source, url string) services.BaseResponse {
	return services.BaseResponse{
		Currency: cur,
		Rate:     rate,
		Source:   source,
		URL:      url}
}
