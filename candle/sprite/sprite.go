package sprite

import "github.com/hajimehoshi/ebiten/v2"

type Sprite struct {
	Img  *ebiten.Image
	X, Y int
}
