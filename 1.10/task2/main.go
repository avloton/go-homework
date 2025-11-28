/*
Задача 2: Полиморфное отображение данных

Описание: Создайте интерфейс Shape с методом Area() float64. Реализуйте структуры Circle и Rectangle, которые реализуют этот интерфейс.
Напишите функцию, которая принимает слайс фигур ([]Shape) и выводит площадь каждой фигуры.

Требования:

    Интерфейс Shape.
    Структуры Circle (с радиусом) и Rectangle (с длиной и шириной).
    Функция для вывода площадей всех фигур в слайсе.
*/

package main

import (
	"fmt"
	"math"
)

type Shape interface { 
    Area() float64 
}
type Circle struct {
    radius float64
}
type Rectangle struct {
    lenght float64
    width float64
}
func (c Circle) Area() float64 {
    return math.Pi*c.radius*c.radius
}
func (r Rectangle) Area() float64 {
    return r.lenght*r.width
}
func AreaSlice(slice []Shape) {
    for _,v := range slice {
        fmt.Printf("Площадь фигуры %T равна %.4f\n", v, v.Area())
    }
}
func main() {
    myFigures := []Shape {
        Circle{radius: 1},
        Rectangle{lenght: 4, width: 8},
    }

	AreaSlice(myFigures)
}