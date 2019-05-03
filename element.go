package SnakeMasters

import (
	"log"
	"math/rand"
)

const (
	viewRange = 6
	viewLen   = 1 + 2*viewRange

	elEmpty = 0
	elEat   = -1
	elWall  = -2
	elHead  = 2
	elBody  = 1
)

func (w *World) findElement(el int) (x, y int) {
	x = rand.Intn(w.lenX)
	y = rand.Intn(w.lenY)
	ys := y
	n := 0

	for {
		if w.area[x%w.lenX][y%w.lenY] == el {
			return x % w.lenX, y % w.lenY
		}
		x++
		y = ys + x%w.lenX
		n++
		if n > w.lenX*w.lenY {
			log.Fatal("findElement: The element no found.")
		}
	}
}

func (w *World) addWall() {
	x, y := w.findElement(elEmpty)
	w.area[x][y] = elWall
}

func (w *World) addWallN(n int) {
	for x := 0; x < n; x++ {
		w.addWall()
	}
}

func (w *World) addEat() {
	x, y := w.findElement(elEmpty)
	w.area[x][y] = elEat
}

func (w *World) addEatN(n int) {
	for x := 0; x < n; x++ {
		w.addEat()
	}
}

func (w *World) delEat() {
	x, y := w.findElement(elEat)
	w.area[x][y] = elEmpty
}

func (w *World) delEatN(n int) {
	for x := 0; x < n; x++ {
		w.addEat()
	}
}

func (w *World) visiString(s *snake) {
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
	s.Visibility = as
}

func elString(n int) string {
	switch n {
	case elEmpty:
		return "."
	case elEat:
		return "*"
	case elWall:
		return "#"
	case elHead:
		return "@"
	case elBody:
		return "o"
	}
	return "E"
}
