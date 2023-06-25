package actions

import (
	"crypto/rand"
	"encoding/base64"
	"log"

	"github.com/ggbois-inc/docshieldapi/internal/database"
	"github.com/ggbois-inc/docshieldapi/internal/ipfs"
)

func generateRandomString(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}
	return base64.URLEncoding.EncodeToString(bytes)[:length]
}

func GetDocumentByLink(shortcode string) []byte {
	doc := database.GetDocumentByCode(shortcode)
	file := ipfs.GetFile(doc.CID)
	return file
}
