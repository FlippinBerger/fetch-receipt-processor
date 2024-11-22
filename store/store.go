package store

import (
	"github.com/flippinberger/fetch-receipt-processor/model"
	"github.com/google/uuid"
)

type Store struct {
	receipts map[string]model.Receipt
}

var receiptStore *Store

func newStore() {
	receiptStore = &Store{
		receipts: make(map[string]model.Receipt),
	}
}

func GetStore() *Store {
	if receiptStore == nil {
		newStore()
	}

	return receiptStore
}

func (s *Store) GenerateID() string {
	id := ""
	for {
		id = uuid.NewString()
		if _, ok := s.receipts[id]; !ok {
			return id
		}
	}
}

func (s *Store) StoreReceipt(id string, r model.Receipt) {
	s.receipts[id] = r
}

func (s *Store) GetReceipt(id string) (*model.Receipt, error) {
	if r, ok := s.receipts[id]; !ok {
		return nil, &NotFoundError{}
	} else {
		return &r, nil
	}
}
