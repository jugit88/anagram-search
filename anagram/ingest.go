package anagram

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

func init() {
	Client = RedisClient()
	// read in file
	file := "dictionary.txt"
	words, err := readLines(file)
	if err != nil {
		log.Fatalln("failed to read file.")
	}
	// create a pipeline to process transactions in batches
	pipe := Client.Pipeline()
	pipe.Expire("pipeline_counter", time.Second)

	startCacheWrites := time.Now() //track for cache writes

	sliceLength := len(words)
	for i := 0; i < sliceLength; i++ {
		key, _ := GenerateKey(words[i])
		pipe.SAdd(key, words[i])
	}
	pipe.Exec()

	t := time.Now()
	elapsed := t.Sub(startCacheWrites)
	message := fmt.Sprintf("INGEST FINISHED: loaded and parsed file into memory and completed %d writes to Redis in %s", sliceLength, elapsed.String())
	log.Println(message)
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// SortString sorts a given string by split and join method
func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
