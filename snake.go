package SnakeMasters

const (
	startLength = 4
	energeStart = 10
)

type snake struct {
	Num    int
	Body   []cell
	Energe int
	Dead   bool
	dir    direction
}

type direction struct {
	dx, dy int
}

type cell struct {
	X, Y int
}

func (u *User) addNewSnake(w *World) {
	s := u.snakeCreate(w)
	s.Num = len(u.Snakes)
	u.Snakes = append(u.Snakes, s)
}

func (u *User) snakeCreate(w *World) snake {
	var s snake
	x, y := w.findElement(ElEmpty)
	s.Body = make([]cell, startLength)
	s.Energe = energeStart

	for n := range s.Body {
		s.Body[n].X = x
		s.Body[n].Y = y
	}
	w.area[s.Body[0].X][s.Body[0].Y] = ElHead
	return s
}
