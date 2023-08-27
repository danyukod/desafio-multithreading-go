package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type CepResponse struct {
	Result string
	API    string
}

func main() {

	c1 := make(chan CepResponse)
	c2 := make(chan CepResponse)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		api := "https://cdn.apicep.com/file/apicep/18230-000.json"
		resp, err := http.Get(api)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		result := CepResponse{
			Result: string(b),
			API:    api,
		}

		c1 <- result

		wg.Done()
	}()

	go func() {
		api := "http://viacep.com.br/ws/18230-000/json/"
		resp, err := http.Get(api)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		if err != nil {
			panic(err)
		}

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		result := CepResponse{
			Result: string(b),
			API:    api,
		}

		c2 <- result

		wg.Done()
	}()

	select {
	case resp := <-c1:
		fmt.Println(resp)
	case resp := <-c2:
		fmt.Println(resp)
	case <-time.After(1 * time.Second):
		print("Timed out")
	}

}
