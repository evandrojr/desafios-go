package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/kr/pretty"
)

// const url = "http://localhost:8080"

const url = "https://economia.awesomeapi.com.br/json/last/USD-BRL"

type Cotacao struct {
	Usdbrl struct {
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

func main() {

	req, err := http.Get(url)
	if err != nil {
		log.Fatal("Erro ao fazer a requisição:" + err.Error())
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal("Erro ao fazer a requisição:" + err.Error())
	}
	// fmt.Println(string(res))

	var data Cotacao
	err = json.Unmarshal(res, &data)
	pretty.Println(data)
}
