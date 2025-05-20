package account

import "fmt"

type BonusAccount struct {
	Owner   string
	Balance float64
	Bonus   float64
}

func NewBonusAccount(owner string, bonus float64) *BonusAccount {
	return &BonusAccount{
		Owner: owner,
		Bonus: bonus,
	}
}

func (ba *BonusAccount) Deposit(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("account: the amount must be greater than zero")
	}

	bonus := amount * ba.Bonus / 100
	ba.Balance += amount + bonus
	return nil
}

func (ba *BonusAccount) Withdraw(amount float64) error {
	if amount > ba.Balance {
		return fmt.Errorf("account: insufficient funds. Your balance: %.2f", ba.Balance)
	}

	if amount < 0 {
		return fmt.Errorf("account: the amount must be greater than zero")
	}
	ba.Balance -= amount
	return nil
}

func (ba *BonusAccount) GetBalance() float64 {
	return ba.Balance
}

func (ba *BonusAccount) GetOwner() string {
	return ba.Owner
}
