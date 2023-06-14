package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type coords struct {
	x int
	y int
}

// rectangle is the smallest unit of UI. All other elements are made up of rectangles.
type rectangle struct {
	img *ebiten.Image
	x   int
	y   int
}

// square is 5 rectangles: 4 for the border, with another one for the middle
type square [5]*rectangle

// shape is a collection of squares representing a single tetris piece
type shape []*square

func MakeRectangle(size coords, clr color.Color, position coords) *rectangle {
	img := ebiten.NewImage(size.x, size.y)
	img.Fill(clr)
	return &rectangle{img, position.x, position.y}
}

func MakeSquare(clr color.Color, position coords) *square {
	border := color.RGBA{171, 166, 166, 0xff}
	arr := square{}
	arr[0] = MakeRectangle(coords{46, 2}, border, coords{position.x, position.y})
	arr[1] = MakeRectangle(coords{46, 2}, border, coords{position.x + 2, position.y + 46})
	arr[2] = MakeRectangle(coords{2, 46}, border, coords{position.x + 46, position.y})
	arr[3] = MakeRectangle(coords{2, 46}, border, coords{position.x, position.y + 2})
	arr[4] = MakeRectangle(coords{44, 44}, clr, coords{position.x + 2, position.y + 2})
	return &arr
}
