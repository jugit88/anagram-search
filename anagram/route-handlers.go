package anagram

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetAnagrams gets all anagrams for a given work
func GetAnagrams(c *gin.Context) {
	word := c.Param("word")
	word = strings.TrimSuffix(word, ".json")
	key, normalized := GenerateKey(word)
	// get all elements from key
	result, err := Client.SMembers(key).Result()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, failureMessage("could not read from cache"))
	}
	anagrams := &Anagrams{Words: result}
	// handle limit query param
	limit, exists := c.GetQuery("limit")
	if exists {
		anagrams.applyLimit(limit)
	}
	// remove word from response
	anagrams.removeWord(normalized)

	c.JSON(http.StatusOK, gin.H{
		"anagrams": anagrams.Words,
	})
}

// UpdateCorpus updates the corpus with a list of words supplied by client
func UpdateCorpus(c *gin.Context) {
	var requestBody RequestBody
	words, err := requestBody.parseRequestBody(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, failureMessage("request body was missing or malformed"))
	}
	for _, w := range words {
		key, normalized := GenerateKey(w)
		err := Client.SAdd(key, normalized).Err()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, failureMessage("could not write to cache"))
		}
	}
	c.Status(http.StatusCreated)
}

// DeleteWord deletes word specified in path
func DeleteWord(c *gin.Context) {
	word := c.Param("word")
	word = strings.TrimSuffix(word, ".json")
	key, normalized := GenerateKey(word)
	err := Client.SRem(key, 0, normalized).Err()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, failureMessage("could not delete from cache"))
	}
	c.Status(http.StatusNoContent)
}

// DropCorpus drops everthing in the corpus
func DropCorpus(c *gin.Context) {
	err := Client.FlushDBAsync().Err()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, failureMessage("could not flush the cache"))
	}
	c.Status(http.StatusNoContent)
}

// IsAnagram checks to see if a given set of words are anagrams of eachother
func IsAnagram(c *gin.Context) {
	var requestBody RequestBody
	words, err := requestBody.parseRequestBody(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, failureMessage("request body was missing or malformed"))
	}
	isAnagram := true
	// sort each word and check if they are equal
	for i := 1; i < len(words); i++ {
		w1, _ := GenerateKey(words[i-1])
		w2, _ := GenerateKey(words[i])
		if w1 != w2 {
			isAnagram = false
		}
	}
	c.JSON(http.StatusOK, gin.H{"isSetAnagrams": isAnagram})
}

// DeleteKey deletes the key and all assocated values from the database
func DeleteKey(c *gin.Context) {
	word := c.Param("word")
	key, _ := GenerateKey(word)
	// async delete
	err := Client.Unlink(key).Err()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, failureMessage("could not delete from cache"))
	}
	c.Status(http.StatusNoContent)
}

func failureMessage(message string) gin.H {
	return gin.H{"error": message}
}

func (anagrams *Anagrams) removeWord(word string) {
	for i, w := range anagrams.Words {
		if w == word {
			anagrams.Words = append(anagrams.Words[:i], anagrams.Words[i+1:]...)
		}
	}
}

func (requestBody *RequestBody) parseRequestBody(c *gin.Context) ([]string, error) {
	err := c.BindJSON(&requestBody)
	return requestBody.Words, err
}

func (anagrams *Anagrams) applyLimit(limit string) {
	i, err := strconv.Atoi(limit)
	if err != nil || i < 0 {
		log.Println("limit must be a whole number, ignoring limit")
	} else {
		// this ensures index can never be out of range
		if i <= len(anagrams.Words) {
			anagrams.Words = anagrams.Words[:i]
		}
	}
}
