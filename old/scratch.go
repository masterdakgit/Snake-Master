package main

import (
	"SnakeMasters"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

var (
	w SnakeMasters.World
	sd SnakeMasters.SnakeSlice
)

func main(){
	w.Create(50, 15, 100, 30)
	go w.ListenAndServe(":5301")
	go func() {
		for {
			w.Generation()
			w.Gen++
			time.Sleep(200 * time.Millisecond)
		}
	}()

	conn, err := net.Dial("tcp", "localhost:5301")
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	name := "masterdak"

	fmt.Fprintln(conn, name)

	input := bufio.NewScanner(conn)

	for {
		if input.Scan(){
			fmt.Println(input.Text())
			if input.Text()[0:1] == "{" {
				err = json.Unmarshal(input.Bytes(), &sd)
				if err != nil {
					log.Println(err)
				}
				for y := range sd.Area[0]{
					for x := range sd.Area{
						switch sd.Area[x][y] {
						case SnakeMasters.ElBody:
							fmt.Print("o")
						case SnakeMasters.ElEat:
							fmt.Print("*")
						case SnakeMasters.ElEmpty:
							fmt.Print(".")
						case SnakeMasters.ElHead:
							fmt.Print("@")
						case SnakeMasters.ElWall:
							fmt.Print("#")
						}
					}
					fmt.Println()
				}
			}
			for range sd.Snakes{
				r := rand.Intn(4)
				fmt.Fprintln(conn, "rlud"[r:r+1])
			}
		}
		time.Sleep(10 * time.Millisecond)
	}
}