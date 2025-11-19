/*
Функция высшего порядка: передача функции как аргумента
Создайте функцию applyOperation(a, b int, op func(int, int) int) int, которая применяет переданную функцию op к числам a и b.

Создайте несколько функций-операций: сложение, вычитание, умножение.
В основной программе вызовите applyOperation с разными операциями и выведите результаты.
*/

package main

import "fmt"

func main() {
	x := 3
	y := 5
	fmt.Println(applyOperation(x, y, sum))
	fmt.Println(applyOperation(x, y, subtraction))
	fmt.Println(applyOperation(x, y, multiply))
}

func sum(a, b int) int {
	return a + b
}
func subtraction(a, b int) int {
	return a - b
}
func multiply(a, b int) int {
	return a * b
}
func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}