package repository

import "github.com/Petro-vich/stepik.PRO.Go/src/exes_chatgpt/bank_account_add_interface/pkg/bankiface"

type AccountRepository interface {
	Save(acc bankiface.BankAccount) error
	Load(owner string) (bankiface.BankAccount, error)
	LoadAll() ([]bankiface.BankAccount, error)
}
