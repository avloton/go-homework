/*Задача 2:
Напишите программу, которая подсчитывает количество гласных букв (а, е, ё, и, о, у, ы, э, ю, я) в введённой пользователем строке.
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
	fmt.Println(countLetters(input0))
}

func countLetters(str string) int {
	count := 0
	for _,v := range str {
		if v == 'а' || v ==  'е' || v == 'ё' || v == 'и' || v == 'о' || v == 'у' || v == 'ы' || v == 'э' || v == 'ю' || v == 'я' {
			count++
		}
	}
	return count
}