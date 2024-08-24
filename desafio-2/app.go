package main

import (
	"fmt"
	"net/http"
	"os"
)

const cep = "41830430"

func main() {

	requestURL := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

}

// func cep1() {

// }

// func cep2() {

// }
