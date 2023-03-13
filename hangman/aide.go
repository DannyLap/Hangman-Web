package hangman

import (
	"math/rand"
	"time"
)

func Aide(string1 string, Word_string []string) []string {
	letters2show := len(Word_string)/2 - 1 //on définit le nombre de lettre à afficher avant que le joueur ne commence à jouer
	rand.Seed(time.Now().UnixNano())
	for index := 0; index <= letters2show-1; index++ {
		c := string(string1[rand.Intn(len(string1))])             //on obtient une lettre au hasard du mot
		for index2 := 0; index2 <= len(Word_string)-1; index2++ { //on parcourt les index du mot
			if string1[index2] == c[0] { //si une lettre du mot mystère correspond à la lettre obtenu au hasard
				Word_string[index2] = string(c[0]) //on supprime le tiret du bas que l'on remplace par la lettre obtenu au hasard
			}
		}
	}
	return Word_string
}
