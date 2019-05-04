package SnakeMasters

import (
	"encoding/json"
	"fmt"
	"image/color"
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

type jsonSent struct {
	Area *[][]int
	User *User
}

func (w *World) loginName(conn net.Conn) string {
	_, err := fmt.Fprint(conn, "Welcome to the Snake Masters!\r\n")
	var name string

	for {
		_, err = fmt.Fprint(conn, "Enter you name:\r\n")

		name = ""
		_, err = fmt.Fscanln(conn, &name)

		ans := w.correctName(name)
		_, err = fmt.Fprint(conn, ans+"\r\n")

		if err != nil {
			err = conn.Close()
			errProc(err)
			return "E"
		}

		if ans[0:6] == "Hellow" {
			return name
		}
	}
}

func (w *World) correctName(name string) string {
	if len(name) < 3 || len(name) > 16 {
		return "Name must consist 3-16 char."
	}
	if mathc, err := regexp.MatchString("^[a-zA-Z0-9]*$", name); !mathc {
		return "Name must consist a-z, A-Z, 0-9 char."
	} else {
		errProc(err)
	}
	if w.userNum[name] > 0 {
		return "Name is busy."
	}

	mutex.Lock()
	w.addNewUser(name)
	mutex.Unlock()

	return "Hellow, " + name + "!"
}

func (w *World) addNewUser(name string) {
	var u User

	R := uint8(32 + rand.Intn(192))
	G := uint8(32 + rand.Intn(192))
	B := uint8(32 + rand.Intn(192))
	u.Color = color.RGBA{R, G, B, 255}

	userNum := len(w.users)
	w.userNum[name] = userNum
	w.users = append(w.users, u)

	w.users[userNum].addNewSnake(w)
}

func (w *World) deleteUser(name string) {
	log.Println("Delete user: ", name)

	for _, s := range w.users[w.userNum[name]].Snakes {
		s.die(w, &w.users[w.userNum[name]])
	}

	w.users[w.userNum[name]].disconnect = true
	delete(w.userNum, name)
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

		var jsonData jsonSent
		jsonData.Area = &w.area
		jsonData.User = u

		b, err := json.Marshal(jsonData)
		if err != nil {
			log.Println(err)
			return
		}

		_, err = conn.Write(b)
		errProc(err)

		_, err = fmt.Fprint(conn, "\r\n")
		if err != nil {
			log.Println(err)
			return
		}

		move := ""
		for n := range u.Snakes {
			if u.Snakes[n].Dead {
				continue
			}
		reEnter:
			_, err = fmt.Fscanln(conn, &move)

			if err != nil {
				return
			}
			s := w.setMove(move, &u.Snakes[n])
			if s != "" {
				fmt.Fprintln(conn, s)
				goto reEnter
			}
		}

	}
}
