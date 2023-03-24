package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type ApiCEP struct {
	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	ch1 := make(chan ApiCEP)
	ch2 := make(chan ViaCEP)

	cep := os.Args[1]
	if cep == "" {
		println("CEP is required.\nTry again: 'go run main.go xxxxx-xxx' ")
	}
	if !strings.Contains(cep, "-") {
		println("Invalid CEP format.\nTry again: 'go run main.go xxxxx-xxx'")
	}

	go func() {
		url := fmt.Sprintf("https://cdn.apicep.com/file/apicep/%s.json", cep)
		println("requesting apicep: ", url)

		req, _ := http.NewRequest(http.MethodGet, url, nil)
		res, _ := http.DefaultClient.Do(req)
		body, _ := ioutil.ReadAll(res.Body)
		if res.StatusCode != 200 {
			println("viacep response: ", string(body))
		}
		if res.StatusCode == 429 {
			println("viacep response: ", "blocked by apicep (too many requests)")
		}

		var apiCep ApiCEP
		_ = json.Unmarshal(body, &apiCep)

		ch1 <- apiCep
	}()

	go func() {
		url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", strings.Replace(cep, "-", "", 1))
		println("requesting viacep: ", url)

		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://viacep.com.br/ws/%s/json/", strings.Replace(cep, "-", "", 1)), nil)
		res, _ := http.DefaultClient.Do(req)
		body, _ := ioutil.ReadAll(res.Body)
		if res.StatusCode != 200 {
			println("viacep response: ", string(body))
		}
		if res.StatusCode == 429 {
			println("viacep response: ", "blocked by viacep (too many requests)")
		}

		var viaCep ViaCEP
		_ = json.Unmarshal(body, &viaCep)

		ch2 <- viaCep
	}()

	select {
	case apicep := <-ch1:
		fmt.Printf("apicep returned first: %v\n", apicep)
	case viacep := <-ch2:
		fmt.Printf("viacep returned first: %v\n", viacep)
	case <-time.After(time.Second):
		println("timeout")
	}
}
