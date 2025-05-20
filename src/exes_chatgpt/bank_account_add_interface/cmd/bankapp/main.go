package main

import (
	"fmt"
	"github.com/Petro-vich/stepik.PRO.Go/src/exes_chatgpt/bank_account_add_interface/intrental/account"
	"github.com/Petro-vich/stepik.PRO.Go/src/exes_chatgpt/bank_account_add_interface/pkg/bankiface"
)

func main() {
	var acc bankiface.BankAccount = account.NewAccount("Оля")
	var bonusAcc bankiface.BankAccount = account.NewBonusAccount("Иван", 6.7)

	if err := acc.Deposit(2000); err != nil {
		fmt.Println("Ошибка пополнения:", err)
	}
	fmt.Printf("Баланс %s: %.2f₽\n", acc.GetOwner(), acc.GetBalance())

	if err := acc.Withdraw(2000); err != nil {
		fmt.Println("Ошибка снятия:", err)
	}

	if err := acc.Withdraw(800); err == nil {
		fmt.Printf("Снятие успешно. Новый баланс: %.2f₽\n", acc.GetBalance())
	} else {
		fmt.Printf("Не достаточно средстав. Ваш баланс: %.2f\n", acc.GetBalance())
	}

	if err := bonusAcc.Deposit(500); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Пополнение успешно, ваш счет: %.2f\n", bonusAcc.GetBalance())
	}

}
