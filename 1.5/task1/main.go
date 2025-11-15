/*
Задача 1:
Напишите программу, которая запрашивает у пользователя ввод строки, а затем выводит число - количество символов в строке
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Print("Введите строку: ")
	reader0 := bufio.NewReader(os.Stdin)
	rawInput0, err := reader0.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	input0 := strings.TrimSpace(rawInput0)
	fmt.Printf("Количество символов в строке: %v\n", utf8.RuneCountInString(input0))
}