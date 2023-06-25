package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ggbois-inc/docshieldapi/internal/actions"
	driver "github.com/ggbois-inc/docshieldapi/internal/database"
	"github.com/julienschmidt/httprouter"
)

func homePage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := driver.CreateUser("xyz")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

func fileUpload(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	actions.UploadDocument(r.Header.Get("meta_id"), file, fileHeader.Filename)
	fmt.Fprint(w, "DONE\n")
}

func getFiles(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	docs := actions.GetDocuments(r.Header.Get("meta_id"))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(docs)
}

func getFileByCode(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	data, doc := actions.GetDocumentByLink(p.ByName("code"))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", doc.Filename))
	http.ServeContent(w, r, doc.Filename, doc.CreatedOn, bytes.NewReader(data))
}

func createRoutes(router *httprouter.Router) {
	router.GET("/", homePage)
	router.POST("/api/upload", fileUpload)
	router.GET("/api/files", getFiles)
	router.GET("/api/file/:code", getFileByCode)
}
