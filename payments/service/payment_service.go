package service

import (
	"github.com/google/uuid"
)

type PaymentService struct {
	store *IdempotencyStore
}

func NewPaymentService() *PaymentService {
	return &PaymentService{
		store: NewIdempotencyStore(),
	}
}

func (s *PaymentService) CreatePayment(idempotencyKey string, amount int64, currency string) Payment {

	// check existing payment
	if payment, ok := s.store.Get(idempotencyKey); ok {
		return payment
	}

	payment := Payment{
		ID:       uuid.New().String(),
		Amount:   amount,
		Currency: currency,
		Status:   "created",
	}

	s.store.Set(idempotencyKey, payment)

	return payment
}
