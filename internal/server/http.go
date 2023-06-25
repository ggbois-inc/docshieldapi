package server

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func CreateHttpServer() *http.Server {
	godotenv.Load()
	return &http.Server{
		Addr: os.Getenv("HOST"),
	}
}
