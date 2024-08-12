package game

import (
	"candle/config"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *Game) Draw(screen *ebiten.Image) {
	// Debug
	if config.Global.DebugEnable {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %.2f, FPS: %.2f", ebiten.ActualTPS(), ebiten.ActualFPS()))
	}

	switch g.GameState {
	case mainMenu:
		g.drawMainMenu()
	case inGame:
		g.drawInGame()
	default:
		ebitenutil.DebugPrint(screen, "not implemented")
	}
}

func (g *Game) drawMainMenu() {
	// TODO
}

func (g *Game) drawInGame() {
	// TODO
}
