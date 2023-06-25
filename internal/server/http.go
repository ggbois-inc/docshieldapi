package server

import (
	"log"
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
	log.Printf("Starting server on %s", os.Getenv("HOST"))
	return &http.Server{
		Addr:    os.Getenv("HOST"),
		Handler: handler,
	}
}
