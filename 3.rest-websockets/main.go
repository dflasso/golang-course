package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/dflasso/rest-ws/handlers"
	"github.com/dflasso/rest-ws/middleware"
	"github.com/dflasso/rest-ws/server"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loadinf .env file")
	}

	PORT := os.Getenv("PORT")
	JWTSecret := os.Getenv("JWT_SECRET")
	DatabaseUrl := os.Getenv("DATABASE_URL")

	serverInstace, err := server.NewServer(context.Background(), &server.Config{
		JWTSecret:   JWTSecret,
		Port:        PORT,
		DatabaseUrl: DatabaseUrl,
	})

	if err != nil {
		log.Fatal(err)
	}

	serverInstace.Start(BindRoutes)
}

/*
En esta función se definen los endpoints
*/
func BindRoutes(s server.Server, r *mux.Router) {

	r.Use(middleware.CheckAuthMiddleware(s))

	/*
		rutas
	*/
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)

	//auth
	r.HandleFunc("/signup", handlers.SignUpHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/login", handlers.LoginHandle(s)).Methods(http.MethodPost)

	//user
	r.HandleFunc("/me", handlers.MeHandler(s)).Methods(http.MethodGet)
}
