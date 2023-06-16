package main

import "image/color"

// square is 5 rectangles: 4 for the border, with another one for the middle
type square [5]*rectangle

func MakeSquare(clr color.Color, position coords) *square {
	arr := square{}
	// Top
	arr[0] = makeRectangle(coords{46, 2}, color.RGBA{184, 184, 184, 0xff}, coords{position.x, position.y})
	// Bottom
	arr[1] = makeRectangle(coords{46, 2}, color.RGBA{136, 136, 136, 0xff}, coords{position.x + 2, position.y + 46})
	// Right
	arr[2] = makeRectangle(coords{2, 46}, color.RGBA{200, 200, 200, 0xff}, coords{position.x + 46, position.y})
	// Left
	arr[3] = makeRectangle(coords{2, 46}, color.RGBA{150, 150, 150, 0xff}, coords{position.x, position.y + 2})
	// Middle
	arr[4] = makeRectangle(coords{44, 44}, clr, coords{position.x + 2, position.y + 2})
	return &arr
}
