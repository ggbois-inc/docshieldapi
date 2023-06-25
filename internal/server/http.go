package server

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func CreateHttpServer() *http.Server {
	godotenv.Load()
	log.Printf("Launching server on %s", os.Getenv("HOST"))
	router := httprouter.New()
	createRoutes(router)
	return &http.Server{
		Addr:    os.Getenv("HOST"),
		Handler: router,
	}
}
