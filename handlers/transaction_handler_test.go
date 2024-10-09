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
	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockTransactionServiceI(ctrl)

	// Create the handler with the mock service
	handler := NewTransactionHandler(mockService)

	// Define test cases
	tests := []struct {
		name           string
		input          interface{}
		expectedStatus int
		expectedError  string
		mockBehavior   func()
	}{
		{
			name: "Successful Transaction Creation",
			input: map[string]interface{}{
				"account_id":        int64(1),
				"operation_type_id": int64(2),
				"amount":            100.0,
				"event_date":        "2020-01-01T10:32:07.7199222",
			},
			expectedStatus: http.StatusCreated,
			mockBehavior: func() {
				mockService.EXPECT().
					CreateTransaction(int64(1), int64(2), 100.0, gomock.Any()).
					Return(&models.Transaction{}, nil)
			},
		},
		{
			name:           "Invalid Request Body",
			input:          "invalid json",
			expectedStatus: http.StatusBadRequest,
			expectedError:  "Invalid request\n",
			mockBehavior:   func() {}, // No service call expected
		},
		{
			name: "Invalid Date Format",
			input: map[string]interface{}{
				"account_id":        int64(1),
				"operation_type_id": int64(2),
				"amount":            100.0,
				"event_date":        "invalid-date",
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  "Invalid date format\n",
			mockBehavior:   func() {}, // No service call expected
		},
		{
			name: "Service Error",
			input: map[string]interface{}{
				"account_id":        int64(1),
				"operation_type_id": int64(2),
				"amount":            100.0,
				"event_date":        "2020-01-01T10:32:07.7199222",
			},
			expectedStatus: http.StatusInternalServerError,
			expectedError:  "Service error\n",
			mockBehavior: func() {
				mockService.EXPECT().
					CreateTransaction(int64(1), int64(2), 100.0, gomock.Any()).
					Return(nil, errors.New("Service error"))
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Set up the request body
			body, _ := json.Marshal(tc.input)
			req := httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewReader(body))
			w := httptest.NewRecorder()

			tc.mockBehavior()

			handler.CreateTransaction(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)

			if tc.expectedError != "" {
				assert.Equal(t, tc.expectedError, w.Body.String())
			}
		})
	}
}
