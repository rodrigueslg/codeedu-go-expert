package server

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type CoinHandler struct {
	service *CoinService
}

func NewCoinHandler(s *CoinService) *CoinHandler {
	h := &CoinHandler{
		service: s,
	}
	return h
}

func (h *CoinHandler) GetUsdBrlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	cp, err := h.service.GetUsdBrlPrice()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	price, err := strconv.ParseFloat(cp.UsdBrl.Bid, 32)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(price)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
