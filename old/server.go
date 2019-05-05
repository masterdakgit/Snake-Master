package old

import (
	"encoding/json"
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
		//go w.handleConnection(conn)
	}
}

func errProc(err error) {
	if err != nil {
		log.Println(err)
	}
}

/*
func (w *World) handleConnection(conn net.Conn) {
	if len(w.userNum) > 20 {
		sentAnswer("Too many connections, log in later.", nil, conn)
		return
	}
	//name := w.loginName(conn)

	if name == "E" {
		sentAnswer("Error to enter name.", nil, conn)
		return
	}

	w.game(&w.users[w.userNum[name]], conn)
	sentAnswer("Connection or answer error.", nil, conn)
	w.deleteUser(name)
}
*/
func sentAnswer(ans string, data *JsonData, conn net.Conn) bool {
	var a JsonOutput
	a.Answer = ans
	a.Data = data

	var b []byte
	b, err := json.Marshal(a)
	if err != nil {
		log.Println(err)
	}

	_, err = conn.Write(b)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
