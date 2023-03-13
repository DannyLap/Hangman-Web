package hangman

func Winloose(Word_string []string, Life int) bool {
	Win := true
	for i := 0; i < len(Word_string); i++ { // creation d'un tableau de rune de longueur équivalente au mot
		if string(Word_string[i]) == string("_") {
			Win = false
		}
	}
	return Win
}

func Loose(Life int) bool {
	Loose := false
	if Life <= 0 {
		Loose = true
	}
	return Loose
}
