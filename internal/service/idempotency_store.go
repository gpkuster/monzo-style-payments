package service

import "sync"

// IdempotencyStore uses array fo MVP
type IdempotencyStore struct {
	data map[string]Payment
	mu   sync.Mutex // thread safety
}

func NewIdempotencyStore() *IdempotencyStore {
	return &IdempotencyStore{
		data: make(map[string]Payment),
	}
}

func (s *IdempotencyStore) Get(key string) (Payment, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	payment, ok := s.data[key]
	return payment, ok
}

func (s *IdempotencyStore) Set(key string, payment Payment) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = payment
}
