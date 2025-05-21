package repository

import (
	"encoding/json"
	"errors"
	"github.com/Petro-vich/stepik.PRO.Go/src/exes_chatgpt/bank_account_add_interface/intrental/dto"
	"github.com/Petro-vich/stepik.PRO.Go/src/exes_chatgpt/bank_account_add_interface/pkg/bankiface"
	"os"
	"sync"
)

type jsonRepo struct {
	filePath string
	mu       sync.Mutex
}

func NewJSONRepo(filePath string) AccountRepository {
	return &jsonRepo{filePath: filePath}
}

func (r *jsonRepo) LoadAll() ([]bankiface.BankAccount, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	/*
		os.Stat() (FileInfo, error) Возвращает структуру, описывающую File, а также ошибку, если есть
		В данном случае, если мы получим ошибку и функция проверки ошибки на существование файла os.IsNotExist()
		вернет true. Мы вернем пустой интерфейс. То есть создадим новый и выйдем без ошибок
	*/

	data, err := os.ReadFile(r.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// Файл не существует — возвращаем пустой список без ошибки
			return []bankiface.BankAccount{}, nil
		}
		// Любая другая ошибка — вернуть её
		return nil, err
	}

	/*
		DTO(Data Transfer Object) - Объект для передачи данных
				используется для передачи данных между слоями(Например между repository и service)
				преобразование в другие типа Enity -> DTO -> JSON
		Объявляем переменную dtos типа "срез AccountDTO из пакета dto"
		json.Unmarshal(data, &dtos) Декодирует JSON data и заполняет структуру dtos
	*/

	var dtos []dto.AccountDTO
	if err := json.Unmarshal(data, &dtos); err != nil {
		return nil, err
	}

	var result []bankiface.BankAccount
	for _, d := range dtos {
		result = append(result, d.ToEntity())
	}
	return result, nil
}

func (r *jsonRepo) Save(acc bankiface.BankAccount) error {
	existing, err := r.LoadAll()
	if err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	var allDTOs []dto.AccountDTO
	for _, a := range existing {
		if a.GetOwner() != acc.GetOwner() {
			allDTOs = append(allDTOs, dto.FromEntity(a))
		}
	}

	// Обновляем / добавляем текущий аккаунт
	allDTOs = append(allDTOs, dto.FromEntity(acc))

	data, err := json.MarshalIndent(allDTOs, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(r.filePath, data, 0o644)
}

func (r *jsonRepo) Load(owner string) (bankiface.BankAccount, error) {
	all, err := r.LoadAll()
	if err != nil {
		return nil, err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	for _, acc := range all {
		if acc.GetOwner() == owner {
			return acc, nil
		}
	}
	return nil, errors.New("account not found")
}
