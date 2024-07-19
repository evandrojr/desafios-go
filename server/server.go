package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/kr/pretty"
)

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

func logConsoleAndBrowser(msg string, w http.ResponseWriter) {
	log.Println(msg)
	w.Write([]byte(msg))
}

func getCotacao(ctxBg context.Context, w http.ResponseWriter, r *http.Request, data Cotacao) (Cotacao, error) {
	ctx, cancel := context.WithTimeout(ctxBg, 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		logConsoleAndBrowser("Erro ao preparar a requisição:"+err.Error(), w)
		return Cotacao{}, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logConsoleAndBrowser("Erro ao fazer a requisição:"+err.Error(), w)
		return Cotacao{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logConsoleAndBrowser("Erro ao ler a requisição:"+err.Error(), w)
		return Cotacao{}, err
	}
	// log.Println(string(body))

	err = json.Unmarshal(
		body,
		&data,
	)

	if err != nil {
		logConsoleAndBrowser("Erro json.Unmarshal:"+err.Error(), w)
		return Cotacao{}, err
	}

	// // jsonString:= fmt.Sprintf(res)
	fmt.Fprint(w, string(body))
	return data, nil
}

func cotacaoHandler(w http.ResponseWriter, r *http.Request) {

	var data Cotacao
	ctxBg := context.Background()

	data, err := getCotacao(ctxBg, w, r, data)
	if err != nil {
		logConsoleAndBrowser("Erro getCotacao:"+err.Error(), w)
	}
	pretty.Println(data)

}

func main() {
	port := "8080"
	mux := http.NewServeMux()
	mux.HandleFunc("GET /cotacao/", cotacaoHandler)
	pretty.Println("Rodando na porta:", port)
	http.ListenAndServe("localhost:"+port, mux)
}
