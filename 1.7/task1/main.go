/*
Напишите программу, которая создает массив из 10 целых чисел, заполняет его случайными значениями от 1 до 100.
Затем скопируйте этот массив в слайс и отсортируйте его по возрастанию. Выведите исходный массив и отсортированный слайс.
*/

package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	var array [10]int
	for i := range len(array) {
		array[i] = int(rand.Int31n(100))
	}
	slice := array[:]
	fmt.Printf("Исходный массив: %v\n", array)
	sort.Ints(slice)
	fmt.Printf("Отсортированный слайс: %v\n", slice)
}