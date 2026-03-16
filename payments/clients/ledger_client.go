package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Entry struct {
	AccountID string `json:"account_id"`
	Amount    int64  `json:"amount"`
}

type CreateTransactionRequest struct {
	Reference string  `json:"reference"`
	Entries   []Entry `json:"entries"`
}

// LedgerClient is a client for interacting with a remote ledger service via HTTP requests.
type LedgerClient struct {
	baseURL string
	client  *http.Client
}

func NewLedgerClient() *LedgerClient {

	url := os.Getenv("LEDGER_SERVICE_URL")

	if url == "" {
		url = "http://localhost:8081"
	}

	return &LedgerClient{
		baseURL: url,
		client:  &http.Client{},
	}
}

func (c *LedgerClient) CreateTransaction(reference string, entries []Entry) error {
	log.Printf("LedgerClient: creating transaction with reference %s\n", reference)

	reqBody := CreateTransactionRequest{
		Reference: reference,
		Entries:   entries,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	resp, err := c.client.Post(
		c.baseURL+"/transactions",
		"application/json",
		bytes.NewBuffer(body),
	)

	if err != nil {
		log.Printf("LedgerClient: error posting to ledger: %v\n", err)
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("LedgerClient: ledger returned status %d\n", resp.StatusCode)
		return fmt.Errorf("ledger returned status %d", resp.StatusCode)
	}

	log.Printf("LedgerClient: transaction %s created successfully\n", reference)
	return nil
}
