package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
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

const domainName = "localhost:8000"

func getIndexPage(c *gin.Context){
	http.ServeFile(c.Writer, c.Request, "../web/index.html")
}

func redirectHandler(c *gin.Context){
	hashval := c.Param("hashval")
	url, err := GetURLMapping(hashval)
	if err != nil{
		http.NotFound(c.Writer, c.Request)
	}
	http.Redirect(c.Writer, c.Request, url, http.StatusFound)
}

func postHandler(c *gin.Context){
	url := c.PostForm("url")
	newURL := shortURL(url)
	result, err := json.Marshal(tinyURL{domainName + "/" + newURL, nil})
	if err != nil {
		fmt.Fprint(c.Writer, "{\"error\": \"%v\"}", err)
		return
	}
	InsertURLMapping(url, newURL)
	fmt.Fprintf(c.Writer, "%s", result)
}

func main() {
	router := gin.Default()
	router.GET("/", getIndexPage)
	router.GET("/:hashval", redirectHandler)
	router.POST("/", postHandler)
	router.Run("localhost:8000")
}
