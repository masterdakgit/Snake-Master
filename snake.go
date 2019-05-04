package SnakeMasters

const (
	startLength = 4
	energeStart = 10
)

type snake struct {
	Body   []cell
	Energe int
	dir    direction
	dead   bool
}

type direction struct {
	dx, dy int
}

type cell struct {
	X, Y int
}

func (u *User) addNewSnake(w *World) {
	s := u.snakeCreate(w)
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

	return s
}
