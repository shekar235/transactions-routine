package services

import (
	"fmt"
	"transactions-routine/models"
	"transactions-routine/repository"
)

type AccountServiceI interface {
	CreateAccount(documentNumber string) (*models.Account, error)
	GetAccount(accountID int64) (*models.Account, error)
}

type AccountService struct {
	accountRepo repository.AccountRepository
}

func NewAccountService(accountRepo repository.AccountRepository) AccountServiceI {
	return &AccountService{accountRepo: accountRepo}
}

func (s *AccountService) CreateAccount(documentNumber string) (*models.Account, error) {
	account := &models.Account{DocumentNumber: documentNumber}
	err := s.accountRepo.CreateAccount(account)
	return account, err
}

func (s *AccountService) GetAccount(accountID int64) (*models.Account, error) {
	account, err := s.accountRepo.GetAccountByID(accountID)
	if err != nil {
		return nil, fmt.Errorf("account not found")
	}
	return account, nil
}
