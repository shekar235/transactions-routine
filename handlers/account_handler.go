package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"transactions-routine/services"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service services.AccountServiceI
}

func NewAccountHandler(service services.AccountServiceI) *AccountHandler {
	return &AccountHandler{service: service}
}

func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var request struct {
		DocumentNumber string `json:"document_number"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	account, err := h.service.CreateAccount(request.DocumentNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

func (h *AccountHandler) GetAccount(w http.ResponseWriter, r *http.Request) {
	accountID, err := strconv.ParseInt(mux.Vars(r)["accountId"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	account, err := h.service.GetAccount(accountID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(account)
}
