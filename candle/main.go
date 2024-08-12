package main

import (
	"candle/config"
	"candle/game"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	config.Init(256, 240, 2) // NES settings
	g := &game.Game{}
	g.Init()

	ebiten.SetWindowSize(config.Global.ScreenWidth*config.Global.ScreenScaleCoef, config.Global.ScreenHeight*config.Global.ScreenScaleCoef)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Wick's End")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
