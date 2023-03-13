package hangman

func Slice2str(String1 string, String2 string, Word_string []string) string {
	for i := 0; i < len(String1); i++ { // on parcourt le mot
		String2 += Word_string[i] // on ajoute les lettres du slice une par une à une nouvelle variable string
		String2 += " "            // à l'ajout de chaque lettre, on ajoute un espace
	}
	return String2
}
