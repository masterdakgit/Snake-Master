package SnakesMaster

var dir [4]direction

type direction struct {
	dx, dy int
}

func setDir() {
	dir[0].dx = -1
	dir[0].dy = 0

	dir[1].dx = 1
	dir[1].dy = 0

	dir[2].dx = 0
	dir[2].dy = -1

	dir[3].dx = 0
	dir[3].dy = 1
}
