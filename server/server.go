package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/evandrojr/desafio-go-1/shared"
	"github.com/kr/pretty"
	_ "github.com/mattn/go-sqlite3"
)

const url = "https://economia.awesomeapi.com.br/json/last/USD-BRL"

var db *sql.DB

const requestTimeout = 200 * time.Millisecond
const dbTimeout = 10 * time.Millisecond

func initializeDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "cotacoes.db")
	if err != nil {
		return nil, err
	}

	query := `
    CREATE TABLE IF NOT EXISTS cotacoes (
        code TEXT,
		codein TEXT,
		name TEXT,
		high TEXT,
		low TEXT,
		varBid TEXT,
		pctChange TEXT,
		bid TEXT,
		ask TEXT,
		timestamp TEXT,
		create_date TEXT
    );`
	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func logErrorConsoleAndBrowser(msg string, w http.ResponseWriter) {
	log.Println(msg)
	http.Error(w, msg, http.StatusInternalServerError)
}

func getCotacao(ctxBg context.Context, data *shared.Cotacao) (*shared.Cotacao, error) {
	ctx, cancel := context.WithTimeout(ctxBg, requestTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		errPersonalizado := errors.New("Erro ao preparar a requisição: \n" + err.Error())
		return &shared.Cotacao{}, errPersonalizado
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		errPersonalizado := errors.New("Erro ao fazer a requisição: \n" + err.Error())
		return &shared.Cotacao{}, errPersonalizado
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		errPersonalizado := errors.New("Erro ao ler a requisição: \n" + err.Error())
		return &shared.Cotacao{}, errPersonalizado
	}

	err = json.Unmarshal(
		body,
		data,
	)

	if err != nil {
		errPersonalizado := errors.New("Erro json.Unmarshal: \n" + err.Error() + "\n" + err.Error())
		return &shared.Cotacao{}, errPersonalizado
	}

	// // jsonString:= fmt.Sprintf(res)

	return data, nil
}

func saveCotacao(ctx context.Context, data *shared.Cotacao) error {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	err := insertWithTimeout(ctx, db, data)
	if err != nil {
		log.Println("Erro ao inserir:", err)
	} else {
		log.Println("Inserido com sucesso")
	}
	return nil
}

func insertWithTimeout(ctx context.Context, db *sql.DB, data *shared.Cotacao) error {
	query := `INSERT INTO cotacoes (code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	c := *data
	_, err := db.ExecContext(ctx, query, c.Usdbrl.Code, c.Usdbrl.Codein, c.Usdbrl.Name, c.Usdbrl.High, c.Usdbrl.Low, c.Usdbrl.VarBid, c.Usdbrl.PctChange, c.Usdbrl.Bid, c.Usdbrl.Ask, c.Usdbrl.Timestamp, c.Usdbrl.CreateDate)
	return err
}

func cotacaoHandler(w http.ResponseWriter, r *http.Request) {

	var data shared.Cotacao
	ctxRequisicao := context.Background()

	_, err := getCotacao(ctxRequisicao, &data)
	if err != nil {
		logErrorConsoleAndBrowser("Erro getCotacao: "+err.Error(), w)

		return
	} else {
		pretty.Println(data)
	}
	// fmt.Fprint(
	// 	w,
	// 	string(body),
	// )

	// Criei um novo contexto, pois se usasse o outro acho que poderia dar problema com 2 timeouts
	ctxSalvamento := context.Background()
	err = saveCotacao(ctxSalvamento, &data)
	if err != nil {
		logErrorConsoleAndBrowser("Erro SaveCotacao: "+err.Error(), w)
	}

	jsonCotacao, err := json.Marshal(data)
	if err != nil {
		logErrorConsoleAndBrowser("Erro no json.Marshal da cotação: "+err.Error(), w)
	}

	fmt.Fprint(
		w,
		string(jsonCotacao),
	)

}

func main() {
	var err error
	db, err = initializeDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	port := "8080"
	mux := http.NewServeMux()
	mux.HandleFunc("GET /cotacao/", cotacaoHandler)
	pretty.Println("Rodando na porta:", port)

	err = http.ListenAndServe("localhost:"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
	pretty.Println("Saindo")
}
