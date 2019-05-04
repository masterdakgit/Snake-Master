package old

import (
	"encoding/json"
	"fmt"
	"image/color"
	"log"
	"net"
	"time"
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

func (w *World) game(cl client, conn net.Conn) {
	gOld := 0
	for {

		for {
			if gOld < w.Gen {
				gOld = w.Gen
				break
			}
			time.Sleep(1 * time.Millisecond)
		}

		var ssio SnakeSlice
		ssio.Color = cl.color
		ssio.Snakes = w.clSnake[cl.num]
		ssio.Area = w.area

		b, err := json.Marshal(ssio)
		if err != nil {
			log.Println(err)
			break
		}

		conn.Write(b)
		_, err = fmt.Fprint(conn, "\r\n")
		if err != nil {
			log.Println(err)
			break
		}

		move := ""
		for n := range w.clSnake[cl.num] {
		reEnter:
			fmt.Fscanln(conn, &move)
			if n >= len(w.clSnake[cl.num]) {
				break
			}
			s := w.setMove(move, &w.clSnake[cl.num][n])
			if s != "" {
				fmt.Fprintln(conn, s)
				goto reEnter
			}
		}
	}
}
