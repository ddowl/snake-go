package main

import tl "github.com/JoelOtter/termloop"

// impl Drawable, DynamicPhysical
type Pellet struct {
	*tl.Entity
	coord Coord
}

func NewPellet() *Pellet {
	pellet := new(Pellet)
	pellet.Entity = tl.NewEntity(1, 1, 2, 1)
	x, y := 6, 4
	pellet.coord.x = x
	pellet.coord.y = y
	pellet.SetPosition(x, y)
	return pellet
}

func (pellet *Pellet) Draw(screen *tl.Screen) {
	RenderSquare(screen, pellet.coord, ColoredCell(tl.ColorRed))
}

func (pellet *Pellet) Position() (int, int) {
	return pellet.coord.x, pellet.coord.y
}

func (pellet *Pellet) Size() (int, int) {
	return 1, 1
}

func (pellet *Pellet) Collide(tl.Physical) {

}
