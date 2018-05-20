package anagram

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetAnagrams gets all anagrams for a given work
func GetAnagrams(c *gin.Context) {
	name := c.Param("word")
	name = strings.TrimSuffix(name, ".json")
	key := SortString(name)
	// get all elements from key
	val, err := Client.LRange(key, 0, -1).Result()
	if err != nil {
		// TODO: improve error handling
		c.Error(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"anagrams": val,
	})
}

// UpdateCorpus updates the corpus with a list of words supplied by client
func UpdateCorpus(c *gin.Context) {
	var requestBody RequestBody
	err := c.BindJSON(&requestBody)
	if err != nil {
		panic(err)
	}
	length := len(requestBody.Words)
	for i := 0; i < length; i++ {
		word := requestBody.Words[i]
		key := SortString(word)
		err := Client.LPush(key, word).Err()
		if err != nil {
			// TODO: improve error handling
			c.Error(err)
		}
	}
	c.JSON(http.StatusCreated, gin.H{"user": requestBody})
}

// DeleteWord deletes word specified in path
func DeleteWord(c *gin.Context) {
	word := c.Param("word")
	word = strings.TrimSuffix(word, ".json")
	key := SortString(word)
	err := Client.LRem(key, 0, word).Err()
	if err != nil {
		// TODO: improve error handling
		c.Error(err)
	}
	c.Status(http.StatusNoContent)
}

// DropCorpus drops everthing in the corpus
func DropCorpus(c *gin.Context) {
	err := Client.FlushDBAsync().Err()
	status := http.StatusNoContent
	if err != nil {
		// TODO better error handling failed to drop corpus improve error handling
		c.Error(err)
	}
	c.Status(status)
}
