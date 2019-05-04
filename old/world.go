package old

import (
	"image/color"
)

const (
	energeStart = 10
)

type World struct {
	clMap      map[string]client
	clSnake    [][]snake
	area       [][]int
	balance    int
	lenX, lenY int
	Gen        int
}

type SnakeSlice struct {
	Color  color.RGBA
	Snakes []snake
	Area   [][]int
}

type client struct {
	num   int
	name  string
	color color.RGBA
}

type snake struct {
	Body   []cell
	Energe int
	dir    direction
	dead   bool
}

type direction struct {
	dx, dy int
}

type cell struct {
	X, Y int
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
			if w.area[x][y] != ElWall && w.area[x][y] != ElEmpty {
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
		w.area[x][0] = ElWall
		w.area[x][w.lenY-1] = ElWall
	}
	for y := range w.area[0] {
		w.area[0][y] = ElWall
		w.area[w.lenX-1][y] = ElWall
	}
}

func (w *World) Generation() {
	for _, cl := range w.clMap {
		for sn, s := range w.clSnake[cl.num] {

			if s.dead {
				continue
			}

			w.move(&s, cl.num)
			s.Energe--
			if s.Energe <= 0 {
				s.eatSomeSelf(w, cl.num, sn)
			}
		}
	}
	w.setBalance()
}
