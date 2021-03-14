package foodzy

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Food struct {
	image  *ebiten.Image
	name   string
	vx, vy float64
	px     int
	py     int
}

func (e *Food) Update(g *Game) {}

func (e *Food) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	cw, ch := e.image.Size()
	sw, sh := screen.Size()

	tx := sw/2 - cw/2 + e.px
	ty := sh/2 - ch/2 + e.py
	op.GeoM.Translate(float64(tx), float64(ty))
	screen.DrawImage(e.image, op)
}
func NewFood(name string, imageBytes []byte) *Food {
	img, _ := LoadImage(imageBytes)

	food := &Food {
		name: name,
		image: ebiten.NewImageFromImage(img),
	}
	log.Printf("created foood: %s", name)
	return food
}
