package service

import (
	"github.com/jegasape/spirex.git/internal/entity"
)

type TransactionService interface {
	Add(entity.Detail) entity.Detail
	Edit(entity.Detail) entity.Detail
	Delete(entity.Detail) entity.Detail
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

func (ts *transactionService) Edit(v entity.Detail) entity.Detail {
	for i, e := range ts.transations {
		if e.Id == v.Id {
			ts.transations[i] = v
			return v
		}
	}
	return entity.Detail{}
}

func (ts *transactionService) Delete(v entity.Detail) entity.Detail {
	for i, e := range ts.transations {
		if e.Id == v.Id {
			ts.transations = append(ts.transations[:i], ts.transations[i+1:]...)
			return v
		}
	}
	return entity.Detail{}
}

func (ts *transactionService) FindAll() []entity.Detail {
	return ts.transations
}
