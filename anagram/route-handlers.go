package anagram

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)
// GetAnagrams gets all anagrams for a given work
func GetAnagrams(c *gin.Context) {
	name := c.Param("word")
	trimedName := strings.TrimSuffix(name, ".json")
	// TODO: call db to get anagrams for trimmedName
	c.JSON(http.StatusOK, gin.H{
		"status":  "posted",
		"message": trimedName,
	})
}
// UpdateCorpus updates the corpus with a list of words supplied by client
func UpdateCorpus(c *gin.Context) {
	var requestBody RequestBody
	err := c.BindJSON(&requestBody)
	// TODO: call db to add words to corpus
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusCreated, gin.H{"user": requestBody})
}
// DeleteWord deletes word specified in path
func DeleteWord(c *gin.Context) {
	// TODO
	// name := c.Param("word")
	// trimedName := strings.TrimSuffix(name, ".json")
	c.Status(http.StatusNoContent)
}
// DropCorpus drops everthing in the corpus
func DropCorpus(c *gin.Context) {
	// TODO
	c.Status(http.StatusNoContent)
}
