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

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func init() {
	var anagrams Anagrams
	startReadFile := time.Now()
	file := "dictionary.txt"
	words, _ := readLines(file)
	t1 := time.Now()
	fmt.Println(t1.Sub(startReadFile))
	sliceLength := len(words)
	client := RedisClient()
	startWriteCache := time.Now()
	for i := 0; i < sliceLength; i++ {
		word := words[i]
		key := sortString(word)
		client.LPush(key, &anagrams)
	}

	t := time.Now()
	elapsed := t.Sub(startWriteCache)
	fmt.Println(elapsed)
	// ch := make(chan string, )
}
