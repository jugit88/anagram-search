package main

import (
	"flag"
	"fmt"
	"os"
	// "strings"
	"app/anagram"

	"github.com/gin-gonic/gin"
)

func main() {
	// check corpus file is passed on command line
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Input file is missing.")
		os.Exit(1)
	}
	fileName := os.Args[0]
	// populate cache on separate goroutine
	go anagram.ReadLines(fileName)

	// initialize http server/middleware
	router := gin.Default()

	// routes
	router.GET("/anagrams/:word", anagram.GetAnagrams)

	router.POST("/words.json", anagram.UpdateCorpus)

	router.DELETE("/words/:word", anagram.DeleteWord)

	router.DELETE("/words.json", anagram.DropCorpus)

	router.Run()
}
