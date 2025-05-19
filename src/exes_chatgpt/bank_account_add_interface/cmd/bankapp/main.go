package main

import (
	"fmt"
	"github.com/Petro-vich/stepik.PRO.Go/src/exes_chatgpt/bank_account_add_interface/intrental/account"
	"github.com/Petro-vich/stepik.PRO.Go/src/exes_chatgpt/bank_account_add_interface/pkg/bankiface"
)

func main() {
	var acc bankiface.BankAccount = account.NewAccount("Оля")

	if err := acc.Deposit(2000); err != nil {
		fmt.Println("Ошибка пополнения:", err)
	}
	fmt.Printf("Баланс %s: %.2f₽\n", acc.GetOwner(), acc.GetBalance())

	if err := acc.Withdraw(2000); err != nil {
		fmt.Println("Ошибка снятия:", err)
	}

	if err := acc.Withdraw(800); err == nil {
		fmt.Printf("Снятие успешно. Новый баланс: %.2f₽\n", acc.GetBalance())
	}
}
