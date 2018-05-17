package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Words []string `json:"words"`
}

func getAnagrams(c *gin.Context) {
	name := c.Param("word")
	trimedName := strings.TrimSuffix(name, ".json")
	c.JSON(http.StatusOK, gin.H{
		"status":  "posted",
		"message": trimedName,
	})
}

func updateCorpus(c *gin.Context) {
	var requestBody RequestBody
	err := c.BindJSON(&requestBody)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusCreated, gin.H{"user": requestBody})
}

func main() {
	router := gin.Default()

	// get all anagrams for a given work
	router.GET("/anagrams/:word", getAnagrams)

	// update the corpus with a list of words supplied by client
	router.POST("/words.json", updateCorpus)

	router.Run()
}
