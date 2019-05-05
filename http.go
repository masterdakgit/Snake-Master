package SnakeMasters

import (
	"encoding/json"
	"image/png"
	"log"
	"net/http"
	"time"
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
	http.HandleFunc("/game/", w.gameHTTP)
	http.HandleFunc("/", loadHTML)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Println("ListenHTTP:", err)
	}
}

func (w *World) gameHTTP(rw http.ResponseWriter, r *http.Request) {
	name := r.FormValue("user")
	session := r.FormValue("session")
	move := r.FormValue("move")
	if name != "" {
		if session != "" && w.userSession[name] == session {
			if move != "" {
				for n := range move {
					s := w.setMove(move[n:n+1], &w.users[w.userNum[name]].Snakes[n])
					if s != "" {
						var output JsonOutput
						output.Answer = s
						jsonSent(&output, rw)
					}
				}
			}

			w.users[w.userNum[name]].antiSleepSec = 0

			if !w.users[w.userNum[name]].antiDDoS {
				w.sentGameData(&w.users[w.userNum[name]], rw)
				w.users[w.userNum[name]].antiDDoS = true
				go w.users[w.userNum[name]].unsetDDoS()
			} else {
				var output JsonOutput
				output.Answer = "Between request must be pause 10 ms."
				jsonSent(&output, rw)
			}
		} else {
			w.addNewSession(name, rw)
		}
	}
}

func (u *User) unsetDDoS() {
	time.Sleep(antiDDoS)
	u.antiDDoS = false
}

func (w *World) sleeper(name string) {
	for {
		time.Sleep(1 * time.Second)
		w.users[w.userNum[name]].antiSleepSec++

		if w.users[w.userNum[name]].antiSleepSec > antiSleepSec {
			w.deleteUser(name)
			return
		}
	}
}

func (w *World) addNewSession(name string, rw http.ResponseWriter) {
	var output JsonOutput
	ans, session := w.correctName(name)
	output.Answer = ans
	if session != "" {
		output.Session = session
	}
	jsonSent(&output, rw)
	go w.sleeper(name)
}

func jsonSent(output *JsonOutput, rw http.ResponseWriter) {
	encoder := json.NewEncoder(rw)
	err := encoder.Encode(output)
	if err != nil {
		panic(err)
	}
}

func (w *World) sentGameData(u *User, rw http.ResponseWriter) {
	var output JsonOutput
	var data JsonData
	data.Area = &w.area
	data.Snakes = &u.Snakes
	output.Answer = "Sent game data."
	output.Data = &data
	jsonSent(&output, rw)
}
