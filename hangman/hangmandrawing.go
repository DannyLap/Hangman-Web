package hangman

import (
	"fmt"
	"io/ioutil"
)

func Hangmandrawing(Life_d int) []string {
	var hangman []string

	data, err := ioutil.ReadFile("hangman/hangman.txt") // lire le fichier txt
	if err != nil {
		fmt.Println(err)
	}
	for index := 0 + (71 * Life_d); index <= 70+(71*Life_d); index++ {
		//on parcourt une partie du document .txt correspondant au nombre de tentatives restantes
		hangman = append(hangman, string(data[index])) //que l'on imprime
	}
	return hangman
}
