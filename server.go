package SnakeMasters

import (
	"fmt"
	"image/color"
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
	w.game(w.clMap[name].num, conn)
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
		_, err = fmt.Fprint(conn, "\n\rEnter you name: ")

		name = ""
		_, err = fmt.Fscanln(conn, &name)

		ans := w.correctName(name)
		_, err = fmt.Fprint(conn, "\n\r")
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

func (w *World) game(cl int, conn net.Conn) {
	for {
		for n := range w.clSnake[cl] {
			as := w.areaString(&w.clSnake[cl][n])
			_, err := fmt.Fprintf(conn, "\n\rSnake position: %d, %d\n\n\r",
				w.clSnake[cl][n].body[0].x, w.clSnake[cl][n].body[0].y)
			_, err = fmt.Fprint(conn, as+"\n\n\r")

			for {
				_, err = fmt.Fprint(conn, "Your move: ")
				move := ""
				_, err = fmt.Fscanln(conn, &move)
				m := w.setMove(move, &w.clSnake[cl][n])

				if m == "" {
					w.move(&w.clSnake[cl][n])
					break
				} else {
					_, err = fmt.Fprint(conn, m+"\n\n\r")
				}

				if err != nil {
					err = conn.Close()
					errProc(err)
					return
				}
			}

			if err != nil {
				err = conn.Close()
				errProc(err)
				return
			}
		}
	}
}
