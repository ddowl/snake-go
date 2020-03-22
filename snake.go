package main

import tl "github.com/JoelOtter/termloop"

func main() {
	game := tl.NewGame()
	screen := game.Screen()
	screen.SetLevel(tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlue,
		Fg: tl.ColorWhite,
		Ch: ' ',
	}))
	screen.AddEntity(tl.NewRectangle(2, 1, 40, 20, tl.RgbTo256Color(201, 189, 91)))
	screen.SetFps(30)
	game.Start()
}
