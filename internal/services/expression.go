package services

import (
	"fmt"
)

var marginality float64

// Инициализируем карту
var Spisok = make(map[string]float64)

// Функция добавления значения в карту
func Read(coin float64, source string) {
	Spisok[source] = coin // Добавляем или обновляем значение по ключу
}

// Функция для определения самой дешёвой и самой дорогой покупки
func Sort(sp map[string]float64, volume float64, balance float64) (string, string, string) {
	if len(sp) == 0 {
		return "Нет данных для анализа!", "", "" // Обработка случая, когда карта пустая
	}

	var minKey, maxKey string
	var minValue, maxValue float64

	// Инициализируем значениями из первого элемента карты
	first := true
	for key, value := range sp {
		if first {
			minKey, maxKey = key, key
			minValue, maxValue = value, value
			first = false
		} else {
			if value < minValue {
				minValue = value
				minKey = key
			}
			if value > maxValue {
				maxValue = value
				maxKey = key
			}
		}
	}

	minV := minValue * volume
	maxV := maxValue * volume
	marginV := (maxValue - minValue) * volume

	// StringResult(minV, maxV, marginV, minKey, maxKey)
	SumMarginality(marginV, balance)

	// Формируем строки результата
	// minResult := fmt.Sprintf("Покупай на [%s] (%.2f)", minKey, minValue)
	// maxResult := fmt.Sprintf("– Продавай на [%s] (%.2f)", maxKey, maxValue)
	// raznica := fmt.Sprintf("– Мaржа %.2f\n", maxValue-minValue)


	// return minResult, maxResult, raznica
	return StringResult(minV, maxV, marginV, minKey, maxKey)
}

func StringResult(a, b, c float64, minKey, maxKey string) (string, string, string) {
	// Формируем строки результата
	minResult := fmt.Sprintf("Покупай на [%s] (%.5f)", minKey, a)
	maxResult := fmt.Sprintf("– Продавай на [%s] (%.5f)", maxKey, b)
	raznica := fmt.Sprintf("– Мaржа %.5f\n", c)

	return minResult, maxResult, raznica
}

func SumMarginality(sum float64, balance float64) float64 {
	marginality += sum 
	balance += marginality * 0.02 // маржинальность += суммарная маржа + комиссия за переводы, взял за основу 2%
	fmt.Printf("Суммарная маржа: %.5f\n", marginality)
	fmt.Printf("Баланс: %.5f\n", balance)
	return marginality
}
