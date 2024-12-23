package services

import (
	"fmt"
)

// Инициализируем карту
var Spisok = make(map[string]float64)

// Функция добавления значения в карту
func Read(coin float64, source string) {
	Spisok[source] = coin // Добавляем или обновляем значение по ключу
}

// Функция для определения самой дешёвой и самой дорогой покупки
func Sort(sp map[string]float64) (string, string, string) {
	if len(sp) == 0 {
		return "Нет данных", "Нет данных", "" // Обработка случая, когда карта пустая
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

	// Формируем строки результата
	minResult := fmt.Sprintf("Покупай на [%s] (%.2f)", minKey, minValue)
	maxResult := fmt.Sprintf("– Продавай на [%s] (%.2f)", maxKey, maxValue)
	raznica := fmt.Sprintf("– Мaржа %.2f", maxValue - minValue)

	return minResult, maxResult, raznica
}
