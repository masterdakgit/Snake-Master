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
	name := w.loginName(conn)

	_, err := fmt.Fscanln(conn)
	errProc(err)

	w.deleteUser(name)
}
