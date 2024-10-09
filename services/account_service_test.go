package services

import (
	"testing"
	"transactions-routine/mocks"
	"transactions-routine/models"

	"github.com/stretchr/testify/assert"
)

type AccountTestCase struct {
	Description    string
	DocumentNumber string
	ExpectedID     int64
	ExpectError    bool
}

func TestCreateAccountCases(t *testing.T) {
	mockRepo := mocks.NewMockAccountRepository()
	service := NewAccountService(mockRepo)

	testCases := []AccountTestCase{
		{"Valid Account Creation", "12345678900", 1, false},
		{"Valid Account Creation with Different Document Number", "98765432100", 2, false},
		// Add more cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			account, err := service.CreateAccount(tc.DocumentNumber)

			if tc.ExpectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.ExpectedID, account.AccountID)
				assert.Equal(t, tc.DocumentNumber, account.DocumentNumber)
			}
		})
	}
}

type GetAccountTestCase struct {
	Description string
	AccountID   int64
	ExpectError bool
}

func TestGetAccountCases(t *testing.T) {
	mockRepo := mocks.NewMockAccountRepository()
	service := NewAccountService(mockRepo)

	// Pre-populate with an account
	mockRepo.CreateAccount(&models.Account{DocumentNumber: "12345678900"})

	testCases := []GetAccountTestCase{
		{"Account Exists", 1, false},
		{"Account Does Not Exist", 99, true},
	}

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			account, err := service.GetAccount(tc.AccountID)

			if tc.ExpectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.AccountID, account.AccountID)
			}
		})
	}
}
