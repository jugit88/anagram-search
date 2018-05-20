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
	err := anagram.Client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	router.GET("/anagrams/:word", anagram.GetAnagrams)

	router.POST("/words.json", anagram.UpdateCorpus)

	router.DELETE("/words/:word", anagram.DeleteWord)

	router.DELETE("/words.json", anagram.DropCorpus)

	router.GET("/system_health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	router.Run()

}
