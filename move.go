package SnakeMasters

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

func (w *World) move(s *snake) {
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
		return
	case elEat:
		return
	case elEmpty:
	}

	lastN := len(s.body) - 1
	w.area[s.body[lastN].x][s.body[lastN].y] = elEmpty

	for n := lastN; n > 0; n-- {
		s.body[n].x = s.body[n-1].x
		s.body[n].y = s.body[n-1].y
	}

	w.area[s.body[lastN].x][s.body[lastN].y] = elBody
	w.area[s.body[0].x][s.body[0].y] = elBody
	w.area[x][y] = elHead
	s.body[0].x = x
	s.body[0].y = y

}
