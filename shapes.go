package main

import (
	"image/color"
)

type coords struct {
	x int
	y int
}

type iShape struct {
	shapeImpl
}

func (s *iShape) Rotate() {
	squares := s.squares()
	if s.isHorizontal() {
		squares[0].position.x += 2
		squares[0].position.y -= 1
		squares[1].position.x += 1
		squares[2].position.y += 1
		squares[3].position.x -= 1
		squares[3].position.y += 2
	} else {
		squares[0].position.x -= 2
		squares[0].position.y += 1
		squares[1].position.x -= 1
		squares[2].position.y -= 1
		squares[3].position.x += 1
		squares[3].position.y -= 2
	}
}

func (s *iShape) isHorizontal() bool {
	return s.squares()[0].position.y == s.squares()[1].position.y
}

// i block
func makeI(position coords) shape {
	shp := iShape{}
	clr := color.RGBA{0, 255, 255, 0xff}
	shp.shapeImpl = append(shp.shapeImpl, &square{position: coords{x: position.x, y: position.y - 3}, color: clr})
	shp.shapeImpl = append(shp.shapeImpl, &square{position: coords{x: position.x, y: position.y - 2}, color: clr})
	shp.shapeImpl = append(shp.shapeImpl, &square{position: coords{x: position.x, y: position.y - 1}, color: clr})
	shp.shapeImpl = append(shp.shapeImpl, &square{position: coords{x: position.x, y: position.y}, color: clr})
	return &shp
}

// j block
func makeJ(position coords) shape {
	clr := color.RGBA{0, 0, 255, 0xff}
	return &shapeImpl{
		&square{position: coords{x: position.x, y: position.y - 1}, color: clr}, // middle square
		&square{position: coords{x: position.x, y: position.y - 2}, color: clr}, // top square
		&square{position: coords{x: position.x - 1, y: position.y}, color: clr}, // branch square
		&square{position: coords{x: position.x, y: position.y}, color: clr},     // bottom square
	}
}

// l block
func makeL(position coords) shape {
	shp := shapeImpl{}
	clr := color.RGBA{255, 165, 0, 0xff}
	shp = append(shp, &square{position: coords{x: position.x, y: position.y - 1}, color: clr}) // middle square
	shp = append(shp, &square{position: coords{x: position.x, y: position.y - 2}, color: clr})
	shp = append(shp, &square{position: coords{x: position.x + 1, y: position.y}, color: clr})
	shp = append(shp, &square{position: coords{x: position.x, y: position.y}, color: clr})
	return &shp
}

type oShape struct {
	shapeImpl
}

func (s *oShape) Rotate() {}

// o block
func makeO(position coords) shape {
	shp := oShape{}
	clr := color.RGBA{255, 255, 0, 0xff}
	shp.shapeImpl = append(shp.shapeImpl, &square{position: coords{x: position.x, y: position.y - 1}, color: clr})
	shp.shapeImpl = append(shp.shapeImpl, &square{position: coords{x: position.x + 1, y: position.y - 1}, color: clr})
	shp.shapeImpl = append(shp.shapeImpl, &square{position: coords{x: position.x + 1, y: position.y}, color: clr})
	shp.shapeImpl = append(shp.shapeImpl, &square{position: coords{x: position.x, y: position.y}, color: clr})
	return &shp
}

// s block
func makeS(position coords) shape {
	shp := shapeImpl{}
	clr := color.RGBA{0, 128, 0, 0xff}
	shp = append(shp, &square{position: coords{x: position.x, y: position.y}, color: clr}) // middle square
	shp = append(shp, &square{position: coords{x: position.x, y: position.y - 1}, color: clr})
	shp = append(shp, &square{position: coords{x: position.x + 1, y: position.y - 1}, color: clr})
	shp = append(shp, &square{position: coords{x: position.x - 1, y: position.y}, color: clr})
	return &shp
}

// t block
func makeT(position coords) shape {
	shp := shapeImpl{}
	clr := color.RGBA{128, 0, 128, 0xff}
	shp = append(shp, &square{position: coords{x: position.x, y: position.y}, color: clr}) // middle square
	shp = append(shp, &square{position: coords{x: position.x, y: position.y - 1}, color: clr})
	shp = append(shp, &square{position: coords{x: position.x - 1, y: position.y}, color: clr})
	shp = append(shp, &square{position: coords{x: position.x + 1, y: position.y}, color: clr})
	return &shp
}

// z block
func makeZ(position coords) shape {
	shp := shapeImpl{}
	clr := color.RGBA{255, 0, 0, 0xff}
	shp = append(shp, &square{position: position, color: clr}) // middle square
	shp = append(shp, &square{position: coords{x: position.x, y: position.y - 1}, color: clr})
	shp = append(shp, &square{position: coords{x: position.x - 1, y: position.y - 1}, color: clr})
	shp = append(shp, &square{position: coords{x: position.x + 1, y: position.y}, color: clr})
	return &shp
}

var shapeFuncs = [7](func(coords) shape){makeI, makeJ, makeL, makeO, makeS, makeT, makeZ}
