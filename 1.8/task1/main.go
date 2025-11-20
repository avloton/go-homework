/*
Реализуйте функцию divide(a, b float64) (float64, error), которая делит a на b. Если b равно нулю, возвращайте ошибку. В основной программе вызовите эту функцию и обработайте возможную ошибку.
*/

package main

import (
	"fmt"
)

func main() {
	d, err := divide(3, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Ошибка. Второй параметр равен нулю.")
	}
	result := a/b
	return result, nil
}