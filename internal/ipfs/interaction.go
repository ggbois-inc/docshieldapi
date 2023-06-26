package ipfs

import (
	"bytes"
	"io"
	"log"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/joho/godotenv"
)

var sh *shell.Shell

func init() {
	godotenv.Load()
	sh = shell.NewShell("localhost:5001")
}

func UploadFile(reader io.Reader) string {
	content, _ := encryptReader([]byte(os.Getenv("PASS")), reader)
	r, err := sh.Add(bytes.NewReader(content))
	if err != nil {
		log.Fatal(err)
	}
	sh.Pin(r)
	log.Println("File Uploaded")
	return r
}

func GetFile(cid string) []byte {
	i, _ := sh.Cat(cid)
	resp, _ := io.ReadAll(i)
	resp, _ = decryptReader([]byte(os.Getenv("PASS")), resp)
	return resp
}
