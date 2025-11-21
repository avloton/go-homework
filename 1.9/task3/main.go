
/*
Задача 3: Структура "Банковский счет" и методы для работы с балансом
Описание:
Создайте структуру BankAccount с полями: Owner (владелец счета) и Balance (текущий баланс).
Реализуйте методы:
Deposit(amount float64), увеличивающий баланс.
Withdraw(amount float64), уменьшающий баланс, если хватает средств, иначе выводит сообщение о недостатке средств.
Что нужно сделать:
Объявить структуру и методы.
Создать счет, пополнить его, попытаться снять деньги и вывести итоговый баланс.
*/

package main

import "fmt"

type BankAccount struct {
	Owner string
	Balance float64
}
func (a *BankAccount) Deposit(amount float64) {
	fmt.Printf("Кладем на счет сумму: %v\n", amount)
	a.Balance += amount
}
func (a *BankAccount) Withdraw(amount float64) {
	fmt.Printf("Попытка снять сумму: %v\n", amount)
	currentBalance := a.Balance
	if (currentBalance - amount) < 0 {
		fmt.Printf("Ошибка, на счете недостаточно средств: %v\n", a.Balance)
	} else {
		a.Balance -= amount
	}
}

func main() {
	account := BankAccount {Owner: "Владелец", Balance: 0.0}
	account.Deposit(101)
	fmt.Println(account)
	account.Withdraw(102)
	fmt.Println(account)
	account.Withdraw(101)
	fmt.Println(account)
}