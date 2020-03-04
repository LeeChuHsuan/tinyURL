package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const domainName = "localhost:8000"

type tinyURL struct {
	URL   string `json:"url"`
	Error error  `json:"error,omitempty"`
}

func newtinyURL(hashval string, err error) tinyURL {
	newURL := domainName + "/" + hashval
	return tinyURL{newURL, err}
}

func hashURL(url string) string {
	hasher := sha256.New()
	hasher.Write([]byte(url))
	hashvalue := hex.EncodeToString(hasher.Sum(nil))
	return hashvalue[:8]
}

func getIndexPage(c *gin.Context) {
	http.ServeFile(c.Writer, c.Request, "../web/index.html")
}

func redirectHandler(c *gin.Context) {
	hashval := c.Param("hashval")
	var record *urlMapping
	record, err := record.GetByPrimaryKey(hashval)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
	}
	http.Redirect(c.Writer, c.Request, record.URL, http.StatusFound)
}

func postHandler(c *gin.Context) {
	url := c.PostForm("url")
	URLHash := hashURL(url)
	response, err := json.Marshal(newtinyURL(URLHash, nil))
	if err != nil {
		fmt.Fprint(c.Writer, "{\"error\": \"%v\"}", err)
		return
	}

	newRecord := newurlMapping(url, URLHash)
	err = newRecord.InsertDB()
	if err != nil {
		fmt.Fprint(c.Writer, "{\"error\": \"%v\"}", err)
		return
	}
	c.Writer.WriteHeader(http.StatusCreated)
	fmt.Fprintf(c.Writer, "%s", response)
}

func main() {
	router := gin.Default()
	router.GET("/", getIndexPage)
	router.GET("/:hashval", redirectHandler)
	router.POST("/", postHandler)
	router.Run("localhost:8000")
}
