package webapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/internal/entity"
	"net/http"
	"net/url"
)

type WebApiRepo struct {
	serverAddress string
}

func New(address string) *WebApiRepo {
	return &WebApiRepo{serverAddress: address}
}

func (repo *WebApiRepo) GetBalance(ctx context.Context, w entity.Wallet) ([]string, error) {
	type WalletBalanceResponse struct {
		Address string `json:"address"`
		Balance string `json:"balance"`
	}
	req_url := repo.serverAddress + "/get_wallet_balance"

	params := url.Values{}
	params.Add("address", w.Address)
	fullURL := fmt.Sprintf("%s?%s", req_url, params.Encode())

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
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}
	balance := parseAnswer(result.Balance)

	return balance, nil
}

func parseAnswer(ba string) []string {
	strSliceOutput := []string{}
	lastPos := 0
	for i := range ba {
		if ba[i] == '\n' {
			s := string(ba[lastPos : i+1])
			strSliceOutput = append(strSliceOutput, s)
			lastPos = i + 1
		}
	}

	return strSliceOutput
}
