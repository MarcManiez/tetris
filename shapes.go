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

func (s *shape) move() {
	for _, sqr := range *s {
		for _, rect := range *sqr {
			rect.y += 48
		}
	}
}

type shapes [7]*shape

func makeRectangle(size coords, clr color.Color, position coords) *rectangle {
	img := ebiten.NewImage(size.x, size.y)
	img.Fill(clr)
	return &rectangle{img, position.x, position.y}
}

func MakeSquare(clr color.Color, position coords) *square {
	arr := square{}
	arr[0] = makeRectangle(coords{46, 2}, color.RGBA{184, 184, 184, 0xff}, coords{position.x, position.y})
	arr[1] = makeRectangle(coords{46, 2}, color.RGBA{136, 136, 136, 0xff}, coords{position.x + 2, position.y + 46})
	arr[2] = makeRectangle(coords{2, 46}, color.RGBA{200, 200, 200, 0xff}, coords{position.x + 46, position.y})
	arr[3] = makeRectangle(coords{2, 46}, color.RGBA{150, 150, 150, 0xff}, coords{position.x, position.y + 2})
	arr[4] = makeRectangle(coords{44, 44}, clr, coords{position.x + 2, position.y + 2})
	return &arr
}

// i block
func makeI(position coords) *shape {
	shp := shape{}
	clr := color.RGBA{0, 255, 255, 0xff}
	shp = append(shp, MakeSquare(clr, coords{position.x, position.y - 3*48}))
	shp = append(shp, MakeSquare(clr, coords{position.x, position.y - 2*48}))
	shp = append(shp, MakeSquare(clr, coords{position.x, position.y - 48}))
	shp = append(shp, MakeSquare(clr, position))
	return &shp
}

// j block
func makeJ(position coords) *shape {
	shp := shape{}
	clr := color.RGBA{0, 0, 255, 0xff}
	shp = append(shp, MakeSquare(clr, coords{position.x, position.y - 2*48}))
	shp = append(shp, MakeSquare(clr, coords{position.x, position.y - 48}))
	shp = append(shp, MakeSquare(clr, coords{position.x - 48, position.y}))
	shp = append(shp, MakeSquare(clr, position))
	return &shp
}

// l block
func makeL(position coords) *shape {
	shp := shape{}
	clr := color.RGBA{255, 165, 0, 0xff}
	shp = append(shp, MakeSquare(clr, coords{position.x, position.y - 2*48}))
	shp = append(shp, MakeSquare(clr, coords{position.x, position.y - 48}))
	shp = append(shp, MakeSquare(clr, coords{position.x + 48, position.y}))
	shp = append(shp, MakeSquare(clr, position))
	return &shp
}

// o block
func makeO(position coords) *shape {
	shp := shape{}
	clr := color.RGBA{255, 255, 0, 0xff}
	shp = append(shp, MakeSquare(clr, coords{position.x, position.y + 48}))
	shp = append(shp, MakeSquare(clr, coords{position.x + 48, position.y + 48}))
	shp = append(shp, MakeSquare(clr, coords{position.x + 48, position.y}))
	shp = append(shp, MakeSquare(clr, position))
	return &shp
}

// s block
func makeS(position coords) *shape {
	shp := shape{}
	clr := color.RGBA{0, 128, 0, 0xff}
	shp = append(shp, MakeSquare(clr, coords{position.x, position.y - 48}))
	shp = append(shp, MakeSquare(clr, coords{position.x + 48, position.y - 48}))
	shp = append(shp, MakeSquare(clr, coords{position.x - 48, position.y}))
	shp = append(shp, MakeSquare(clr, position))
	return &shp
}

// t block
func makeT(position coords) *shape {
	shp := shape{}
	clr := color.RGBA{128, 0, 128, 0xff}
	shp = append(shp, MakeSquare(clr, coords{position.x, position.y - 48}))
	shp = append(shp, MakeSquare(clr, coords{position.x - 48, position.y}))
	shp = append(shp, MakeSquare(clr, coords{position.x + 48, position.y}))
	shp = append(shp, MakeSquare(clr, position))
	return &shp
}

// z block
func makeZ(position coords) *shape {
	shp := shape{}
	clr := color.RGBA{255, 0, 0, 0xff}
	shp = append(shp, MakeSquare(clr, coords{position.x, position.y - 48}))
	shp = append(shp, MakeSquare(clr, coords{position.x - 48, position.y - 48}))
	shp = append(shp, MakeSquare(clr, coords{position.x + 48, position.y}))
	shp = append(shp, MakeSquare(clr, position))
	return &shp
}

var shapeFuncs = [7](func(coords) *shape){makeI, makeJ, makeL, makeO, makeS, makeT, makeZ}
