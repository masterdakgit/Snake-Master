package SnakeMasters

type World struct {
	area       [][]int
	lenX, lenY int
	balance    int
	snakes     []snake
	clients    []client
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
	w.snakes = make([]snake, 0)
	w.clients = make([]client, 0)
	w.setBalance()
	w.addWallN(wall)
}
