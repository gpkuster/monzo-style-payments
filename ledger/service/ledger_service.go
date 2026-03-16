package service

import (
	"errors"

	"github.com/google/uuid"
)

type LedgerService struct{}

func NewLedgerService() *LedgerService {
	return &LedgerService{}
}

func (s *LedgerService) CreateTransaction(reference string, entries []Entry) (Transaction, error) {

	var sum int64

	for _, e := range entries {
		sum += e.Amount
	}

	if sum != 0 {
		return Transaction{}, errors.New("entries must sum to zero")
	}

	tx := Transaction{
		ID:        uuid.New().String(),
		Reference: reference,
		Entries:   entries,
	}

	return tx, nil
}
