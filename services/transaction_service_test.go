package services

import (
	"testing"
	"time"
	"transactions-routine/mocks"
	"transactions-routine/models"

	"github.com/stretchr/testify/assert"
)

type TransactionTestCase struct {
	Description    string
	AccountID      int64
	OperationType  int64
	Amount         float64
	EventDate      time.Time
	ExpectedAmount float64
	ExpectError    bool
}

func TestCreateTransactionCases(t *testing.T) {
	mockAccountRepo := mocks.NewMockAccountRepository()
	mockTransactionRepo := mocks.NewMockTransactionRepository()
	service := NewTransactionService(mockAccountRepo, mockTransactionRepo)

	// Pre-populate with an account
	mockAccountRepo.CreateAccount(&models.Account{DocumentNumber: "12345678900"})

	testCases := []TransactionTestCase{
		{"Normal Purchase", 1, 1, 50.0, time.Now(), -50.0, false},
		{"Purchase with Installments", 1, 2, 23.5, time.Now(), -23.5, false},
		{"Withdrawal", 1, 3, 18.7, time.Now(), -18.7, false},
	}

	for _, tc := range testCases {
		t.Run(tc.Description, func(t *testing.T) {
			transaction, err := service.CreateTransaction(tc.AccountID, tc.OperationType, tc.Amount, tc.EventDate)

			if tc.ExpectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.ExpectedAmount, transaction.Amount)
				assert.Equal(t, tc.OperationType, transaction.OperationTypeID)
			}
		})
	}
}
