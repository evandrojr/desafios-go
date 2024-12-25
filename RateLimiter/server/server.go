package main

import (
	"log"
	"net/http"

	"github.com/evandrojr/ratelimiter/configs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	configs.LoadConfig()

	r := chi.NewRouter()

	// Middleware padrão do Chi
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Middleware personalizado
	r.Use(RequestLogger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Bem-vindo!"))
	})

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá, mundo!"))
	})

	log.Println("Servidor rodando em :8080")
	http.ListenAndServe(":8080", r)
}

// Middleware personalizado para log de requisições
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Método: %s, URL: %s, Host: %s", r.Method, r.URL.Path, r.Host)
		next.ServeHTTP(w, r)
	})
}
