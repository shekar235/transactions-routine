package repository

import (
	"testing"
	"transactions-routine/models"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	repo := NewInMemoryAccountRepository()

	tests := []struct {
		name     string
		account  *models.Account
		expected int64
	}{
		{
			name:     "Create Account Success",
			account:  &models.Account{DocumentNumber: "123456789"},
			expected: 1,
		},
		{
			name:     "Create Another Account Success",
			account:  &models.Account{DocumentNumber: "987654321"},
			expected: 2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.CreateAccount(tc.account)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, tc.account.AccountID)

			repo.mu.Lock()
			defer repo.mu.Unlock()
			_, exists := repo.accounts[tc.expected]
			assert.True(t, exists)
		})
	}
}

func TestGetAccountByID(t *testing.T) {
	repo := NewInMemoryAccountRepository()

	// Prepopulate repository with accounts
	repo.CreateAccount(&models.Account{DocumentNumber: "123456789"})
	repo.CreateAccount(&models.Account{DocumentNumber: "987654321"})

	tests := []struct {
		name          string
		accountID     int64
		expectedDocNo string
		expectError   bool
	}{
		{
			name:          "Get Existing Account",
			accountID:     1,
			expectedDocNo: "123456789",
			expectError:   false,
		},
		{
			name:          "Get Another Existing Account",
			accountID:     2,
			expectedDocNo: "987654321",
			expectError:   false,
		},
		{
			name:        "Account Not Found",
			accountID:   3,
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			account, err := repo.GetAccountByID(tc.accountID)

			if tc.expectError {
				assert.Error(t, err)
				assert.Nil(t, account)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, account)
				assert.Equal(t, tc.expectedDocNo, account.DocumentNumber)
			}
		})
	}
}
