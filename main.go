package main

import (
	server "github.com/ggbois-inc/docshieldapi/internal/server"
)

func main() {
	server.CreateHttpServer().ListenAndServe()
}
