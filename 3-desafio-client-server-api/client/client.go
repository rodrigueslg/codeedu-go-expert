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
		fmt.Printf("client: can't create request: %s\n", err)
		return
	}

	ctxAPI, cancelAPI := context.WithTimeout(req.Context(), APITimeout)
	defer cancelAPI()

	req = req.WithContext(ctxAPI)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: request failed: %s\n", err)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: can't read api response: %s\n", err)
		return
	}

	price := string(body)

	file, err := os.Create("./client/cotacao.txt")
	if err != nil {
		fmt.Printf("client: can't create file: %s\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Dólar: %s\n", price))
	if err != nil {
		fmt.Printf("client: can't write file: %s\n", err)
		return
	}
}
