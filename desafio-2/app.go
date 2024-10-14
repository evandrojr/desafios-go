package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const cep = "41830460"

func requestCep(url string, ret chan Ret) {

	start := time.Now()

	res, err := http.Get(url)
	if err != nil {
		log.Printf("error making http request: %s\n", err)
		ret <- Ret{Err: err.Error()}
		return
	}

	// fmt.Printf("client: status code: %d\n", res.StatusCode)

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		// log.Println(res.StatusCode, err.Error())
		ret <- Ret{Err: err.Error(), HttpCode: res.StatusCode}
		return
	}
	// fmt.Println(string(body))
	t := time.Now()

	ret <- Ret{
		HttpCode: res.StatusCode,
		Body:     string(body),
		Elapsed:  t.Sub(start),
	}
}

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

func main() {

	brasilapi := make(chan Ret)
	viacep := make(chan Ret)
	requestURL := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	go requestCep(requestURL, brasilapi)

	requestURL = fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	go requestCep(requestURL, viacep)

	select {

	case ret := <-brasilapi:
		fmt.Println("Brasil Api foi mais rápida!")
		body := BrasilApiBody{}
		if err := json.Unmarshal([]byte(ret.Body), &body); err != nil {
			panic(err)
		}
		fmt.Println(body.Street+",", body.Neighborhood+",", body.City+",", body.State+", CEP:", body.Cep)

	case ret := <-viacep:
		fmt.Println("Viacep foi mais rápida:\n", ret.Body)
		body := ViacepBody{}
		if err := json.Unmarshal([]byte(ret.Body), &body); err != nil {
			panic(err)
		}
		fmt.Println(body.Logradouro+",", body.Unidade+",", body.Complemento+",", body.Localidade+",", body.Uf+", CEP:", body.Cep)

	case <-time.After(1 * time.Second):
		fmt.Println("Timeout")

	}

}
