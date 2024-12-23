package utils

import (
	"fmt"
	"informant-crypto/internal/interfaces"
	"informant-crypto/internal/services"
	"strconv"
)

func ToString(b services.BaseResponse, f interfaces.Transformator, body []byte) {
	coin := parseToFloat(f, body)

	services.Read(coin, b.Source) // todo: удалить

	fmt.Printf("[%s]Текущий курс %s/%s: %.2f\n", b.Source,
		b.Currency,
		b.Rate,
		coin)
}


func parseToFloat(f interfaces.Transformator, body []byte) float64 {
	finish, err := strconv.ParseFloat(f.Transformate(body), 64)
	if err != nil {
		fmt.Printf("Ошибка преобразования строки в число: %+v", string(body))
	}
	return finish
}