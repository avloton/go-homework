package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("./web/templates/index.html")
		if err != nil {
			log.Println(err)
			return
		}
		tmpl.Execute(w, nil)
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("./web/templates/about.html")
		if err != nil {
			log.Println(err)
			return
		}
		tmpl.Execute(w, nil)
	}
}

func ContactsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("./web/templates/contacts.html")
		if err != nil {
			log.Println(err)
			return
		}
		tmpl.Execute(w, nil)
	}
}

func MenuHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("./web/templates/menu.html")
		if err != nil {
			log.Println(err)
			return
		}
		tmpl.Execute(w, nil)
	}
}

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("./web/templates/order.html")
		if err != nil {
			log.Println(err)
			return
		}
		tmpl.Execute(w, nil)
	}
}