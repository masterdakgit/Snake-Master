package SnakeMasters

import (
	"crypto/md5"
	"fmt"
	"hash"
	"io"
	"log"
	"time"
)

type client struct {
	num         int
	hash        hash.Hash
	write, read string
}

func (w *World) addClient() {
	var c client
	c.num = len(w.clients)
	c.hash = md5.New()

	_, err := io.WriteString(c.hash, string(c.num)+time.Now().String())
	if err != nil {
		log.Fatal("addClient:", err)
	}

	fmt.Printf("%x", c.hash.Sum(nil))

	w.clients = append(w.clients, c)
}
