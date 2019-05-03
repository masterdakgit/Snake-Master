package SnakeMasters

import (
	"encoding/json"
	"fmt"
	"image/color"
	"log"
	"net"
)

const (
	startLength = 4
)

var (
	colorHead = color.RGBA{0, 0, 0, 255}
)

func (w *World) ListenAndServe(port string) {
	listener, err := net.Listen("tcp", port)

	errProc(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			err = conn.Close()
			errProc(err)
			continue
		}
		go w.handleConnection(conn)
	}
}

func (w *World) handleConnection(conn net.Conn) {
	name := w.loginName(conn)
	w.game(w.clMap[name], conn)
	delete(w.clMap, name)
}

func errProc(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func (w *World) loginName(conn net.Conn) string {
	_, err := fmt.Fprint(conn, "Welcome to the Snake Masters!\n\r")
	var name string

	for {
		_, err = fmt.Fprint(conn, "Enter you name:\n\r")

		name = ""
		_, err = fmt.Fscanln(conn, &name)

		ans := w.correctName(name)
		_, err = fmt.Fprint(conn, ans+"\n\r")

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

func (w *World) game(cl client, conn net.Conn) {
	for {
		var ssio SnakeSlice
		ssio.Color = cl.color
		ssio.Snakes = w.clSnake[cl.num]

		for n := range ssio.Snakes {
			w.visiString(&ssio.Snakes[n])
		}

		b, err := json.Marshal(ssio)
		if err != nil {
			log.Println(err)
			break
		}

		conn.Write(b)
		_, err = fmt.Fprint(conn, "\n\r")
		if err != nil {
			log.Println(err)
			break
		}

		move := ""
		for n := range w.clSnake[cl.num] {
		reEnter:
			fmt.Fscanln(conn, &move)
			s := w.setMove(move, &w.clSnake[cl.num][n])
			if s != "" {
				fmt.Fprintln(conn, s)
				goto reEnter
			} else {
				w.move(&w.clSnake[cl.num][n], cl.num)
			}
		}
	}
}
