package webapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/internal/entity"
	"net/http"
)

func (repo *WebApiRepo) GetBalance(ctx context.Context, w entity.Wallet) ([]string, error) {
	type WalletBalanceResponse struct {
		Address string `json:"address"`
		Balance string `json:"balance"`
	}
	serverURL := repo.serverAddress + "/get_wallet_balance"

	req, err := http.NewRequestWithContext(ctx, "GET", serverURL, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	q := req.URL.Query()
	q.Add("address", w.Address)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var result WalletBalanceResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}
	balance := parseAnswer(result.Balance)

	return balance, nil
}

func (repo *WebApiRepo) GetTransactionsHistory(ctx context.Context, w entity.Wallet) ([]string, error) {
	type WalletTxHistoryResponse struct {
		History string `json:"history"`
	}
	serverURL := repo.serverAddress + "/get_transactions_history"

	req, err := http.NewRequestWithContext(ctx, "GET", serverURL, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	q := req.URL.Query()
	q.Add("address", w.Address)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var result WalletTxHistoryResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}
	history := parseAnswer(result.History)

	return history, nil
}

func (repo *WebApiRepo) SendCurrency(ctx context.Context, w entity.Wallet, amount, currency, receiver, mine string) (string, error) {
	type WalletSendCurrencyResponse struct {
		SendResult string `json:"sendResult"`
	}
	serverURL := repo.serverAddress + "/send_currency"

	req, err := http.NewRequestWithContext(ctx, "POST", serverURL, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	q := req.URL.Query()
	q.Add("amount", amount)
	q.Add("currency", currency)
	q.Add("sender", w.Address)
	q.Add("receiver", receiver)
	q.Add("mine", mine)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var result WalletSendCurrencyResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", fmt.Errorf("error parsing JSON: %v", err)
	}

	return result.SendResult, nil
}
