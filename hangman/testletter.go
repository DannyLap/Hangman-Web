package hangman

func Testletter(Word_string []string, String1 string, String2 string, Letter []string, Life int, Life_d int) (string, int, int) {
	var wrongLetter bool = true
	if len(Letter) != 0 {
		for i := 0; i < len(String1); i++ { // on parcourt le mot
			if string(String1[i]) == string(Letter[0]) {
				Word_string[i] = string(Letter[0])
				wrongLetter = false
			}
		}
		if wrongLetter {
			Life--
			Life_d++
		}
		String2 = ""
		for i := 0; i < len(String1); i++ { // on parcourt le mot
			String2 += Word_string[i] // on ajoute les lettres du slice une par une à une nouvelle variable string
			String2 += " "            // à l'ajout de chaque lettre, on ajoute un espace
		}
	}
	return String2, Life, Life_d
}
