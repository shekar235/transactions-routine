package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"transactions-routine/mocks"
	"transactions-routine/models"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockAccountServiceI(ctrl)
	handler := NewAccountHandler(mockService)

	tests := []struct {
		name           string
		input          interface{}
		mockBehavior   func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Successful Account Creation",
			input: map[string]string{
				"document_number": "123456789",
			},
			mockBehavior: func() {
				mockService.EXPECT().
					CreateAccount("123456789").
					Return(&models.Account{AccountID: 1, DocumentNumber: "123456789"}, nil)
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "Invalid Request Body",
			input:          "invalid-json",
			mockBehavior:   func() {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid request\n",
		},
		{
			name: "Service Error",
			input: map[string]string{
				"document_number": "123456789",
			},
			mockBehavior: func() {
				mockService.EXPECT().
					CreateAccount("123456789").
					Return(nil, errors.New("Service error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Service error\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			body, _ := json.Marshal(tc.input)
			req := httptest.NewRequest(http.MethodPost, "/accounts", bytes.NewReader(body))
			w := httptest.NewRecorder()

			tc.mockBehavior()
			handler.CreateAccount(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
			if tc.expectedBody != "" {
				assert.Equal(t, tc.expectedBody, w.Body.String())
			}
		})
	}
}

func TestGetAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockAccountServiceI(ctrl)
	handler := NewAccountHandler(mockService)

	tests := []struct {
		name           string
		accountID      string
		mockBehavior   func()
		expectedStatus int
		expectedBody   string
	}{
		{
			name:      "Successful Account Retrieval",
			accountID: "1",
			mockBehavior: func() {
				mockService.EXPECT().
					GetAccount(int64(1)).
					Return(&models.Account{AccountID: 1, DocumentNumber: "123456789"}, nil)
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid Account ID",
			accountID:      "invalid",
			mockBehavior:   func() {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid account ID\n",
		},
		{
			name:      "Account Not Found",
			accountID: "2",
			mockBehavior: func() {
				mockService.EXPECT().
					GetAccount(int64(2)).
					Return(nil, errors.New("Account not found"))
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   "Account not found\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/accounts/"+tc.accountID, nil)
			req = mux.SetURLVars(req, map[string]string{"accountId": tc.accountID})
			w := httptest.NewRecorder()

			tc.mockBehavior()
			handler.GetAccount(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
			if tc.expectedBody != "" {
				assert.Equal(t, tc.expectedBody, w.Body.String())
			}
		})
	}
}
