package repository

import (
	"sync"
	"transactions-routine/models"
)

type TransactionRepository interface {
	CreateTransaction(transaction *models.Transaction) error
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
