package SnakeMasters

import "image/color"

const (
	startLength = 4
)

type snake struct {
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
