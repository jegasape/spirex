package controller

import (
	"encoding/json"
	"net/http"

	"github.com/jegasape/spirex/internal/entity"
	"github.com/jegasape/spirex/internal/service"
)

type TransactionController interface {
	Add(w http.ResponseWriter, r *http.Request) (bool, error)
	Edit(w http.ResponseWriter, r *http.Request) (bool, error)
	Delete(w http.ResponseWriter, r *http.Request) (bool, error)
	FindAll(w http.ResponseWriter, r *http.Request)
}

type controller struct {
	service service.TransactionService
}

func New(service service.TransactionService) TransactionController {
	return &controller{
		service: service,
	}
}

func (c *controller) Add(w http.ResponseWriter, r *http.Request) (bool, error) {
	var entity entity.Detail
	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		http.Error(w, "Invalid payload...", http.StatusBadRequest)
		return false, nil
	}

	c.service.Add(entity)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&entity)
	return true, nil
}

func (c *controller) Edit(w http.ResponseWriter, r *http.Request) (bool, error) {
	return true, nil
}

func (c *controller) Delete(w http.ResponseWriter, r *http.Request) (bool, error) {
	return true, nil
}

func (c *controller) FindAll(w http.ResponseWriter, r *http.Request) {
	entity := c.service.FindAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&entity)
}
