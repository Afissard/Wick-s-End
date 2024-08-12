package game

import (
	"candle/inputs"
)

type Game struct {
	GameState int
	Inputs    inputs.Inputs
}

const ( // game state
	mainMenu   int = iota // stater screen
	inGame                // the game
	inGameMenu            // [escape] menu for inventory, save and settings
)
