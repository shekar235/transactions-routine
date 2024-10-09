package repository

import (
	"testing"
	"transactions-routine/models"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	repo := NewInMemoryTransactionRepository()

	tests := []struct {
		name        string
		transaction *models.Transaction
		expectedID  int64
	}{
		{
			name: "Create First Transaction",
			transaction: &models.Transaction{
				AccountID:       1,
				OperationTypeID: 2,
				Amount:          100.00,
			},
			expectedID: 1,
		},
		{
			name: "Create Second Transaction",
			transaction: &models.Transaction{
				AccountID:       1,
				OperationTypeID: 3,
				Amount:          200.00,
			},
			expectedID: 2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.CreateTransaction(tc.transaction)
			assert.NoError(t, err)
			assert.Equal(t, tc.expectedID, tc.transaction.TransactionID)

			repo.mu.Lock()
			defer repo.mu.Unlock()
			_, exists := repo.transactions[tc.expectedID]
			assert.True(t, exists)
		})
	}
}
