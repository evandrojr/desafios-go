package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// const url = "http://localhost:8080"

const url = "https://evandrojr.org"

func main() {

	req, err := http.Get(url)
	if err != nil {
		log.Fatal("Erro ao fazer a requisição:" + err.Error())
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	fmt.Println(req.Body)
	if err != nil {
		log.Fatal("Erro ao fazer a requisição:" + err.Error())
	}
	fmt.Println(string(res))
}
