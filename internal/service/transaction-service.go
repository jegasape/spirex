package service

import "github.com/jegasape/spirex.git/internal/entity"

type TransactionService interface {
	Add(entity.Detail) entity.Detail
	FindAll() []entity.Detail
}

type transactionService struct {
	transations []entity.Detail
}

func New() TransactionService {
	return &transactionService{}
}

func (ts *transactionService) Add(v entity.Detail) entity.Detail {
	ts.transations = append(ts.transations, v)
	return v
}

func (ts *transactionService) FindAll() []entity.Detail {
	return ts.transations
}
