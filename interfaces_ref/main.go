package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type NotifyClient interface {
	notifyClient(Account) error
}

type Simplenotify struct{}

func (s Simplenotify) notifyClient(account Account) error {
	slog.Info("new account created sent from simple notify", "username", account.Username, "email", account.Email)
	return nil
}

type ComplexNotify struct{}

func (c ComplexNotify) notifyClient(account Account) error {
	slog.Info("new account created sent from complex notify", "username", account.Username, "email", account.Email)
	return nil
}

type Account struct {
	Username string
	Email    string
}

type AccountHandler struct {
	NotifyClient
}

func (a AccountHandler) handleCreateAccount(w http.ResponseWriter, r *http.Request) {
	var account Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		slog.Error("failed to decode reponse body", "err", err)
		return
	}

	if err := a.notifyClient(account); err != nil {
		slog.Error("failed to notify account created", "err", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func main() {
	mux := http.NewServeMux()

	var accountHandler AccountHandler

	accountHandler = AccountHandler{ComplexNotify{}}

	mux.HandleFunc("POST /getbalance", accountHandler.handleCreateAccount)

	accountHandler = AccountHandler{
		Simplenotify{},
	}
	mux.HandleFunc("POST /account", accountHandler.handleCreateAccount)

	http.ListenAndServe(":3000", mux)
}
