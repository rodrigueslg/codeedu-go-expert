package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		req, _ := http.NewRequest(http.MethodGet, "https://cdn.apicep.com/file/apicep/"+"31744596"+".json", nil)
		res, _ := http.DefaultClient.Do(req)
		body, _ := ioutil.ReadAll(res.Body)

		ch1 <- string(body)
	}()

	go func() {
		req, _ := http.NewRequest(http.MethodGet, "http://viacep.com.br/ws/"+"31744596"+"/json/", nil)
		res, _ := http.DefaultClient.Do(req)
		body, _ := ioutil.ReadAll(res.Body)

		ch2 <- string(body)
	}()

	select {
	case apicep := <-ch1:
		fmt.Printf("apicep returned first: %s\n", apicep)
	case viacep := <-ch2:
		fmt.Printf("viacep returned first: %s\n", viacep)
	case <-time.After(time.Second):
		println("timeout")
	}
}
