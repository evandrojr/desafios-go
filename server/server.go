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

func getCotacao(ctxBg context.Context, w http.ResponseWriter, r *http.Request, data *Cotacao) {
	ctx, cancel := context.WithTimeout(ctxBg, 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		logConsoleAndBrowser("Erro ao preparar a requisição:"+err.Error(), w)
		return
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logConsoleAndBrowser("Erro ao fazer a requisição:"+err.Error(), w)
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logConsoleAndBrowser("Erro ao ler a requisição:"+err.Error(), w)
		return
	}
	// log.Println(string(body))

	err = json.Unmarshal(body, &data)
	pretty.Println(*data)

	// // jsonString:= fmt.Sprintf(res)
	fmt.Fprint(w, string(body))
}

func cotacaoHandler(w http.ResponseWriter, r *http.Request) {

	ctxBg := context.Background()
	var data *Cotacao

	getCotacao(ctxBg, w, r, data)

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /cotacao/", cotacaoHandler)
	http.ListenAndServe("localhost:8080", mux)
}
