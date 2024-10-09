package services

import (
	"errors"
	"time"
	"transactions-routine/models"
	"transactions-routine/repository"
)

type TransactionServiceI interface {
	CreateTransaction(accountID, operationTypeID int64, amount float64, eventDate time.Time) (*models.Transaction, error)
}

type TransactionService struct {
	accountRepo     repository.AccountRepository
	transactionRepo repository.TransactionRepository
}

func NewTransactionService(accountRepo repository.AccountRepository, transactionRepo repository.TransactionRepository) TransactionServiceI {
	return &TransactionService{accountRepo: accountRepo, transactionRepo: transactionRepo}
}

func (s *TransactionService) CreateTransaction(accountID, operationTypeID int64, amount float64, eventDate time.Time) (*models.Transaction, error) {
	account, err := s.accountRepo.GetAccountByID(accountID)
	if err != nil {
		return nil, errors.New("account not found")
	}

	if operationTypeID == 1 || operationTypeID == 2 || operationTypeID == 3 {
		amount = -amount
	}

	transaction := &models.Transaction{
		AccountID:       account.AccountID,
		OperationTypeID: operationTypeID,
		Amount:          amount,
		EventDate:       eventDate,
	}

	err = s.transactionRepo.CreateTransaction(transaction)
	return transaction, err
}
