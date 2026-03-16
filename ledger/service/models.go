package service

type Entry struct {
	AccountID string `json:"account_id"`
	Amount    int64  `json:"amount"`
}

type Transaction struct {
	ID        string  `json:"id"`
	Reference string  `json:"reference"`
	Entries   []Entry `json:"entries"`
}
