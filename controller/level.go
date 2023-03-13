package controller

import (
	"fmt"
	"net/http"
	"text/template"
)

type Level struct {
	Title    string
	Body     string
	Template *template.Template
}

func (l Level) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form["test"])
	err := l.Template.Execute(w, l)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
