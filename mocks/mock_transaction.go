package mocks

import (
	"errors"
	"sync"
	"transactions-routine/models" // Update this import path according to your project structure
)

// MockTransactionRepository is a mock implementation of the TransactionRepository interface
type MockTransactionRepository struct {
	Transactions map[int64]*models.Transaction // Exported field
	mu           sync.Mutex
	nextID       int64
}

// NewMockTransactionRepository initializes a new MockTransactionRepository
func NewMockTransactionRepository() *MockTransactionRepository {
	return &MockTransactionRepository{
		Transactions: make(map[int64]*models.Transaction), // Use the exported field
		nextID:       1,
	}
}

// CreateTransaction creates a new transaction in the mock repository
func (m *MockTransactionRepository) CreateTransaction(transaction *models.Transaction) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	transaction.TransactionID = m.nextID
	m.Transactions[m.nextID] = transaction // Use the exported field
	m.nextID++
	return nil
}

// GetTransactionByID retrieves a transaction by ID from the mock repository
func (m *MockTransactionRepository) GetTransactionByID(transactionID int64) (*models.Transaction, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	transaction, exists := m.Transactions[transactionID] // Use the exported field
	if !exists {
		return nil, errors.New("transaction not found")
	}
	return transaction, nil
}
