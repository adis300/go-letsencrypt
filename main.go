package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var relativePath = ""
var letsencryptSecret = ""

// PORT implementation
const PORT = ":80"

func getRelativePath() string {
	if relativePath == "" {
		filename := os.Args[0] // get command line first parameter

		filedirectory := filepath.Dir(filename)

		relativePath, _ = filepath.Abs(filedirectory)
		return relativePath
	}
	return relativePath
}

func loadFile(filepath string) []byte {
	data, _ := ioutil.ReadFile(filepath)
	return data
}

func letsencrypt(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(letsencryptSecret))
}

func main() {
	letsencryptSecret = string(loadFile(getRelativePath() + "/letsencrypt.secret"))
	letsencryptSecret = strings.Replace(letsencryptSecret, "\n", "", -1)
	letsencryptSecret = strings.TrimSpace(letsencryptSecret)

	var challenge = strings.Split(letsencryptSecret, ".")[0]
	http.HandleFunc("/.well-known/acme-challenge/"+challenge, letsencrypt)

	log.Println("Starting Letsencrypt server on " + PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatal("Fatal error happened server on port" + PORT)
	}
}
