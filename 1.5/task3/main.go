/*
Задача 3:
Создайте функцию capitalizeWords(s string) string, которая преобразует каждое слово в строке так, чтобы первая буква была заглавной, а остальные — строчными. Например: "привет мир" → "Привет Мир".
*/

package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	str := "привет мир"
	fmt.Println(capitalizeWords(str))
}

func capitalizeWords(str string) string {
	str_slice := strings.Fields(str)
	var newstr []string = make([]string, len(str_slice))
	for i,v := range str_slice {
		r, size := utf8.DecodeRuneInString(v)
		newstr[i] = string(unicode.ToUpper(r)) + v[size:]
	}
	return strings.Join(newstr, " ")
}