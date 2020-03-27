package main

import (
	"math/rand"

	tl "github.com/JoelOtter/termloop"
)

// impl Drawable, DynamicPhysical
type Pellet struct {
	*tl.Entity
	coord Coord
}

func NewPellet() *Pellet {
	pellet := new(Pellet)
	pellet.Entity = tl.NewEntity(1, 1, 2, 1)
	pellet.move()
	return pellet
}

func (pellet *Pellet) Draw(screen *tl.Screen) {
	RenderSquare(screen, pellet.coord, ColoredCell(tl.ColorRed))
}

func (pellet *Pellet) move() {
	x, y := (rand.Intn(10)*2)+2, rand.Intn(20)+1
	pellet.coord.x, pellet.coord.y = x, y
	pellet.SetPosition(x, y)
}
