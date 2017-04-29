package main

import (
	"github.com/semihs/goform"
	"net/http"
	"text/template"
)

func main() {
	email := goform.NewEmailElement("email", "Email", []*goform.Attribute{}, []goform.ValidatorInterface{
		&goform.RequiredValidator{},
	})
	password := goform.NewPasswordElement("password", "Password", []*goform.Attribute{}, []goform.ValidatorInterface{})
	submit := goform.NewButtonElement("submit", "Login", []*goform.Attribute{})

	form := goform.NewGoForm()
	form.Add(email)
	form.Add(password)
	form.Add(submit)

	tpl, _ := template.New("tpl").Parse(view)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		if r.Method != "POST" {
			tpl.Execute(w, form)
			return
		}
		r.ParseForm()
		form.BindFromRequest(r)

		if !form.IsValid() {
			tpl.Execute(w, form)
			return
		}
	})
	http.ListenAndServe(":2626", nil)

}
