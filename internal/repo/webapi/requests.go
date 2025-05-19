package webapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"main/internal/entity"
	"net/http"
)

func (repo *WebApiRepo) GetBalance(ctx context.Context, w entity.Wallet) ([]string, error) {
	type WalletBalanceResponse struct {
		Address string `json:"address"`
		Balance string `json:"balance"`
	}
	serverURL := repo.serverAddress + "/get_wallet_balance"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, serverURL, nil)
	if err != nil {
		repo.logger.Fatalf("Error creating request: %v", err)
	}

	q := req.URL.Query()
	q.Add("address", w.Address)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		repo.logger.Fatalf("Error sending request: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		repo.logger.Fatalf("Error reading response: %v", err)
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

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, serverURL, nil)
	if err != nil {
		repo.logger.Fatalf("Error creating request: %v", err)
	}

	q := req.URL.Query()
	q.Add("address", w.Address)
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		repo.logger.Fatalf("Error sending request: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		repo.logger.Fatalf("Error reading response: %v", err)
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

	payload := map[string]string{
		"amount":   amount,
		"currency": currency,
		"sender":   w.Address,
		"receiver": receiver,
		"mine":     mine,
	}
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("marshaling payload: %w", err)
	}

	serverURL := repo.serverAddress + "/send_currency"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, serverURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return "", fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		repo.logger.Fatalf("Error sending request: %v", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("unexpected HTTP status %d: %s", resp.StatusCode, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		repo.logger.Fatalf("Error reading response: %v", err)
	}

	var result WalletSendCurrencyResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", fmt.Errorf("error parsing JSON: %v", err)
	}

	return result.SendResult, nil
}
