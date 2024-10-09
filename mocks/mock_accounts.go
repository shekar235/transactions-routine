package mocks

import (
	"errors"
	"sync"
	"transactions-routine/models"
)

// MockAccountRepository is a mock implementation of the AccountRepository interface
type MockAccountRepository struct {
	Accounts map[int64]*models.Account // Exported field
	mu       sync.Mutex
	nextID   int64
}

// NewMockAccountRepository initializes a new MockAccountRepository
func NewMockAccountRepository() *MockAccountRepository {
	return &MockAccountRepository{
		Accounts: make(map[int64]*models.Account), // Use the exported field
		nextID:   1,
	}
}

// CreateAccount creates a new account in the mock repository
func (m *MockAccountRepository) CreateAccount(account *models.Account) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	account.AccountID = m.nextID
	m.Accounts[m.nextID] = account // Use the exported field
	m.nextID++
	return nil
}

// GetAccountByID retrieves an account by ID from the mock repository
func (m *MockAccountRepository) GetAccountByID(accountID int64) (*models.Account, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	account, exists := m.Accounts[accountID] // Use the exported field
	if !exists {
		return nil, errors.New("account not found")
	}
	return account, nil
}
