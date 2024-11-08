package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const cep = "41830460"

type Ret struct {
	Err      string
	HttpCode int
	Body     string
	Elapsed  time.Duration
}

type BrasilApiBody struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type ViacepBody struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func requestCep(ctx context.Context, url string, ret chan<- Ret) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		ret <- Ret{Err: fmt.Sprintf("Erro ao criar requisição: %s", err)}
		return
	}

	start := time.Now()
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		ret <- Ret{Err: fmt.Sprintf("Erro na requisição HTTP: %s", err)}
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		ret <- Ret{Err: fmt.Sprintf("Erro ao ler o corpo da resposta: %s", err), HttpCode: res.StatusCode}
		return
	}

	ret <- Ret{
		HttpCode: res.StatusCode,
		Body:     string(body),
		Elapsed:  time.Since(start),
	}
}

func printResult(ret Ret, apiName string) {
	if apiName == "BrasilAPI" {
		body := BrasilApiBody{}
		if err := json.Unmarshal([]byte(ret.Body), &body); err != nil {
			log.Fatalf("Erro ao decodificar resposta BrasilAPI: %s", err)
		}
		fmt.Printf("BrasilAPI: %s, %s, %s - %s, CEP: %s\n", body.Street, body.Neighborhood, body.City, body.State, body.Cep)
	} else if apiName == "ViaCEP" {
		body := ViacepBody{}
		if err := json.Unmarshal([]byte(ret.Body), &body); err != nil {
			log.Fatalf("Erro ao decodificar resposta ViaCEP: %s", err)
		}
		fmt.Printf("ViaCEP: %s, %s, %s - %s, CEP: %s\n", body.Logradouro, body.Bairro, body.Localidade, body.Uf, body.Cep)
	}
}

func main() {
	brasilapiChan := make(chan Ret)
	viacepChan := make(chan Ret)

	// Foi usado Context para cancelar e liberar o recurso da requisição que fosse mais lenta
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go requestCep(ctx, fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep), brasilapiChan)
	go requestCep(ctx, fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep), viacepChan)

	// Tem um for no Select para esperar até 2 resultados, pois o primeiro resultado poderia ser de falha.
	// Se o primeiro resultado for de falha, o segundo resultado poderia ser de sucesso.
	for i := 0; i < 2; i++ {
		select {
		case ret := <-brasilapiChan:
			if ret.Err != "" {
				log.Printf("Erro BrasilAPI: %s", ret.Err)
			} else {
				fmt.Println("BrasilAPI foi mais rápida!")
				printResult(ret, "BrasilAPI")
				cancel()
				return
			}

		case ret := <-viacepChan:
			if ret.Err != "" {
				log.Printf("Erro ViaCEP: %s", ret.Err)
			} else {
				fmt.Println("ViaCEP foi mais rápida!")
				printResult(ret, "ViaCEP")
				cancel()
				return
			}

		case <-time.After(1 * time.Second):
			fmt.Println("Timeout: Nenhuma das APIs respondeu dentro do limite de 1 segundo.")
			cancel()
			return
		}
	}

	fmt.Println("Ambas as requisições falharam.")
}
