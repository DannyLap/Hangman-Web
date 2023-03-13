package main

import (
	"HangmanWeb/controller"
	"HangmanWeb/hangman"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var menuTemplate = template.Must(template.ParseFiles("templates/menu.html"))
var gameTemplate = template.Must(template.ParseFiles("templates/game.html"))
var levelTemplate = template.Must(template.ParseFiles("templates/level.html"))

var gamePage = new(controller.Game)

func main() {
	Life := 10
	Life_d := 0
	var Word_string []string
	var String1 string
	var String2 string
	var Letter []string
	var Win string
	var Loose string
	var Hangmandrawing []string
	fmt.Println("http://localhost:8080/menu")
	menuPage := controller.Menu{"Menu", "Bonjour ! Que voulez-vous faire ?", menuTemplate}
	levelPage := controller.Level{"Level", "Quel Level choisir", levelTemplate}
	*gamePage = controller.Game{"Game", "Hangman Web", gameTemplate, Life, Life_d, Word_string, String1, String2, Letter, Win, Loose, Hangmandrawing}
	gamePage.String1 = hangman.ChooseWord()
	gamePage.Word_string = hangman.Word2underscore(gamePage.Word_string, gamePage.String1)
	gamePage.Word_string = hangman.Aide(gamePage.String1, gamePage.Word_string)
	gamePage.String2 = hangman.Slice2str(gamePage.String1, gamePage.String2, gamePage.Word_string)

	http.Handle("/menu", menuPage)
	http.Handle("/game", gamePage)
	http.Handle("/level", levelPage)

	http.Handle("/", http.FileServer(http.Dir("/static/menu.css")))

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
