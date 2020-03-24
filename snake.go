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
	snake.move()
	for _, coord := range snake.coords {
		RenderSquare(screen, coord, ColoredCell(tl.ColorGreen))
	}
}

func (snake *Snake) Tick(event tl.Event) {
	switch event.Key {
	case tl.KeyArrowDown:
		if snake.dir != north {
			snake.dir = south
		}
	case tl.KeyArrowUp:
		if snake.dir != south {
			snake.dir = north
		}
	case tl.KeyArrowLeft:
		if snake.dir != east {
			snake.dir = west
		}
	case tl.KeyArrowRight:
		if snake.dir != west {
			snake.dir = east
		}
	}
}

func (snake *Snake) move() {
	cs := snake.coords
	head := cs[len(cs)-1]
	snake.coords = append(cs[1:], nextCoord(head, snake.dir))
}

func nextCoord(coord Coord, dir Direction) Coord {
	x, y := coord.x, coord.y
	switch dir {
	case north:
		return Coord{x, y - 1}
	case south:
		return Coord{x, y + 1}
	case east:
		return Coord{x + 2, y}
	case west:
		return Coord{x - 2, y}
	default:
		panic("unknown direction")
	}
}
