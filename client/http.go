package human

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var (
	host  string
	mutex sync.Mutex
)

type jsonOutput struct {
	Answer  string    `json:"answer"`
	Session string    `json:"session,omitempty"`
	Data    *jsonData `json:"data,omitempty"`
}

type jsonData struct {
	Area   *[][]int
	Snakes *[]snake
}

type snake struct {
	Num    int
	Body   []cell
	Energe int
	Dead   bool
}

type cell struct {
	X, Y int
}

func Create(site string) {
	host = site
	human = make(map[string]humanData)
}

func HumanGame(rw http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	user := r.FormValue("user")
	session := r.FormValue("session")

	if user != "" {
		if human[user].name != user || human[user].session != session || session == "" {
			var newHuman humanData
			newHuman.name = user
			human[user] = newHuman
			humanConnection(user, rw)
			return
		}
	}

	http.ServeFile(rw, r, "human.html")
}

var human map[string]humanData

type humanData struct {
	name    string
	session string
	resp    http.Response
}

func humanConnection(user string, rw http.ResponseWriter) {
	resp, err := http.Get(host + "/game/?user=" + user)
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}

	var data *jsonOutput = &jsonOutput{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(data)
	if err != nil {
		panic(err)
	}

	if data.Session == "" {
		fmt.Fprintln(rw, data.Answer)
		return
	}

	newHuman := human[user]
	newHuman.session = data.Session
	human[user] = newHuman

	go humanBots(user, data.Session)

	fmt.Fprintf(rw, "<a href=\"/human/?user=%s&session=%s\">Start game</a>",
		user, human[user].session)

}

func humanBots(user, session string) {
	var data *jsonOutput = &jsonOutput{}
	var r http.Response

	for {
		time.Sleep(100 * time.Millisecond)
		if r.Body != nil {
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(data)
			if err != nil && err != io.EOF {
				panic(err)
			}
			m := "_"

			if data.Data != nil && data.Data.Snakes != nil {
				snakes := *data.Data.Snakes

				for n := range snakes {
					if n == 0 {
						continue
					}
					r := rand.Intn(4)
					rs := "lrud"[r : r+1]
					if len(snakes[n].Body) > 8 {
						rs = "/"
					}
					m += rs
				}
			}

			r.Body.Close()
			resp, err := http.Get(host + "/game/?user=" + user + "&session=" + session + "&move=" + m +
				"&humanbot=true")
			if err != nil {
				panic(err)
			}

			r = *resp

			continue
		}

		resp, err := http.Get(host + "/game/?user=" + user + "&session=" + session + "&humanbot=true")
		if err != nil {
			panic(err)
		}

		r = *resp
	}
}

func Key(rw http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		return
	}

	user := r.FormValue("user")
	session := r.FormValue("session")
	key := r.FormValue("key")

	m := ""
	switch key {
	case "ArrowUp":
		m = "u"
	case "ArrowDown":
		m = "d"
	case "ArrowLeft":
		m = "l"
	case "ArrowRight":
		m = "r"
	case " ":
		m = "/"

	}

	var data *jsonOutput = &jsonOutput{}

	mutex.Lock()
	defer mutex.Unlock()

	if human[user].resp.Body != nil {
		decoder := json.NewDecoder(human[user].resp.Body)
		err := decoder.Decode(data)
		if err != nil {
			return
			//panic(err)
		}

		if data.Data != nil {
			snakes := *data.Data.Snakes
			for n := range snakes {
				if n == 0 {
					continue
				}
				r := rand.Intn(4)
				rs := "lrud"[r : r+1]
				if len(snakes[n].Body) > 8 {
					rs = "/"
				}
				m += rs
			}
		}
	}

	if human[user].resp.Body != nil {
		human[user].resp.Body.Close()
	}
	resp, err := http.Get(host + "/game/?user=" + user + "&session=" + session + "&move=" + m)
	if err != nil {
		panic(err)
	}

	newHuman := human[user]
	newHuman.resp = *resp
	human[user] = newHuman

}
