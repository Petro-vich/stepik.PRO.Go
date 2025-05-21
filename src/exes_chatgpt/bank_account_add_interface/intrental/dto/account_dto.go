package dto

import (
	"github.com/Petro-vich/stepik.PRO.Go/src/exes_chatgpt/bank_account_add_interface/intrental/account"
	"github.com/Petro-vich/stepik.PRO.Go/src/exes_chatgpt/bank_account_add_interface/pkg/bankiface"
)

type AccountDTO struct {
	Owner   string  `json:"owner"`
	Balance float64 `json:"balance"`
	Bonus   float64 `json:"bonus,omitempty"`
}

func (d AccountDTO) ToEntity() bankiface.BankAccount {
	if d.Bonus > 0 {
		acc := account.NewBonusAccount(d.Owner, d.Bonus)
		acc.Balance = d.Balance // напрямую, без приведения
		return acc
	}

	acc := account.NewAccount(d.Owner)
	acc.Balance = d.Balance
	return acc
}

func FromEntity(acc bankiface.BankAccount) AccountDTO {
	dto := AccountDTO{
		Owner:   acc.GetOwner(),
		Balance: acc.GetBalance(),
	}

	if ba, ok := acc.(*account.BonusAccount); ok {
		dto.Bonus = ba.Bonus
	}
	return dto
}
