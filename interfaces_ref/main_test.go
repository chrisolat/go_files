package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockNotifyClient struct {
	notifyClientFunc func(Account) error
}

func (m MockNotifyClient) notifyClient(account Account) error {
	return m.notifyClientFunc(account)
}

func TestHandleCreateAccount(t *testing.T) {
	tests := []struct {
		name           string
		account        Account
		notifyClient   NotifyClient
		expectedStatus int
	}{
		{
			name: "success with Simplenotify",
			account: Account{
				Username: "user1",
				Email:    "user1@example.com",
			},
			notifyClient: MockNotifyClient{
				notifyClientFunc: func(account Account) error {
					assert.Equal(t, "user1", account.Username)
					assert.Equal(t, "user1@example.com", account.Email)
					return nil
				},
			},
			expectedStatus: http.StatusOK,
		},
		// {
		// 	name: "invalid request body",
		// 	account: Account{
		// 		Username: "",
		// 		Email:    "",
		// 	},
		// 	notifyClient: MockNotifyClient{
		// 		notifyClientFunc: func(account Account) error {
		// 			return nil
		// 		},
		// 	},
		// 	expectedStatus: http.StatusBadRequest,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accountHandler := AccountHandler{NotifyClient: tt.notifyClient}

			// Prepare request body
			body, _ := json.Marshal(tt.account)
			req := httptest.NewRequest(http.MethodPost, "/account", bytes.NewReader(body))
			w := httptest.NewRecorder()

			accountHandler.handleCreateAccount(w, req)

			// Check status code
			assert.Equal(t, tt.expectedStatus, w.Result().StatusCode)

			if tt.expectedStatus == http.StatusOK {
				// Check response body
				var respAccount Account
				err := json.NewDecoder(w.Result().Body).Decode(&respAccount)
				assert.NoError(t, err)
				assert.Equal(t, tt.account.Username, respAccount.Username)
				assert.Equal(t, tt.account.Email, respAccount.Email)
			}
		})
	}
}
