package sprite

import "github.com/hajimehoshi/ebiten/v2"

func (s *Sprite) Draw(surface *ebiten.Image) {
	if s.X > surface.Bounds().Dx() || s.Y > surface.Bounds().Dy() {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(s.X), float64(s.Y))
		surface.DrawImage(s.Img, op)
	}
}
