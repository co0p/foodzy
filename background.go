package foodzy

import (
	"github.com/co0p/foodzy/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Background struct {
	image *ebiten.Image
}

func (b *Background) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(b.image, op)
}

func (b *Background) Update(game *Game) {}

func NewBackground() *Background {
	img, _ := LoadImage(assets.Background)

	background := &Background{
		image: ebiten.NewImageFromImage(img),
	}
	log.Printf("created background")
	return background
}
