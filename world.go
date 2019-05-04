package SnakeMasters

import (
	"image"
	"image/color"
)

type World struct {
	userNum    map[string]int
	users      []User
	area       [][]int
	lenX, lenY int
	balance    int
	Imgage     *image.RGBA
	Gen        int
}

type User struct {
	Color      color.RGBA
	Snakes     []snake
	disconnect bool
}

func (w *World) Create(x, y, balance, wall int) {
	w.Imgage = image.NewRGBA(image.Rect(0, 0, x*bar+1+infoPanelX, y*bar+1))

	w.users = make([]User, 1)
	w.userNum = make(map[string]int)

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
	mutex.Unlock()
}
