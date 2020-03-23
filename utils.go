package main

import tl "github.com/JoelOtter/termloop"

type Coord struct {
	x int
	y int
}

func ColoredCell(color tl.Attr) *tl.Cell {
	return &tl.Cell{
		Fg: color,
		Bg: color,
		Ch: ' ',
	}
}

func RenderSquare(screen *tl.Screen, coord Coord, cell *tl.Cell) {
	screen.RenderCell(coord.x, coord.y, cell)
	screen.RenderCell(coord.x+1, coord.y, cell)
}
