package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func homePage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello world\n")
}

func createRoutes(router *httprouter.Router) {
	router.GET("/", homePage)
}
