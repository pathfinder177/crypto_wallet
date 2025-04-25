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

func parseAnswer(joined string) []string {
	strSliceOutput := []string{}
	lastPos := 0
	for i := range joined {
		if joined[i] == '\n' {
			s := string(joined[lastPos : i+1])
			strSliceOutput = append(strSliceOutput, s)
			lastPos = i + 1
		}
	}

	return strSliceOutput
}
