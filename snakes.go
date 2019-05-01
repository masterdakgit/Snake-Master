package SnakesMaster

import "image/color"

const (
	startLength = 4
)

type snakes struct {
	num     int
	command int
	body    []cell
	parent  []int
	childs  []int
}

type cell struct {
	x, y  int
	color color.RGBA
}
