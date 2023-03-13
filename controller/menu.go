package controller

import (
	"net/http"
	"text/template"
)

type Menu struct {
	Title    string
	Body     string
	Template *template.Template
}

func (m Menu) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	err := m.Template.Execute(w, m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
