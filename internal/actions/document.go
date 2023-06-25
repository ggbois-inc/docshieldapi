package actions

import (
	"io"
	"log"

	"github.com/ggbois-inc/docshieldapi/internal/database"
	"github.com/ggbois-inc/docshieldapi/internal/ipfs"
)

func UploadDocument(meta_id string, file io.Reader, filename string) {
	log.Printf("Running Upload action for %s", meta_id)
	database.CreateUser(meta_id)
	cid := ipfs.UploadFile(file)
	database.CreateDocument(meta_id, filename, cid, generateRandomString(10))
}

func GetDocuments(meta_id string) []database.Document {
	log.Printf("Running Get Files action for %s", meta_id)
	docs := database.GetDocuments(meta_id)
	return docs
}
