/*
Задача 1
Создайте программу, которая читает лог-файл server.log, подсчитывает количество строк, содержащих слово "error", и выводит это число.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("server.log")
	if err != nil {
		fmt.Printf("Ошибка открытия файла: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	var countErr int
	for sc.Scan() {
		line := sc.Text()
		if strings.Contains(strings.ToLower(line), "error") {
			countErr++
		}
	}

	fmt.Println("Количество строк, которые содержат слово 'error':", countErr)
} 