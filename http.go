package SnakeMasters

import (
	"fmt"
	"log"
	"net/http"
)

func (w *World) server(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Welcom to the Snake Masters!")

}

func (w *World) StartServer() {
	http.HandleFunc("/", w.server)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("startServer:", err)
	}
}
