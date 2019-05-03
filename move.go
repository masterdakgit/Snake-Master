package SnakeMasters

import "fmt"

func (w *World) setMove(move string, s *snake) string {
	switch move {
	case "l":
		s.dir.dx = -1
		s.dir.dy = 0
		return ""
	case "r":
		s.dir.dx = 1
		s.dir.dy = 0
		return ""
	case "u":
		s.dir.dx = 0
		s.dir.dy = -1
		return ""
	case "d":
		s.dir.dx = 0
		s.dir.dy = 1
		return ""
	}
	return `You must enter: "l", "r", "u" or "d".`
}

func (w *World) move(s *snake, cl int) {
	x := s.body[0].x
	y := s.body[0].y

	x = x + s.dir.dx
	y = y + s.dir.dy

	switch w.area[x][y] {
	case elWall:
		return
	case elHead:
		return
	case elBody:
		if len(s.body) > 1 && s.body[1].x == x && s.body[1].y == y {
			s.div(w, cl)
		}
		return
	case elEat:
		s.eat(w)
	case elEmpty:
	}

	lastN := len(s.body) - 1

	for n := range s.body {
		w.area[s.body[n].x][s.body[n].y] = elEmpty
	}

	for n := lastN; n > 0; n-- {
		s.body[n].x = s.body[n-1].x
		s.body[n].y = s.body[n-1].y
	}

	s.body[0].x = x
	s.body[0].y = y

	for n := range s.body {
		w.area[s.body[n].x][s.body[n].y] = elBody
	}

	w.area[x][y] = elHead
}

func (s *snake) eat(w *World) {
	var c cell
	nLast := len(s.body) - 1

	c.x = s.body[nLast].x
	c.y = s.body[nLast].y
	c.color = s.body[nLast].color

	s.body = append(s.body, c)
}

func (s *snake) div(w *World, cl int) {
	var newSnake snake
	sLen := len(s.body)
	newSnake.body = make([]cell, sLen-sLen/2)

	for n := sLen / 2; n < sLen; n++ {
		newSnake.body[n-sLen/2].x = s.body[n].x
		newSnake.body[n-sLen/2].y = s.body[n].y
		newSnake.body[n-sLen/2].color = s.body[n].color
	}

	s.body = s.body[:sLen/2]
	fmt.Println(sLen, len(s.body), len(newSnake.body))
	w.clSnake[cl] = append(w.clSnake[cl], newSnake)
}
