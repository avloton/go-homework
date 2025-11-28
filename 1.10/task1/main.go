/*
Задача 1: Реализация системы оплаты с использованием интерфейсов

Описание: Создайте интерфейс PaymentProcessor с методом Process(amount float64) string, который возвращает строку с результатом обработки платежа.
Реализуйте два типа платежных систем: CreditCard и CryptoWallet, каждый из которых реализует интерфейс PaymentProcessor.
В функции main создайте список платежных систем и вызовите метод Process для каждого, выводя результат на экран.

Требования:

    Интерфейс PaymentProcessor.
    Структуры CreditCard и CryptoWallet, реализующие интерфейс.
    В main создайте массив/слайс этих структур и вызовите их методы.
    P.S. Подумайте, какие поля могут быть у каждой структуры
*/

package main

import "fmt"


type PaymentProcessor interface {
	Process(float64) string
}
type CreditCard struct {
	name string
	amount float64
}
type CryptoWallet struct {
	name string
	amount float64
}
func (cc CreditCard) Process(amount float64) string {
	return fmt.Sprintf("В платежной системе '%s' успешно проведена операция на сумму %.2f", cc.name, amount)
}
func (cw CryptoWallet) Process(amount float64) string {
	return fmt.Sprintf("В платежной системе '%s' успешно проведена операция на сумму %.2f", cw.name, amount)
}

func main() {
	paymentSystems := []PaymentProcessor{
		CreditCard{
			name: "Кредитная карта",
			amount: 100,
		},
		CryptoWallet{
			name: "Криптокошелек",
			amount: 10,
		},
	}
	for _,p := range paymentSystems {
		result := p.Process(1)
		fmt.Println(result)
	}

}