package SnakeMasters

import (
	"fmt"
	"log"
	"net"
)

func (w *World) ListenAndServe(port string) {
	listener, err := net.Listen("tcp", port)
	errProc(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			err = conn.Close()
			errProc(err)
			continue
		}
		go w.handleConnection(conn)
	}
}

func errProc(err error) {
	if err != nil {
		log.Println(err)
	}
}

func (w *World) handleConnection(conn net.Conn) {
	if len(w.userNum) > 10 {
		_, err := fmt.Fprintln(conn, "Snake Masters: Too many connections, log in later.")
		if err != nil {
			log.Println(err)
		}
		conn.Close()
		return
	}
	name := w.loginName(conn)
	w.game(&w.users[w.userNum[name]], conn)
	w.deleteUser(name)
}
