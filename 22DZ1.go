package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Задание 1. Подсчёт чисел в неупорядоченом массиве после заданного числа

/*
Заполните массив неупорядоченными числами на основе генератора случайных чисел. Введите число. Программа должна найти
это число и вывести, сколько чисел находится в массиве после введённого. При отсутствии введённого числа в массиве - вывести 0.
Для удобства проверки вевести массив на экран.
*/

const n = 10

func main() {
	rand.Seed(time.Now().UnixNano())
	var a [n]int
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(10)
	}

	fmt.Println("Введите число от 1 до 9:")
	var value int
	fmt.Scan(&value)
	for i := 0; i < n; i++ {
		fmt.Printf("a[%v]=%v\n", i, a[i])
	}

	index, numbers := find(a, value)
	if index == 0 {
		fmt.Println("Индекс введённого числа", value, "в массиве равен:", index, "Это значит такого числа нет в массиве.")
	} else {
		fmt.Println("Индекс введённого числа", value, "в массиве равен:", index, "Количество элементов после этого числа в массиве:", numbers)
	}
}

// Функция поиска индекса введённого числа в массиве и оставшегося количества элементов после этого числа
func find(a [n]int, value int) (indexValue int, remainderNumbers int) {
	indexValue = 0
	for i := 0; i < n; i++ {
		if a[i] == value {
			indexValue = i
			remainderNumbers = n - 1 - indexValue
			break
		}
	}
	return indexValue, remainderNumbers
}
