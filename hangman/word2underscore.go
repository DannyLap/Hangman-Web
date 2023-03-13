package hangman

func Word2underscore(Word_string []string, String1 string) []string {
	for index := 0; index < len(String1); index++ { // creation d'un tableau de rune de longueur équivalente au mot
		Word_string = append(Word_string, string('_')) // composé de tiret du bas
	}
	return Word_string
}
