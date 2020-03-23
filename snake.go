package main

import tl "github.com/JoelOtter/termloop"

type Direction int

const (
	north Direction = iota
	south
	east
	west
)

type Snake struct {
	*tl.Entity
	coords []Coord
	dir    Direction
}

func NewSnake() *Snake {
	snake := new(Snake)
	snake.Entity = tl.NewEntity(1, 1, 2, 1)

	snake.coords = []Coord{
		{10, 6},
		{10, 7},
		{10, 8}, // head
	}
	snake.dir = south
	return snake
}

func (snake *Snake) Draw(screen *tl.Screen) {
	for _, coord := range snake.coords {
		RenderSquare(screen, coord, ColoredCell(tl.ColorGreen))
	}
}
