package game

import (
	"candle/config"
	"fmt"
	"log"
)

func (g *Game) Update() error {
	g.Inputs.GetInputs()

	if g.Inputs.Debug {
		config.Global.DebugEnable = !config.Global.DebugEnable
	}

	switch g.GameState {
	case mainMenu:
		g.updateMainMenu()
	case inGame:
		g.updateInGame()
	default:
		return fmt.Errorf("unknown game state : %d", g.GameState)
	}

	return nil
}

func (g *Game) updateMainMenu() {
	if g.Inputs.Enter {
		g.GameState = inGame
	}
}

func (g *Game) updateInGame() {
	if g.Inputs.Menu {
		g.GameState = mainMenu
	} else {
		// TODO : game loop
		log.Println("todo")
	}
}
