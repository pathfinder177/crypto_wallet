package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

const (
	serverAddress string = "http://localhost:3003"
)

type WalletBalanceResponse struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
}

// FIXME make as usecases and add context to timeout
func getWalletBalance(walletAddr string) (string, error) {

	params := url.Values{}
	params.Add("address", walletAddr)
	fullURL := fmt.Sprintf("%s?%s", serverAddress+"/get_wallet_balance", params.Encode())

	resp, err := http.Get(fullURL)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	var result WalletBalanceResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", fmt.Errorf("error parsing JSON: %v", err)
	}

	return result.Balance, nil
}
