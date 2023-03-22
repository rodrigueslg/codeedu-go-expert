package server

import (
	"net/http"
)

type CoinServer struct {
	mux         *http.ServeMux
	coinHandler *CoinHandler
}

func NewServer(ch *CoinHandler) *CoinServer {
	return &CoinServer{
		mux:         http.NewServeMux(),
		coinHandler: ch,
	}
}

func (s *CoinServer) Run() error {
	s.mux.HandleFunc("/cotacao", s.coinHandler.GetUsdBrlHandler)

	return http.ListenAndServe(":8080", s.mux)
}
