package utils

import (
	"fmt"
	"informant-crypto/internal/interfaces"
	"informant-crypto/internal/services"
	"strconv"
)

func ToString(b services.BaseResponse, f interfaces.Transformator, body []byte) {
	coin := parseToFloat(f, body)

	if coin == 0.00 {
		fmt.Printf("! Сервис [%s] был пропущен из-за отсутствия [%s] у него\n", b.Source, b.Rate)
		return
	}

	services.Read(coin, b.Source) // Анализ цен и вывод что покупать/продавать

	// Установим ширину для полей
	fmt.Printf("[%-8s] 	Текущий курс %s/%s: %.2f\n", 
		b.Source, // Выравниваем источник по ширине 10
		b.Currency, // Валюта 1
		b.Rate,     // Валюта 2
		coin)       // Цена
}

func parseToFloat(f interfaces.Transformator, body []byte) float64 {
	finish, err := strconv.ParseFloat(f.Transformate(body), 64)
	if err != nil {
		// fmt.Printf("Ошибка преобразования строки в число: %+v", string(body))
		// fmt.Printf("У сервиса проблема с преобразованием валюты в число, возможно, она отсутствует у него - ")
	}
	return finish
}