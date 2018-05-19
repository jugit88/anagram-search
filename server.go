package main

import (
	"app/anagram"

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

	router.Run()

}
