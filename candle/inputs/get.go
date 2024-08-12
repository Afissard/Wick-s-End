package inputs

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (i *Inputs) GetInputs() {
	// Keyboard
	// * INFO: Ebiten only use qwerty layout -> no need for key-mapping for azerty
	// Directional inputs
	i.Left = ebiten.IsKeyPressed(ebiten.KeyA)
	i.Right = ebiten.IsKeyPressed(ebiten.KeyD)
	i.Up = ebiten.IsKeyPressed(ebiten.KeyW)
	i.Down = ebiten.IsKeyPressed(ebiten.KeyS)

	// In-game inputs
	// TODO : choose key
	// Special inputs
	// TODO: add delay for update ... -> currently may causing flickering
	i.Menu = inpututil.IsKeyJustPressed(ebiten.KeyEscape)
	i.Enter = inpututil.IsKeyJustPressed(ebiten.KeyEnter)
	i.Debug = inpututil.IsKeyJustPressed(ebiten.KeyF3)

	// Game Controller
	// TODO: setup controller support + inputs receiver
	// Directional inputs
	// In-game inputs
	// Special inputs
}
