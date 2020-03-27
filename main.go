package main

import (
	tl "github.com/JoelOtter/termloop"
	"math/rand"
	"time"
)

var gameOver bool

var game *tl.Game
func main() {
	rand.Seed(time.Now().UnixNano())
	game = tl.NewGame()
	game.SetDebugOn(true)

	gameOver = false

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlue,
		Fg: tl.ColorWhite,
		Ch: ' ',
	})
	level.AddEntity(tl.NewRectangle(2, 1, 40, 20, tl.RgbTo256Color(201, 189, 91)))
	level.AddEntity(NewPellet())
	level.AddEntity(NewSnake())

	screen := game.Screen()
	screen.SetLevel(level)
	screen.SetFps(10)
	game.Start()
}
