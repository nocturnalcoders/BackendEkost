package transaction

import (
	"backendEkost/kost"
	"errors"
)

type service struct {
	repository     Repository
	kostRepository kost.Repository
}

type Service interface {
	GetTransactionsByKostID(input GetKostTransactionsInput) ([]Transaction, error)
	GetTransactionsByUserID(userID int) ([]Transaction, error)
}

func NewService(repository Repository, kostRepository kost.Repository) *service {
	return &service{repository, kostRepository}
}

func (s *service) GetTransactionsByKostID(input GetKostTransactionsInput) ([]Transaction, error) {
	//get Kost
	//Check Kost.UserID != user_id_yang_melakukan_request

	kost, err := s.kostRepository.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if kost.UserID != input.User.ID {
		return []Transaction{}, errors.New("Not an owner of the Kost")
	}

	transactions, err := s.repository.GetByKostID(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) GetTransactionsByUserID(userID int) ([]Transaction, error) {
	transactions, err := s.repository.GetByUserID(userID)

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
