package main

import (
	"fmt"
	"strconv"
	"time"

	"informant-crypto/internal/interfaces"
	"informant-crypto/internal/services"
)

func main() {
	apiClients := []interfaces.ApiFetcher{
		services.CoinBaseData{Currency: "BTC", Rate: "USDT"},
		services.OkxData{Currency: "BTC", Rate: "USDT"},
	}

	for {
		for _, client := range apiClients {
			rate, err := getBtcRate(client)
			if err != nil {
				fmt.Printf("[%s] Ошибка получения данных: %v\n", client.GetSource(), err)
				continue
			}
			fmt.Printf("[%s] Текущий курс %s/%s: %.2f\n", client.GetSource(), client.GetCurrency(), client.GetRate(), rate)
		}
		time.Sleep(1 * time.Second) // Задержка 1 сек
	}
}

// Функция для получения курса
func getBtcRate(a interfaces.ApiFetcher) (float64, error) {
	rateStr, err := a.FetchData()
	if err != nil {
		return 0, err
	}

	rate, err := strconv.ParseFloat(rateStr, 64)
	if err != nil {
		return 0, fmt.Errorf("ошибка преобразования строки в число: %w", err)
	}
	return rate, nil
}
