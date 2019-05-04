package old

import (
	"log"
	"math/rand"
)

const (
	viewRange = 6
	viewLen   = 1 + 2*viewRange

	ElEmpty = 0
	ElEat   = -1
	ElWall  = -2
	ElHead  = 2
	ElBody  = 1
)

func (w *World) findElement(el int) (x, y int) {
	x = rand.Intn(w.lenX)
	y = rand.Intn(w.lenY)
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
			panic("findElement: The element no found.")
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

func (w *World) VisiString(s *snake) {
	x := s.Body[0].X
	y := s.Body[0].Y
	x0 := x - viewRange
	x1 := x + viewRange
	y0 := y - viewRange
	y1 := y + viewRange

	n := 0
	var as [viewLen][viewLen]string

	for y := y0; y <= y1; y++ {
		for x := x0; x <= x1; x++ {
			if y < 0 || y >= w.lenY || x < 0 || x >= w.lenX {
				//Out of the edge
				as[x-x0][y-y0] = "#"
				n++
			} else {
				as[x-x0][y-y0] = elString(w.area[x][y])
				n++
			}
		}
	}
	//s.Visibility = as
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
