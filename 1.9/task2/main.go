/*
Задача 2: Структура "Студент" и метод для вычисления среднего балла
Описание:
Создайте структуру Student, которая содержит поля: имя (Name) и список оценок (Grades []float64).
Реализуйте метод AverageGrade() float64, который возвращает средний балл студента.
Что нужно сделать:
Объявить структуру и метод.
Создать студента с несколькими оценками и вывести его средний балл.
*/

package main

import (
	"fmt"
)

type Student struct {
	Name string
	Grades []float64
}

func (s Student) AverageGrade() float64{
	var sum float64
	var count int
	for _,v := range s.Grades {
		sum += v
		count++
	}
	return sum/float64(count)
}

func main() {
	student := Student{Name: "Иван Иванов", Grades:[]float64{4.2, 3.9, 5.0}}
	fmt.Printf("Средний бал студента %s: %.2f", student.Name, student.AverageGrade())
}