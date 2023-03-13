package hangman

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

func ChooseWord() string { //choice of the word taken from Words.txt
	var words string
	var wordsfromtxt []string
	rand.Seed(time.Now().UTC().UnixNano())
	word, err := os.Open(string(os.Args[1]))

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(word)
	nbwords := 0

	for scanner.Scan() {
		nbwords++
		wordsfromtxt = append(wordsfromtxt, scanner.Text())
	}
	indexword := rand.Intn(nbwords)
	words = wordsfromtxt[indexword]
	return words
}
