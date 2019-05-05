package SnakeMasters

import (
	"encoding/json"
	"fmt"
	"image/png"
	"log"
	"net"
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
	if r.Method != "GET" {
		_, err := fmt.Fprintln(rw, `{"answer":"Only GET method handled.."}`)
		if err != nil {
			log.Println(err)
		}
		return
	}

	name := r.FormValue("user")
	session := r.FormValue("session")
	move := r.FormValue("move")

	log.Println(name, "try to connect...")

	if name == "" {
		_, err := fmt.Fprintln(rw, `{"answer":"Request must contain user."}`)
		if err != nil {
			log.Println(err)
		}
		return
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Println(err)
	}

	if len(w.userSession) >= maxSession {
		_, err = fmt.Fprintln(rw, `{"answer":"Too many connection, log in later."}`)
		if err != nil {
			log.Println(err)
		}
		return
	}

	mutexAddSession.Lock()
	if ipMap[ip] >= maxUserToIp {
		_, err = fmt.Fprintln(rw, `{"answer":"Too many connection from one ip, log in later."}`)
		if err != nil {
			log.Println(err)
		}
		mutexAddSession.Unlock()
		return
	}
	mutexAddSession.Unlock()

	mutexAddSession.Lock()
	if session == "" {
		w.addNewSession(name, rw, ip)
		mutexAddSession.Unlock()
		return
	}
	mutexAddSession.Unlock()

	if w.userSession[name] == session {
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

		mmutexSleeper.Lock()
		w.antiSleep[name] = 0
		mmutexSleeper.Unlock()

		if !w.users[w.userNum[name]].antiDDoS {
			mutex.Lock()
			w.sentGameData(&w.users[w.userNum[name]], rw)
			w.users[w.userNum[name]].antiDDoS = true
			go w.users[w.userNum[name]].unsetDDoS()
			mutex.Unlock()
		} else {
			var output JsonOutput
			output.Answer = "Between request must be pause 10 ms."
			jsonSent(&output, rw)
		}
	}

}

func (u *User) unsetDDoS() {
	time.Sleep(antiDDoS)
	u.antiDDoS = false
}

func sleeper(name string, w *World, ip string) {
	for {
		time.Sleep(1 * time.Second)
		mmutexSleeper.Lock()
		w.antiSleep[name]++

		if w.antiSleep[name] > antiSleepSec {
			w.deleteUser(name, ip)
			mmutexSleeper.Unlock()
			return
		}
		mmutexSleeper.Unlock()
	}
}

func (w *World) addNewSession(name string, rw http.ResponseWriter, ip string) {
	var output JsonOutput
	ans, session := w.correctName(name)
	output.Answer = ans
	if session != "" {
		output.Session = session
		log.Println("Add new user: ", name, ip)
		mutex.Lock()
		ipMap[ip]++
		mutex.Unlock()
		w.users[w.userNum[name]].ip = ip
		go sleeper(name, w, ip)
	}
	jsonSent(&output, rw)
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
