package anagram

// Anagrams is a list of anagrams for a given word
type Anagrams struct {
	Anagrams []string `json:"anagrams"`
}

// RequestBody is a list words to be added to the corpus
type RequestBody struct {
	Words []string `json:"words"`
}
