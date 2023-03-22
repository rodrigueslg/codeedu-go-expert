package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

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

func NewServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/usd-brl", GetUsdBrlHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func GetUsdBrlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	cp, err := GetUsdBrlPrice()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cp)
}

func GetUsdBrlPrice() (*CoinPrice, error) {
	res, err := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var cp CoinPrice
	err = json.Unmarshal(body, &cp)
	if err != nil {
		return nil, err
	}

	return &cp, nil
}
