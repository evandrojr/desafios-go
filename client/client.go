package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/kr/pretty"
)

// const url = "http://localhost:8080"

const url = "http://localhost:8080/cotacao"
const requestTimeout = 300 * time.Millisecond

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

	var data Cotacao
	ctxRequisicao := context.Background()

	_, err := getCotacao(ctxRequisicao, &data)
	if err != nil {
		log.Println("Erro getCotacao: " + err.Error())
		return
	} else {
		pretty.Println(data.Usdbrl.Bid)
	}
	conteudo := "Dólar: " + data.Usdbrl.Bid

	err = os.WriteFile("cotacao.txt", []byte(conteudo), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("JSON escrito no arquivo com sucesso!")
}

// req, err := http.Get(url)
// if err != nil {
// 	log.Fatal("Erro ao fazer a requisição:" + err.Error())
// }
// defer req.Body.Close()

// res, err := io.ReadAll(req.Body)
// if err != nil {
// 	log.Fatal("Erro ao fazer a requisição:" + err.Error())
// }
// // fmt.Println(string(res))

// var data Cotacao
// err = json.Unmarshal(res, &data)
// pretty.Println(data)
// }

func getCotacao(ctxBg context.Context, data *Cotacao) (*Cotacao, error) {
	ctx, cancel := context.WithTimeout(ctxBg, requestTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		errPersonalizado := errors.New("Erro ao preparar a requisição: " + err.Error())
		return &Cotacao{}, errPersonalizado
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		errPersonalizado := errors.New("Erro ao fazer a requisição: " + err.Error())
		return &Cotacao{}, errPersonalizado
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		errPersonalizado := errors.New("Erro ao ler a requisição:" + err.Error())
		return &Cotacao{}, errPersonalizado
	}
	pretty.Println(string(body))

	err = json.Unmarshal(
		body,
		data,
	)

	if err != nil {
		errPersonalizado := errors.New("Erro json.Unmarshal: " + err.Error() + err.Error())
		return &Cotacao{}, errPersonalizado
	}

	// // jsonString:= fmt.Sprintf(res)

	return data, nil
}
