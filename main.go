package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type tinyURL struct {
	URL   string `json:"url"`
	Error error  `json:"error,omitempty"`
}

func shortURL(url string) string {
	hasher := sha256.New()
	hasher.Write([]byte(url))
	hashvalue := hex.EncodeToString(hasher.Sum(nil))
	return hashvalue[:8]
}

func handler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		if req.URL.Path == "/" {
			http.ServeFile(w, req, "./index.html")
		} else {
			hashval := req.URL.Path[1:]
			url, err := GetURLMapping(hashval)
			if err != nil{
				http.NotFound(w,req)
			}
			http.Redirect(w, req, url, http.StatusFound)
		}
	} else if req.Method == "POST" {
		if err := req.ParseForm(); err != nil {
			fmt.Fprint(w, "{\"error\": \"%v\"}", err)
			return
		}
		url := req.FormValue("url")
		newURL := shortURL(url)
		result, err := json.Marshal(tinyURL{domainName + "/" + newURL, nil})
		if err != nil {
			fmt.Fprint(w, "{\"error\": \"%v\"}", err)
			return
		}
		InsertURLMapping(url, newURL)
		fmt.Fprintf(w, "%s", result)
	}
}

const domainName = "localhost:8000"

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
