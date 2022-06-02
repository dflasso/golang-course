package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/dflasso/rest-ws/database"
	"github.com/dflasso/rest-ws/repository"
	"github.com/gorilla/mux"
)

type Config struct {
	Port        string
	JWTSecret   string
	DatabaseUrl string
}

/*
Los servidores deben tener la configuración señalada
*/
type Server interface {
	Config() *Config
}

/*
Encargado de gestionar las instancias del tipo Server
*/
type Broker struct {
	config *Config
	router *mux.Router
}

/*
Implementa la interface
*/
func (b *Broker) Config() *Config {
	return b.config
}

/*
Crear una instancia del servidor
Valida que las configuraciones no sean nulas
*/
func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}

	if config.JWTSecret == "" {
		return nil, errors.New("jwt secret is required")
	}

	if config.DatabaseUrl == "" {
		return nil, errors.New("db url secret is required")
	}

	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}

	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)

	repo, err := database.NewPostgresRepository(b.config.DatabaseUrl)

	if err != nil {
		log.Fatal(err)
	}

	repository.SetRepository(repo)

	log.Println("Starting server on port: ", b.Config().Port)

	if err := http.ListenAndServe(b.config.Port, b.router); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
