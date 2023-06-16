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
	shp = append(shp, &square{position: coords{x: position.x, y: position.y - 3}, color: clr})
	shp = append(shp, &square{position: coords{x: position.x, y: position.y - 2}, color: clr})
	shp = append(shp, &square{position: coords{x: position.x, y: position.y - 1}, color: clr})
	shp = append(shp, &square{position: coords{x: position.x, y: position.y}, color: clr})
	return &shp
}

// j block
func makeJ(position coords) *shape {
	shp := shape{}
	clr := color.RGBA{0, 0, 255, 0xff}
	shp = append(shp, &square{position: coords{x: position.x, y: position.y - 2}, color: clr})
	shp = append(shp, &square{position: coords{x: position.x, y: position.y - 1}, color: clr})
	shp = append(shp, &square{position: coords{x: position.x - 1, y: position.y}, color: clr})
	shp = append(shp, &square{position: coords{x: position.x, y: position.y}, color: clr})
	return &shp
}

// l block
func makeL(position coords) *shape {
	shp := shape{}
	clr := color.RGBA{255, 165, 0, 0xff}
	shp = append(shp, &square{position: coords{x: position.x, y: position.y - 2}, color: clr})
	shp = append(shp, &square{position: coords{x: position.x, y: position.y - 1}, color: clr})
	shp = append(shp, &square{position: coords{x: position.x + 1, y: position.y}, color: clr})
	shp = append(shp, &square{position: coords{x: position.x, y: position.y}, color: clr})
	return &shp
}

// o block
func makeO(position coords) *shape {
	shp := shape{}
	clr := color.RGBA{255, 255, 0, 0xff}
	shp = append(shp, &square{position: coords{x: position.x, y: position.y - 1}, color: clr})
	shp = append(shp, &square{position: coords{x: position.x + 1, y: position.y - 1}, color: clr})
	shp = append(shp, &square{position: coords{x: position.x + 1, y: position.y}, color: clr})
	shp = append(shp, &square{position: coords{x: position.x, y: position.y}, color: clr})
	return &shp
}

// s block
func makeS(position coords) *shape {
	shp := shape{}
	clr := color.RGBA{0, 128, 0, 0xff}
	shp = append(shp, &square{position: coords{x: position.x, y: position.y - 1}, color: clr})
	shp = append(shp, &square{position: coords{x: position.x + 1, y: position.y - 1}, color: clr})
	shp = append(shp, &square{position: coords{x: position.x - 1, y: position.y}, color: clr})
	shp = append(shp, &square{position: coords{x: position.x, y: position.y}, color: clr})
	return &shp
}

// t block
func makeT(position coords) *shape {
	shp := shape{}
	clr := color.RGBA{128, 0, 128, 0xff}
	shp = append(shp, &square{position: coords{x: position.x, y: position.y - 1}, color: clr})
	shp = append(shp, &square{position: coords{x: position.x - 1, y: position.y}, color: clr})
	shp = append(shp, &square{position: coords{x: position.x + 1, y: position.y}, color: clr})
	shp = append(shp, &square{position: coords{x: position.x, y: position.y}, color: clr})
	return &shp
}

// z block
func makeZ(position coords) *shape {
	shp := shape{}
	clr := color.RGBA{255, 0, 0, 0xff}
	shp = append(shp, &square{position: coords{x: position.x, y: position.y - 1}, color: clr})
	shp = append(shp, &square{position: coords{x: position.x - 1, y: position.y - 1}, color: clr})
	shp = append(shp, &square{position: coords{x: position.x + 1, y: position.y}, color: clr})
	shp = append(shp, &square{position: position, color: clr})
	return &shp
}

var shapeFuncs = [7](func(coords) *shape){makeI, makeJ, makeL, makeO, makeS, makeT, makeZ}
