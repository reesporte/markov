package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

/* char level markov chain text generator */

// split into char level tokens
func tokenize(words string) []string {
	r := make([]string, 0)

	for _, char := range words {
		if string(char) != "\n" {
			r = append(r, string(char))
		}
	}
	return r
}

// generate ngrams from tokens
func ngrams(tokens []string, ngram_level int) map[string][]string {
	n := make(map[string][]string)
	char := ""
	for i := 0; i < len(tokens)-ngram_level; i++ {
		char = tokens[i]
		if _, ok := n[char]; !ok {
			n[char] = make([]string, 0)
		}
		for j := 1; j < ngram_level; j++ {
			n[char] = append(n[char], strings.Join(tokens[i+1:i+1+j], ""))
		}
	}

	return n
}

// get a random choice from array of tokens
func choice(tokens []string) string {
	if len(tokens) == 0 {
		return " "
	}
	choice := tokens[rand.Intn(len(tokens))]
	return choice
}

// generate quote from list of tokens and ngrams
func quote(tokens []string, ngrams map[string][]string) string {
	last := choice(tokens)
	quote := strings.Title(last)

	count := rand.Intn(15) + 30
	for i := 0; i < count; i++ {
		last = choice(ngrams[string(last[len(last)-1])])
		quote += last
	}
	return quote
}

// read words and generate quote
func main() {
	rand.Seed(time.Now().Unix())
	input, err := ioutil.ReadFile("words.txt")
	if err != nil {
		panic(err)
	}

	tokens := tokenize(string(input))
	ngrams := ngrams(tokens, 40)

	fmt.Println(quote(tokens, ngrams))
}
