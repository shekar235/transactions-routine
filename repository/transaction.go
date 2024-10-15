package repository

import (
	"sync"
	"transactions-routine/models"
)

type TransactionRepository interface {
	CreateTransaction(transaction *models.Transaction) error
	UpdateTransaction(transaction *models.Transaction) (float64, error)
}

type InMemoryTransactionRepository struct {
	mu           sync.Mutex
	transactions map[int64]*models.Transaction
	lastID       int64
}

func NewInMemoryTransactionRepository() *InMemoryTransactionRepository {
	return &InMemoryTransactionRepository{transactions: make(map[int64]*models.Transaction)}
}

func (repo *InMemoryTransactionRepository) CreateTransaction(transaction *models.Transaction) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	repo.lastID++
	transaction.TransactionID = repo.lastID
	repo.transactions[repo.lastID] = transaction
	return nil
}

func (repo *InMemoryTransactionRepository) UpdateTransaction(transaction *models.Transaction) (float64, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	// transaction.TransactionID = repo.lastID
	// repo.transactions[repo.lastID] = transaction
	amount := transaction.Balance
	for id, trx := range repo.transactions {
		//filter
		if trx.AccountID == transaction.AccountID && trx.Balance < 0 {
			if amount == 0 { //basecase
				return 0, nil
			}

			if -trx.Balance < amount { // on more balance

				amount = trx.Balance + amount
				trx.Balance = 0
			} else { // on less balance
				trx.Balance = trx.Balance + amount
				amount = 0
			}
			repo.transactions[id] = trx
		}
	}
	return amount, nil
}
