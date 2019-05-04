package SnakeMasters

import (
	"fmt"
	"image/color"
	"math/rand"
	"net"
	"regexp"
)

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

	w.addNewUser(name)

	return "Hellow, " + name + "!"
}

func (w *World) addNewUser(name string) {
	var u user

	R := uint8(rand.Intn(255))
	G := uint8(rand.Intn(255))
	B := uint8(rand.Intn(255))
	u.color = color.RGBA{R, G, B, 255}

	userNum := len(w.users)
	w.userNum[name] = userNum
	w.users = append(w.users, u)

	w.users[userNum].addNewSnake(w)
}

func (w *World) deleteUser(name string) {
	w.users[w.userNum[name]].disconnect = true
	delete(w.userNum, name)
}
