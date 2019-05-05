package SnakeMasters

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"image/color"
	"io"
	"log"
	"math/rand"
	"net"
	"regexp"
	"sync"
	"time"
)

var (
	mutex sync.Mutex
)

type JsonOutput struct {
	Answer  string    `json:"answer"`
	Session string    `json:"session,omitempty"`
	Data    *JsonData `json:"data,omitempty"`
}

type JsonInput struct {
	Session string
	Moves   []Moves
}

type Moves struct {
	NumSnake  int
	Direction string
}

type JsonData struct {
	Area   *[][]int
	Snakes *[]snake
}

/*
func (w *World) loginName(conn net.Conn) string {
	var userAns *JsonInput

	for {
		sentAnswer("Enter your name.", nil, conn)

		userAns = jsonAnsUnmarshal(conn)

		ans := w.correctName(userAns.Name)
		sentAnswer(ans, nil, conn)

		if ans[0:6] == "Hellow" {
			return userAns.Name
		} else {
			return "E"
		}
	}
}
*/
func (w *World) correctName(name string) (ans, session string) {
	if len(name) < 3 || len(name) > 16 {
		return "Name must consist 3-16 char.", ""
	}
	if mathc, err := regexp.MatchString("^[a-zA-Z0-9]*$", name); !mathc {
		return "Name must consist a-z, A-Z, 0-9 char.", ""
	} else {
		errProc(err)
	}
	if w.userNum[name] > 0 {
		return "Name is busy.", ""
	}

	session = string(rand.Int()) + time.Now().String()
	h := md5.New()

	_, err := io.WriteString(h, session)
	if err != nil {
		panic(err)
	}
	session = fmt.Sprintf("%x", h.Sum(nil))

	mutex.Lock()
	w.addNewUser(name, session)
	mutex.Unlock()

	return "Hellow, " + name + "!", session
}

func (w *World) addNewUser(name, session string) {
	var u User

	R := uint8(32 + rand.Intn(192))
	G := uint8(32 + rand.Intn(192))
	B := uint8(32 + rand.Intn(192))
	u.Color = color.RGBA{R, G, B, 255}

	userNum := len(w.users)
	w.userNum[name] = userNum
	w.userSession[name] = session
	w.users = append(w.users, u)

	w.users[userNum].addNewSnake(w)
}

func (w *World) deleteUser(name string) {
	log.Println("Delete user: ", name)

	for n := range w.users[w.userNum[name]].Snakes {
		w.users[w.userNum[name]].Snakes[n].die(w, &w.users[w.userNum[name]])
	}

	w.users[w.userNum[name]].disconnect = true

	mutex.Lock()
	delete(w.userNum, name)
	mutex.Unlock()
}

func (w *World) game(u *User, conn net.Conn) {
	gOld := 0
	for {

		for {
			if gOld < w.Gen {
				gOld = w.Gen
				break
			}
			time.Sleep(1 * time.Millisecond)
		}

		var data JsonData
		data.Area = &w.area
		data.Snakes = &u.Snakes

		sentAnswer("Sent data.", &data, conn)

		userAns := jsonAnsUnmarshal(conn)

		for n := range userAns.Moves {
			s := w.setMove(userAns.Moves[n].Direction, &u.Snakes[userAns.Moves[n].NumSnake])
			if s != "" {
				sentAnswer(s, nil, conn)
				return
			}
		}

	}
}

func jsonAnsUnmarshal(conn net.Conn) *JsonInput {
	var userAns *JsonInput = &JsonInput{}
	var jb []byte
	b := make([]byte, 64)

	for {
		n, err := conn.Read(b)
		jb = append(jb, b[:n]...)
		if err != nil {
			log.Println(err)
		}
		if n < 64 {
			break
		}
	}

	err := json.Unmarshal(jb, userAns)
	if err != nil {
		panic(err)
	}

	return userAns
}
