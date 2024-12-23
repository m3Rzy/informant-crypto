package utils

import (
	"fmt"
	"informant-crypto/internal/interfaces"
	"informant-crypto/internal/services"
	"log"
	"strconv"
)

func ToString(b services.BaseResponse, f interfaces.Transformator, body []byte) {
	fmt.Printf("[%s]Текущий курс %s/%s: %.2f\n", b.Source,
		b.Currency,
		b.Rate,
		parseToFloat(f, body))
}


func parseToFloat(f interfaces.Transformator, body []byte) float64 {
	finish, err := strconv.ParseFloat(f.Transformate(body), 64)
	if err != nil {
		log.Fatalf("Ошибка преобразования строки в число: %v\n %+v", err.Error(), string(body))
	}
	return finish
}