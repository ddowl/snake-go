package main

import tl "github.com/JoelOtter/termloop"

type Coord struct {
	x int
	y int
}

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
		RenderSquare(screen, pellet.coord, &tl.Cell{
			Fg: tl.ColorRed,
			Bg: tl.ColorRed,
			Ch: ' ',
		})
}

func (pellet *Pellet) Position() (int, int) {
	return pellet.coord.x, pellet.coord.y
}

func (pellet *Pellet) Size() (int, int) {
	return 1, 1
}

func (pellet *Pellet) Collide(tl.Physical) {

}

func RenderSquare(screen *tl.Screen, coord Coord, cell *tl.Cell) {
	screen.RenderCell(coord.x, coord.y, cell)
	screen.RenderCell(coord.x + 1, coord.y, cell)
}

func main() {
	game := tl.NewGame()

	screen := game.Screen()
	screen.SetLevel(tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlue,
		Fg: tl.ColorWhite,
		Ch: ' ',
	}))
	screen.AddEntity(tl.NewRectangle(2, 1, 40, 20, tl.RgbTo256Color(201, 189, 91)))
	screen.AddEntity(NewPellet())
	screen.SetFps(30)
	game.Start()
}
