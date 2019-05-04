package SnakeMasters

import (
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"sort"
	"strconv"
)

const (
	infoPanelX = 160
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
		for _, s := range u.Snakes {
			if s.Dead {
				continue
			}
			for _, b := range s.Body {
				setBar(b.X, b.Y, u.Color, w.Imgage)
			}
			setBar(s.Body[0].X, s.Body[0].Y, colorHead, w.Imgage)
		}
	}
}

func (w *World) imgChange() *image.RGBA {
	w.setInfoPanel()

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

func (w *World) setInfoPanel() {
	for x := w.lenX*bar + 1; x < w.lenX*bar+infoPanelX; x++ {
		for y := 0; y < w.lenY*bar; y++ {
			w.Imgage.Set(x, y, colorHead)
		}
	}

	user := make([]userScore, len(w.userNum))

	k := 0
	for n := range w.userNum {
		user[k].user = n
		user[k].score = w.users[w.userNum[n]].score()
		k++
	}

	sort.Slice(user, func(i, j int) bool {
		return user[i].score > user[j].score
	})

	k = 0
	for n := range user {
		y := 20 + k*20
		x := w.lenX*bar + 10
		addLabel(w.Imgage, x, y, user[n].user, w.users[w.userNum[user[n].user]].Color)
		y = 20 + k*20
		x = w.lenX*bar + 125
		addLabel(w.Imgage, x, y, strconv.Itoa(user[n].score), colorEmpty)
		k++
	}
}

func addLabel(img *image.RGBA, x, y int, label string, col color.RGBA) {
	//col := color.RGBA{0, 0, 0, 255}
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}

type userScore struct {
	user  string
	score int
}

func (u *User) score() int {
	result := 0
	for n := range u.Snakes {
		if u.Snakes[n].Dead {
			continue
		}
		result += len(u.Snakes[n].Body)
	}
	return result
}
