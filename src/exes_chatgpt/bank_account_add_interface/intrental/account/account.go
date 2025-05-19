package account

import (
	"errors"
	"github.com/Petro-vich/stepik.PRO.Go/src/exes_chatgpt/bank_account_add_interface/pkg/bankiface"
)

// Account — базовая реализация BankAccount.
type Account struct {
	Owner   string
	Balance float64
}

// NewAccount конструктор для создания нового счёта.
func NewAccount(owner string) *Account {
	return &Account{Owner: owner}
}

// Deposit пополняет баланс, сумма должна быть неотрицательной.
func (a *Account) Deposit(amount float64) error {
	if amount < 0 {
		return errors.New("сумма для пополнения должна быть положительной")
	}
	a.Balance += amount
	return nil
}

// Withdraw снимает средства, проверяя сумму и наличие денег.
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

// GetBalance возвращает текущий баланс.
func (a *Account) GetBalance() float64 {
	return a.Balance
}

// GetOwner возвращает имя владельца счёта.
func (a *Account) GetOwner() string {
	return a.Owner
}

// Гарантируем, что *Account удовлетворяет интерфейсу BankAccount.
var _ bankiface.BankAccount = (*Account)(nil)
