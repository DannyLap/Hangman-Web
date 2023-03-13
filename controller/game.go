package controller

import (
	"HangmanWeb/hangman"
	"net/http"
	"text/template"
)

type Game struct {
	Title          string
	Body           string
	Template       *template.Template
	Life           int
	Life_d         int
	Word_string    []string
	String1        string
	String2        string
	Letter         []string
	Win            string
	Loose          string
	Hangmandrawing []string
}

func (g *Game) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	g.Letter = r.Form["letterValue"]
	g.String2, g.Life, g.Life_d = hangman.Testletter(g.Word_string, g.String1, g.String2, g.Letter, g.Life, g.Life_d)
	g.Hangmandrawing = hangman.Hangmandrawing(g.Life_d)
	Win := hangman.Winloose(g.Word_string, g.Life)
	if Win {
		g.Win = "  !!! WIN !!!"
	} else if g.Life < 1 {
		g.Win = "  !!! LOST !!!"
	}
	if hangman.Loose(g.Life) {
		g.Loose = "!!! LOOSE !!!"
	}
	err := g.Template.Execute(w, g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
