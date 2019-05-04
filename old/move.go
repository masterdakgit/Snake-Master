package old

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
	x := s.Body[0].X
	y := s.Body[0].Y

	x = x + s.dir.dx
	y = y + s.dir.dy

	switch w.area[x][y] {
	case ElWall:
		return
	case ElHead:
		return
	case ElBody:
		if len(s.Body) > 1 && s.Body[1].X == x && s.Body[1].Y == y {
			s.div(w, cl)
		}
		return
	case ElEat:
		s.eat(w)
	case ElEmpty:
	}

	lastN := len(s.Body) - 1

	for n := range s.Body {
		w.area[s.Body[n].X][s.Body[n].Y] = ElEmpty
	}

	for n := lastN; n > 0; n-- {
		s.Body[n].X = s.Body[n-1].X
		s.Body[n].Y = s.Body[n-1].Y
	}

	s.Body[0].X = x
	s.Body[0].Y = y

	for n := range s.Body {
		w.area[s.Body[n].X][s.Body[n].Y] = ElBody
	}

	w.area[x][y] = ElHead
}

func (s *snake) eat(w *World) {
	var c cell
	nLast := len(s.Body) - 1

	c.X = s.Body[nLast].X
	c.Y = s.Body[nLast].Y

	s.Body = append(s.Body, c)
}

func (s *snake) div(w *World, cl int) {
	sLen := len(s.Body)

	var newSnake snake
	newSnake.Body = make([]cell, sLen-sLen/2)

	for n := sLen / 2; n < sLen; n++ {
		newSnake.Body[n-sLen/2].X = s.Body[n].X
		newSnake.Body[n-sLen/2].Y = s.Body[n].Y
	}

	s.Body = s.Body[:sLen/2]
	w.clSnake[cl] = append(w.clSnake[cl], newSnake)
}

func (s *snake) eatSomeSelf(w *World, cl, sn int) {
	nLast := len(s.Body) - 1

	if nLast < 1 {
		s.die(w, cl, sn)
		return
	}

	w.area[s.Body[nLast].X][s.Body[nLast].Y] = 0
	s.Body = s.Body[:nLast]

	s.Energe = energeStart
}

func (s *snake) die(w *World, cl, sn int) {
	for n := range s.Body {
		w.area[s.Body[n].X][s.Body[n].Y] = 1
	}

	s.dead = true

}

func remove(s []snake, i int) []snake {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
