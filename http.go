package SnakeMasters

import (
	"html/template"
	"log"
	"net/http"
)

func (w *World) server(wr http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

	}
	tmpl, err := template.ParseFiles("index.html")
	login := `<a href="` + string(w.clients[n].hash.Sum(nil)) + `">Start game</a>`
	err = tmpl.Execute(wr, login)

	if err != nil {
		log.Fatal("w.server:", err)
	}
}

func (w *World) StartServer() {
	http.HandleFunc("/", w.server)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("startServer:", err)
	}
}
