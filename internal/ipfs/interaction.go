package ipfs

import (
	"io"
	"log"

	shell "github.com/ipfs/go-ipfs-api"
)

var sh *shell.Shell

func init() {
	sh = shell.NewShell("localhost:5001")
}

func UploadFile(reader io.Reader) string {
	r, err := sh.Add(reader)
	if err != nil {
		log.Fatal(err)
	}
	sh.Pin(r)
	log.Println("File Uploaded")
	return r
}
