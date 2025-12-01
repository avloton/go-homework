/*
Задача 2
Напишите программу, которая принимает список имен файлов в текущей директории, объединяет их содержимое и сохраняет результат в новый файл combined.txt (работать только с текстовыми файлами)
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Ошибка: отсутствуют аргументы")
		os.Exit(1)
	}

	err := CombineTextFiles(os.Args[1:], "combined.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func CombineTextFiles(files []string, combinedFileName string) error {

	combinedFile, err := os.Create(combinedFileName)
	if err != nil {
		return err
	}
	defer combinedFile.Close()

	var resultText string
	var processedFiles int

	for _,filename := range files {

		if filepath.Ext(filename) != ".txt" {
			fmt.Printf("Пропускаем %s (не текстовый файл)\n", filename)
			continue
		}

		textBytes, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка %v\n", err)
			continue
		}

		resultText += string(textBytes)
		processedFiles++

	}

	if processedFiles == 0 {
		return fmt.Errorf("Ни один файл не был обработан")
	}

	_, err = combinedFile.WriteString(resultText)
	if err != nil {
		return err
	}

	fmt.Printf("Файл %s сформирован\n", combinedFileName)
	fmt.Printf("Всего прочитано файлов: %d\n", processedFiles)
	return nil
	
}