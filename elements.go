package SnakesMaster

import (
	"log"
	"math/rand"
)

const (
	elEat  = -1
	elWall = -2
)

func (w *World) addEat() {
	x := rand.Intn(w.lenX)
	y := rand.Intn(w.lenY)
	sy := y
	n := 0

	for {
		if w.area[x%w.lenX][y%w.lenY] == 0 {
			w.area[x%w.lenX][y%w.lenY] = elEat
			break
		}
		x++
		y = sy + x/w.lenX
		n++
		if n > w.lenX*w.lenY {
			log.Fatal("addEat: No place to add eat.")
		}
	}
}

func (w *World) addEatN(n int) {
	for x := 0; x < n; x++ {
		w.addEat()
	}
}

func (w *World) delEat() {
	x := rand.Intn(w.lenX)
	y := rand.Intn(w.lenY)
	sy := y
	n := 0

	for {
		if w.area[x%w.lenX][y%w.lenY] == elEat {
			w.area[x%w.lenX][y%w.lenY] = 0
			break
		}
		x++
		y = sy + x/w.lenX
		n++
		if n > w.lenX*w.lenY {
			log.Fatal("delEat: No eat to delete.")
		}
	}
}

func (w *World) delEatN(n int) {
	for x := 0; x < n; x++ {
		w.delEat()
	}
}

func (w *World) addWall() {
	x := rand.Intn(w.lenX)
	y := rand.Intn(w.lenY)
	sy := y
	n := 0

	for {
		if w.area[x%w.lenX][y%w.lenY] == 0 {
			w.area[x%w.lenX][y%w.lenY] = elWall
			break
		}
		x++
		y = sy + x/w.lenX
		n++
		if n > w.lenX*w.lenY {
			log.Fatal("addWall: No place to add wall.")
		}
	}
}

func (w *World) addWallN(n int) {
	for x := 0; x < n; x++ {
		w.addWall()
	}
}

func (w *World) setBorderWall() {
	for x := range w.area {
		w.area[x][0] = elWall
		w.area[x][w.lenY-1] = elWall
	}
	for y := range w.area[0] {
		w.area[0][y] = elWall
		w.area[w.lenX-1][y] = elWall
	}
}
