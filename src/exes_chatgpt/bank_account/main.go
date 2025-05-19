package main

import (
	"errors"
	"fmt"
)

type Account struct {
	Owner   string
	Balance float64
}

func (a *Account) Deposit(amount float64) error {
	if amount < 0 {
		return errors.New("сумма для пополнения должна быть положительной")
	}
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("сумма для снятия должна быть положительной")
	}
	if amount > a.Balance {
		return errors.New("недостаточно средств на счете")
	}
	a.Balance -= amount
	return nil
}

func (a *Account) balance() float64 {
	return a.Balance
}

func (a *Account) owner() string {
	return a.Owner

}

func NewAccount(owner string) *Account {
	return &Account{Owner: owner}
}

func main() {
	acc := NewAccount("Пётр")

	if err := acc.Deposit(2000); err != nil {
		fmt.Println("Ошибка пополнения:", err)
	}
	fmt.Printf("Баланс %s: %.2f₽\n", acc.owner(), acc.balance())

	if err := acc.Withdraw(2000); err != nil {
		fmt.Println("Ошибка снятия:", err)
	}

	if err := acc.Withdraw(800); err == nil {
		fmt.Printf("Снятие успешно. Новый баланс: %.2f₽\n", acc.balance())
	}
}
