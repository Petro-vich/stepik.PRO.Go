package di

import (
	"github.com/Petro-vich/stepik.PRO.Go/src/exes_chatgpt/bank_account_add_interface/intrental/repository"
	"log"
	"os"
)

func InitRepo() repository.AccountRepository {
	path := os.Getenv("PATH_REPO")
	if path == "" {
		log.Fatalln("PATH_REPO environment variable not set")
	}
	return repository.NewJSONRepo(path)
}
