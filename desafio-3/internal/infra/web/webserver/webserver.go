package webserver

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]map[string]http.HandlerFunc // Método -> Rota -> Manipulador
	WebServerPort string
}

var Verbs = [...]string{"POST", "GET"}

func NewWebServer(serverPort string) *WebServer {

	handlers := make(map[string]map[string]http.HandlerFunc)
	for _, verb := range Verbs {
		handlers[verb] = make(map[string]http.HandlerFunc)
	}

	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      handlers,
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(verb string, path string, handler http.HandlerFunc) {

	verb = strings.ToUpper(verb)
	if s.Handlers[verb] == nil {
		s.Handlers[verb] = make(map[string]http.HandlerFunc)
	}
	s.Handlers[verb][path] = handler
	log.Println("Handler registrado:", verb, path) // Log para verificar registro
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for _, verb := range Verbs {
		for path, handler := range s.Handlers[verb] {
			log.Println("Registrando rota:", verb, path) // Log para depuração
			s.Router.Method(verb, path, handler)
		}
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
