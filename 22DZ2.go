package main

import "fmt"

// Задание 2. Нахождение первого вхождения числа в упорядоченом массиве (числа могут повторяться)

/*
Заполните упорядоченный массив из 12 элементов и введите число. Необхлдимо реализовать поис первого вхождения заданногочисла в массив.
Сложность алгоритма должна быть минимальна.

При вводе массива 1 2 2 2 3 4 5 6 7 8 9 10 и вводе чила 2 программма должна вывести индекс 1.
*/

const m = 12

func main() {
	a := [m]int{1, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(a)

	fmt.Println("Введите любое число из массива")
	var value int
	fmt.Scan(&value)

	fmt.Println("Индекс данного числа в массиве:", findValueIndex(a, value))

}

func findValueIndex(a [m]int, value int) (index int) {
	index = -1
	min := 0
	max := m - 1

MainCycle:
	for max >= min {
		mid := (max + min) / 2
		if a[mid] > value {
			max = mid - 1
		} else if a[mid] < value {
			min = mid + 1
		} else {
			for i := 0; i < mid; i++ {
				if a[i] == value {
					index = i
					break MainCycle
				}
			}
			index = mid
			break
		}
	}
	return index
}
