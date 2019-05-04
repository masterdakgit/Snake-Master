package SnakeMasters

import (
	"image/png"
	"log"
	"net/http"
)

func loadHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func (w *World) loadPict(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		mutex.Lock()
		err := png.Encode(rw, w.Imgage)
		if err != nil {
			log.Println("loadPict:", err)
		}
		mutex.Unlock()
	}
}

func (w *World) ListenHTTP(port string) {
	http.HandleFunc("/pict/", w.loadPict)
	http.HandleFunc("/", loadHTML)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Println("ListenHTTP:", err)
	}
}
