/*
Создайте функцию sumAll, которая принимает произвольное количество целых чисел и возвращает их сумму.
Пример использования:
    fmt.Println(sumAll(1, 2, 3)) // 6
    fmt.Println(sumAll(10, -2, 4, 7)) // 19
*/

package main

import "fmt"

func main() {
	fmt.Println(sumAll(1, 2, 3))
	fmt.Println(sumAll(10, -2, 4, 7))
}

func sumAll(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}