package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/evandrojr/desafio-go-1/shared"
	"github.com/kr/pretty"
)

// const url = "http://localhost:8080"

const url = "http://localhost:8080/cotacao"
const requestTimeout = 300 * time.Millisecond

func main() {

	var data shared.Cotacao
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

	log.Println("JSON escrito no arquivo com sucesso!")
}

func prepareRequest(ctx context.Context) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, errors.New("Erro ao preparar a requisição: " + err.Error())
	}
	return req, nil
}

func makeRequest(req *http.Request) ([]byte, error) {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.New("Erro ao fazer a requisição: " + err.Error())
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("Erro ao ler a requisição: " + err.Error())
	}
	return body, nil
}

func decodeCotacao(body []byte, data *shared.Cotacao) error {
	if err := json.Unmarshal(body, data); err != nil {
		return errors.New("Erro json.Unmarshal: " + err.Error())
	}
	return nil
}

func getCotacao(ctxBg context.Context, data *shared.Cotacao) (*shared.Cotacao, error) {
	ctx, cancel := context.WithTimeout(ctxBg, requestTimeout)
	defer cancel()

	req, err := prepareRequest(ctx)
	if err != nil {
		return &shared.Cotacao{}, err
	}

	body, err := makeRequest(req)
	if err != nil {
		return &shared.Cotacao{}, err
	}

	log.Println(string(body))

	err = json.Unmarshal(
		body,
		data,
	)

	decodeCotacao(body, data)
	if err != nil {
		return &shared.Cotacao{}, err
	}

	// // jsonString:= fmt.Sprintf(res)

	return data, nil
}
