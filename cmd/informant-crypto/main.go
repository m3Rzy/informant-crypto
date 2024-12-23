package main

import (
	"fmt"
	"strconv"
	"time"

	"informant-crypto/internal/interfaces"
	"informant-crypto/internal/services"

)

func main() {
	apiClient1 := services.CoinBaseData{Currency: "BTC", Rate: "USDT"}
	for {
		btc, err := strconv.ParseFloat(dataLog(apiClient1), 64)
		if err != nil {
			fmt.Println("Ошибка преобразования строки в число")
		}

		fmt.Printf(
			"Текущий курс %s/%s: %.2F\n",
			apiClient1.Currency,
			apiClient1.Rate,
			btc,
		)

		time.Sleep(1 * time.Second) // Задержка 1 сек
	}
}

func dataLog(a interfaces.ApiFetcher) string {
	body, err := a.FetchData()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return body
}
