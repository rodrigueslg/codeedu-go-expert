package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	APITimeout = 300 * time.Millisecond
)

func RunClient() {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}

	ctxAPI, cancelAPI := context.WithTimeout(req.Context(), APITimeout)
	defer cancelAPI()

	req = req.WithContext(ctxAPI)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	price := string(body)

	file, err := os.Create("./client/cotacao.txt")
	if err != nil {
		fmt.Printf("Can't create file:\n")
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Dólar: %s\n", price))
	if err != nil {
		panic(err)
	}
}
