package SnakeMasters

import "image/color"

const (
	energeStart = 10
)

type World struct {
	clMap      map[string]client
	clSnake    [][]snake
	area       [][]int
	balance    int
	lenX, lenY int
}

type client struct {
	num   int
	name  string
	color color.RGBA
}

type snake struct {
	body   []cell
	dir    direction
	energe int
}

type direction struct {
	dx, dy int
}

type cell struct {
	x, y  int
	color color.RGBA
}

func (w *World) Create(x, y, balance, wall int) {
	w.clMap = make(map[string]client)
	w.clSnake = make([][]snake, 0)

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
}

func (w *World) currentBalance() int {
	result := 0
	for x := range w.area {
		for y := range w.area[x] {
			if w.area[x][y] != elWall && w.area[x][y] != elEmpty {
				result++
			}
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
		w.area[x][0] = elWall
		w.area[x][w.lenY-1] = elWall
	}
	for y := range w.area[0] {
		w.area[0][y] = elWall
		w.area[w.lenX-1][y] = elWall
	}
}
