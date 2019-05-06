package server

import (
	"log"
	"math/rand"
)

const (
	ElEmpty = 0
	ElEat   = -1
	ElWall  = -2
	ElHead  = 2
	ElBody  = 1
)

func (w *World) findElement(el int) (x, y int) {
	x = 1 + rand.Intn(w.lenX-2)
	y = 1 + rand.Intn(w.lenY-2)
	n := 0

	for {
		if w.area[x%w.lenX][y%w.lenY] == el {
			return x % w.lenX, y % w.lenY
		}
		x++
		if x%w.lenX == 0 {
			y++
		}
		n++
		if n > w.lenX*w.lenY {
			log.Println(x, y, n)
			log.Println("findElement: The element no found.")
		}
	}
}

func (w *World) addWall() {
	x, y := w.findElement(ElEmpty)
	w.area[x][y] = ElWall
}

func (w *World) addWallN(n int) {
	for x := 0; x < n; x++ {
		w.addWall()
	}
}

func (w *World) addEat() {
	x, y := w.findElement(ElEmpty)
	w.area[x][y] = ElEat
}

func (w *World) addEatN(n int) {
	for x := 0; x < n; x++ {
		w.addEat()
	}
}

func (w *World) delEat() {
	x, y := w.findElement(ElEat)
	w.area[x][y] = ElEmpty
}

func (w *World) delEatN(n int) {
	for x := 0; x < n; x++ {
		w.delEat()
	}
}

func elString(n int) string {
	switch n {
	case ElEmpty:
		return "."
	case ElEat:
		return "*"
	case ElWall:
		return "#"
	case ElHead:
		return "@"
	case ElBody:
		return "o"
	}
	return "E"
}
