package anagram

import (
	"bufio"
	"fmt"
	"time"
	// "fmt"
	"os"
	"sort"
	"strings"
	// "time"
)

type sortRunes []rune

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

func SortString(w string) string {
	word := strings.ToLower(w)
	s := strings.Split(word, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// TODO need to add deterministic hashing to remove dups

func init() {
	Client = RedisClient()
	// read in file
	file := "dictionary.txt"
	words, _ := readLines(file)
	sliceLength := len(words)
	// create a pipeline to batch process transactions
	pipe := Client.Pipeline()
	pipe.Expire("pipeline_counter", time.Second)
	startWriteCache := time.Now()
	for i := 0; i < sliceLength; i++ {
		word := words[i]
		key := SortString(word)
		pipe.LPush(key, word)
	}
	pipe.Exec()

	t := time.Now()
	elapsed := t.Sub(startWriteCache)
	fmt.Println(elapsed)
}
