package SnakeMasters

import (
	"image"
	"image/color"
)

const (
	infoPanelY = 64
)

var (
	bar        int = 8
	colorEmpty     = color.RGBA{255, 255, 255, 255}
	colorGreed     = color.RGBA{220, 220, 220, 255}
	colorHead      = color.RGBA{0, 0, 0, 255}
	colorWall      = color.RGBA{170, 170, 170, 255}
	colorEat       = color.RGBA{0, 170, 0, 255}
)

func setBar(x, y int, c color.RGBA, i *image.RGBA) {
	for bx := 0; bx < bar+1; bx++ {
		for by := 0; by < bar+1; by++ {
			i.Set(x*bar+bx, y*bar+by, c)
			if bx%bar == 0 || by%bar == 0 {
				i.Set(x*bar+bx, y*bar+by, colorGreed)
			}
		}
	}
}

func (w *World) setSnakeImg() {
	for _, u := range w.users {
		if u.disconnect {
			continue
		}
		for _, s := range u.snakes {
			for _, b := range s.Body {
				setBar(b.X, b.Y, u.color, w.Imgage)
			}
			setBar(s.Body[0].X, s.Body[0].Y, colorHead, w.Imgage)
		}
	}
}

func (w *World) imgChange() *image.RGBA {
	for x := range w.area {
		for y := range w.area[x] {
			switch w.area[x][y] {
			case ElWall:
				setBar(x, y, colorWall, w.Imgage)
			case ElEmpty:
				setBar(x, y, colorEmpty, w.Imgage)
			case ElEat:
				setBar(x, y, colorEat, w.Imgage)
			}
		}
	}

	w.setSnakeImg()
	return w.Imgage
}
