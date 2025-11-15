/*
Задача 4:
Напишите программу, которая запрашивает у пользователя ввод строки-формулы, а выводит сообщение о правильности написания круглых скобок, например:
Пример 1
строка на вход: (1+1)*(2+2)
вывод: Скобки расставлены верно, 2 открывающиеся, 2 закрывающиеся

Пример 2
Строка на вход: ((1+1) + (2+2) ))
вывод: Скобки расставлены неправильно, 3 открывающиеся, 4 закрывающиеся
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Введите строку: ")
	reader0 := bufio.NewReader(os.Stdin)
	rawInput0, err := reader0.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	input0 := strings.TrimSpace(rawInput0)
	count_opened := 0
	count_closed := 0
	group := 0
	for i,v := range input0 {
		if v == '(' {
			count_opened++
			for ii := i; ii < len(input0); ii++ {
				if input0[ii] == ')' {
					group++
					break
				}
			}
		}
		if v == ')' {
			count_closed++
		}
	}
	if group == count_opened && count_opened == count_closed {
		fmt.Printf("Скобки расставлены верно, %v открывающиеся, %v закрывающиеся\n", count_opened, count_closed)
	} else {
		fmt.Printf("Скобки расставлены неправильно, %v открывающиеся, %v закрывающиеся\n", count_opened, count_closed)
	}
}