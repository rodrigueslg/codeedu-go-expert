package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	APITimeout        = 200 * time.Millisecond
	RepositoryTimeout = 10 * time.Millisecond
)

type CoinService struct {
	repo *CoinRepository
}

type CoinPrice struct {
	UsdBrl struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func NewService(repo *CoinRepository) *CoinService {
	return &CoinService{
		repo: repo,
	}
}

func (s *CoinService) GetUsdBrlPrice() (*CoinPrice, error) {
	req, err := http.NewRequest(http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	ctxAPI, cancelAPI := context.WithTimeout(req.Context(), APITimeout)
	defer cancelAPI()

	req = req.WithContext(ctxAPI)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "server: request failed: %s", err)
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "server: can't read api respose: %s", err)
		return nil, err
	}

	var cp CoinPrice
	err = json.Unmarshal(body, &cp)
	if err != nil {
		fmt.Fprintf(os.Stderr, "server: can't parse api respose: %s", err)
		return nil, err
	}

	ctxRepo, cancelRepo := context.WithTimeout(ctxAPI, RepositoryTimeout)
	defer cancelRepo()
	s.repo.Log(ctxRepo, &cp)

	return &cp, nil
}
