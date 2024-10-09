package repository

import (
	"fmt"
	"sync"
	"transactions-routine/models"
)

type AccountRepository interface {
	CreateAccount(account *models.Account) error
	GetAccountByID(accountID int64) (*models.Account, error)
}

type InMemoryAccountRepository struct {
	mu       sync.Mutex
	accounts map[int64]*models.Account
	lastID   int64
}

func NewInMemoryAccountRepository() *InMemoryAccountRepository {
	return &InMemoryAccountRepository{accounts: make(map[int64]*models.Account)}
}

func (repo *InMemoryAccountRepository) CreateAccount(account *models.Account) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.lastID++
	account.AccountID = repo.lastID
	repo.accounts[repo.lastID] = account
	return nil
}

func (repo *InMemoryAccountRepository) GetAccountByID(accountID int64) (*models.Account, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	account, exists := repo.accounts[accountID]
	if !exists {
		return nil, fmt.Errorf("account not found")
	}
	return account, nil
}
