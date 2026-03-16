package service

import (
	"errors"
	"log"

	"github.com/google/uuid"
)

type LedgerService struct {
	transactions map[string]Transaction
}

func NewLedgerService() *LedgerService {
	return &LedgerService{
		transactions: make(map[string]Transaction),
	}
}

func (s *LedgerService) CreateTransaction(reference string, entries []Entry) (Transaction, error) {
	log.Printf("Creating transaction with reference: %s\n", reference)
	var sum int64

	for _, e := range entries {
		sum += e.Amount
	}

	if sum != 0 {
		log.Printf("CreateTransaction error: entries must sum to zero, sum is %d\n", sum)
		return Transaction{}, errors.New("entries must sum to zero")
	}

	tx := Transaction{
		ID:        uuid.New().String(),
		Reference: reference,
		Entries:   entries,
	}

	s.transactions[tx.ID] = tx

	return tx, nil
}

func (s *LedgerService) GetTransaction(id string) (Transaction, error) {
	log.Printf("Getting transaction from store: %s\n", id)
	tx, ok := s.transactions[id]
	if !ok {
		return Transaction{}, errors.New("transaction not found")
	}

	return tx, nil
}

func (s *LedgerService) GetAllTransactions() []Transaction {
	log.Println("Getting all transactions from store")
	txs := make([]Transaction, 0, len(s.transactions))
	for _, tx := range s.transactions {
		txs = append(txs, tx)
	}

	return txs
}
