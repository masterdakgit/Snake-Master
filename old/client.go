package old

import (
	"image/color"
	"math/rand"
	"regexp"
)

func (w *World) correctName(name string) string {
	if len(name) < 3 || len(name) > 16 {
		return "Name must consist 3-16 char."
	}
	if mathc, err := regexp.MatchString("^[a-zA-Z0-9]*$", name); !mathc {
		return "Name must consist a-z, A-Z, 0-9 char."
	} else {
		errProc(err)
	}
	if w.clMap[name].name == name {
		return "Name is busy."
	}

	w.addNewClient(name)

	return "Hellow, " + name + "!"
}

func (w *World) addNewClient(name string) {
	var c client
	c.name = name
	c.num = len(w.clSnake)

	R := uint8(rand.Intn(255))
	G := uint8(rand.Intn(255))
	B := uint8(rand.Intn(255))
	c.color = color.RGBA{R, G, B, 255}

	w.clSnake = append(w.clSnake, []snake{})
	w.clSnake[c.num] = make([]snake, 0)

	w.clMap[name] = c

	w.addNewSnake(name)
}

func (w *World) addNewSnake(name string) {
	num := w.clMap[name].num
	s := w.snakeCreate(name)
	w.area[s.Body[0].X][s.Body[0].Y] = ElHead
	w.clSnake[num] = append(w.clSnake[num], s)
}

func (w *World) snakeCreate(name string) snake {
	var s snake
	x, y := w.findElement(ElEmpty)
	s.Body = make([]cell, startLength)
	s.Energe = energeStart

	for n := range s.Body {
		s.Body[n].X = x
		s.Body[n].Y = y
	}

	return s
}
