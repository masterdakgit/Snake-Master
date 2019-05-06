package SnakeMasters

import (
	"crypto/md5"
	"fmt"
	"image/color"
	"io"
	"log"
	"math/rand"
	"regexp"
	"sync"
	"time"
)

var (
	mutex, mutexAddSession, mmutexSleeper sync.Mutex
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
	w.users[w.userNum[name]].disconnect = false
	w.antiSleep[name] = 0

	w.users[userNum].addNewSnake(w)
}

func (w *World) deleteUser(name, ip string) {
	log.Println("Delete user: ", name)

	for n := range w.users[w.userNum[name]].Snakes {
		w.users[w.userNum[name]].Snakes[n].die(w, &w.users[w.userNum[name]])
	}

	w.users[w.userNum[name]].disconnect = true

	mutex.Lock()
	delete(w.userNum, name)
	delete(w.userSession, name)
	delete(human, name)
	ipMap[ip]--
	mutex.Unlock()
}

func errProc(err error) {
	if err != nil {
		log.Println(err)
	}
}
