package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"transactions-routine/services"
)

type TransactionHandler struct {
	service services.TransactionServiceI
}

func NewTransactionHandler(service services.TransactionServiceI) *TransactionHandler {
	return &TransactionHandler{service: service}
}

func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {

	var temp struct {
		AccountID       int64   `json:"account_id"`
		OperationTypeID int64   `json:"operation_type_id"`
		Amount          float64 `json:"amount"`
		EventDate       string  `json:"event_date"`
	}

	// Decode the request body into the temporary struct
	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Parse the EventDate string to time.Time
	parsedDate, err := time.Parse("2006-01-02T15:04:05.9999999", temp.EventDate)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	// Map fields to TransactionRequest struct
	request := TransactionRequest{
		AccountID:       temp.AccountID,
		OperationTypeID: temp.OperationTypeID,
		Amount:          temp.Amount,
		EventDate:       parsedDate,
	}

	// service call
	transaction, err := h.service.CreateTransaction(request.AccountID, request.OperationTypeID, request.Amount, request.EventDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)
}
