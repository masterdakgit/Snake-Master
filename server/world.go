package server

import (
	"image"
	"image/color"
	"time"
)

const (
	antiDDoS     = 50 * time.Millisecond
	antiSleepSec = 60
	maxSession   = 20
	maxUserToIp  = 20
)

var (
	ipMap map[string]int
)

type World struct {
	userNum     map[string]int
	userSession map[string]string
	antiSleep   map[string]int
	users       []User
	area        [][]int
	lenX, lenY  int
	balance     int
	Imgage      *image.RGBA
	Gen         int
}

type User struct {
	Color      color.RGBA
	Snakes     []snake
	ip         string
	disconnect bool
	antiDDoS   bool
}

func (w *World) Create(x, y, balance, wall int) {
	w.Imgage = image.NewRGBA(image.Rect(0, 0, x*bar+1+infoPanelX, y*bar+1))

	w.users = make([]User, 1)
	w.userNum = make(map[string]int)
	w.userSession = make(map[string]string)
	w.antiSleep = make(map[string]int)
	ipMap = make(map[string]int)

	w.lenX = x
	w.lenY = y

	w.area = make([][]int, x)
	for n := range w.area {
		w.area[n] = make([]int, y)
	}

	w.setWallEdge()
	w.addWallN(wall)
	w.balance = balance
	w.setBalance()
	w.imgChange()
}

func (w *World) currentBalance() int {
	result := 0
	for x := range w.area {
		for y := range w.area[x] {
			if w.area[x][y] == ElEat {
				result++
			}
		}
	}
	for u := range w.users {
		for s := range w.users[u].Snakes {
			if w.users[u].Snakes[s].Dead {
				continue
			}
			result += len(w.users[u].Snakes[s].Body)
		}
	}
	return result
}

func (w *World) setBalance() {
	currentBalance := w.currentBalance()
	if currentBalance < w.balance {
		w.addEatN(w.balance - currentBalance)
	}
	if currentBalance > w.balance {
		w.delEatN(currentBalance - w.balance)
	}
}

func (w *World) setWallEdge() {
	for x := range w.area {
		w.area[x][0] = ElWall
		w.area[x][w.lenY-1] = ElWall
	}
	for y := range w.area[0] {
		w.area[0][y] = ElWall
		w.area[w.lenX-1][y] = ElWall
	}
}

func (w *World) Generation() {
	for u := range w.users {
		if w.users[u].disconnect {
			continue
		}
		for s := range w.users[u].Snakes {
			if w.users[u].Snakes[s].Dead {
				continue
			}
			w.users[u].Snakes[s].move(w, &w.users[u])
			w.users[u].Snakes[s].Energe--
			if w.users[u].Snakes[s].Energe <= 0 {
				w.users[u].Snakes[s].eatSomeSelf(w, &w.users[u])
			}
		}
	}
	w.setBalance()
	mutex.Lock()
	w.imgChange()
	w.deleteDead()
	mutex.Unlock()
}

func (w *World) deleteDead() {
	for u := range w.users {
		for s := 0; s < len(w.users[u].Snakes); s++ {
			if w.users[u].Snakes[s].Dead {
				if s == len(w.users[u].Snakes)-1 {
					w.users[u].Snakes = w.users[u].Snakes[:s]
				} else {
					w.users[u].Snakes = append(w.users[u].Snakes[:s], w.users[u].Snakes[s+1:]...)
				}
			}
		}
	}
}
