package server

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
	case "/":
		s.dir.dx = 1000
		s.dir.dy = 0
		return ""
	case "_":
		return ""
	}
	return `You must enter: "l", "r", "u", "d", "/" or "_".`
}

func (s *snake) move(w *World, u *User) {
	if s.dir.dx == 1000 {
		s.dir.dx = 0
		s.div(w, u)
		return
	}

	x := s.Body[0].X + s.dir.dx
	y := s.Body[0].Y + s.dir.dy

	switch w.area[x][y] {
	case ElWall:
		return
	case ElHead:
		return
	case ElBody:
		return
	case ElEat:
		s.eat()
	}

	for n := range s.Body {
		w.area[s.Body[n].X][s.Body[n].Y] = ElEmpty
	}

	for n := len(s.Body) - 1; n > 0; n-- {
		s.Body[n] = s.Body[n-1]
	}

	s.Body[0].X = x
	s.Body[0].Y = y

	for n := range s.Body {
		w.area[s.Body[n].X][s.Body[n].Y] = ElBody
	}

	w.area[x][y] = ElHead
}

func (s *snake) eat() {
	var c cell
	c.X = s.Body[len(s.Body)-1].X
	c.Y = s.Body[len(s.Body)-1].Y

	s.Body = append(s.Body, c)
}

func (s *snake) eatSomeSelf(w *World, u *User) {
	nLast := len(s.Body) - 1

	if nLast < 1 {
		s.die(w, u)
		return
	}

	w.area[s.Body[nLast].X][s.Body[nLast].Y] = ElEmpty
	s.Body = s.Body[:nLast]

	s.Energe = energeStart
}

func (s *snake) div(w *World, u *User) {
	sLen := len(s.Body)

	if sLen < 2 {
		return
	}

	var newSnake snake
	newSnake.Body = make([]cell, sLen-sLen/2)

	for n := sLen / 2; n < sLen; n++ {
		newSnake.Body[n-sLen/2].X = s.Body[n].X
		newSnake.Body[n-sLen/2].Y = s.Body[n].Y
	}

	s.Body = s.Body[:sLen/2]
	u.Snakes = append(u.Snakes, newSnake)
}

func (s *snake) die(w *World, u *User) {
	for n := range s.Body {
		w.area[s.Body[n].X][s.Body[n].Y] = ElEat
	}
	s.Dead = true

	if u.liveSnakes() == 0 && !u.disconnect {
		s.Dead = false
		s.Energe = energeStart
	}
}

func (u *User) liveSnakes() int {
	l := 0
	for n := range u.Snakes {
		if !u.Snakes[n].Dead {
			l++
		}
	}
	return l
}
