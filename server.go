package main

import (
	"app/anagram"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// initialize http server/middleware
	router := gin.Default()

	// routes
	router.GET("/anagrams/:word", anagram.GetAnagrams)

	router.POST("/words.json", anagram.UpdateCorpus)

	router.DELETE("/words/:word", anagram.DeleteWord)

	router.DELETE("/words.json", anagram.DropCorpus)

	// check if a given set of words are anagrams
	router.POST("/words/check", anagram.IsAnagram)

	// health check
	router.GET("/system_health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	router.Run()

}
