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
	lastButt Coord
	dir    Direction
}

func NewSnake() *Snake {
	snake := new(Snake)
	snake.coords = []Coord{
		{10, 6},
		{10, 7},
		{10, 8}, // head
	}
	snake.lastButt = Coord{10, 5}
	// Entity follows the head
	snake.Entity = tl.NewEntity(10, 8, 2, 1)
	snake.dir = south
	return snake
}

func (snake *Snake) Draw(screen *tl.Screen) {
	if !gameOver {
		snake.move()
	}
	gameOver = isGameOver(snake)
	for _, coord := range snake.coords {
		RenderSquare(screen, coord, ColoredCell(tl.ColorGreen))
	}
}

func (snake *Snake) Tick(event tl.Event) {
	//game.Log("Snake event, %+v", event)
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

func (snake *Snake) Collide(phys tl.Physical) {
	game.Log("Snake Collision: %#v", phys)
	switch p := phys.(type) {
	case *Pellet:
		p.move()
		snake.coords = append([]Coord{snake.lastButt}, snake.coords...)
	}
}

func (snake *Snake) head() Coord {
	cs := snake.coords
	return cs[len(cs)-1]
}

func (snake *Snake) move() {
	nextHead := nextCoord(snake.head(), snake.dir)
	snake.lastButt = snake.coords[0]
	snake.coords = append(snake.coords[1:], nextHead)
	snake.SetPosition(nextHead.x, nextHead.y)
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

func isGameOver(snake *Snake) bool {
	return isOutOfBounds(snake)
}

func isOutOfBounds(snake *Snake) bool {
	h := snake.head()
	return h.x < 2 || h.y < 2 || h.x > 40 || h.y > 20
}
