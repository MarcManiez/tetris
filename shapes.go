package main

import (
	"image/color"
)

type coords struct {
	x int
	y int
}

type shapes [7]*shape

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
	shp = append(shp, MakeSquare(clr, coords{position.x, position.y - 48}))
	shp = append(shp, MakeSquare(clr, coords{position.x + 48, position.y - 48}))
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
