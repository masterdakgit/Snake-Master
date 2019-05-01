package SnakeMasters

type World struct {
	area       [][]int
	lenX, lenY int
	balance    int
	snake      []snakes
}

func (w *World) Сreate(lenX, lenY, balance, wall int) {
	setDir()
	w.area = make([][]int, lenX)
	for n := range w.area {
		w.area[n] = make([]int, lenY)
	}
	w.lenX = lenX
	w.lenY = lenY
	w.balance = balance
	w.snake = make([]snakes, 0)
	w.setBalance()
	w.addWallN(wall)
}
