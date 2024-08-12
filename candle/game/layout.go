package game

import "candle/config"

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.Global.ScreenWidth, config.Global.ScreenHeight
}
