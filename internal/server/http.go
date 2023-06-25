package server

import (
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"

	"github.com/rs/cors"

	"github.com/joho/godotenv"
)

func CreateHttpServer() *http.Server {
	godotenv.Load()
	router := httprouter.New()
	createRoutes(router)
	handler := cors.AllowAll().Handler(router)
	return &http.Server{
		Addr:    os.Getenv("HOST"),
		Handler: handler,
	}
}
