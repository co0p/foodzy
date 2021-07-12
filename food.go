package foodzy

import (
	"github.com/co0p/foodzy/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
)

type Food struct {
	cooldown int
	active   bool
	image    *ebiten.Image
	name     string
	px       int
	py       int
}

func (e *Food) Update(g *Game) {
	// falling
	e.py = e.py + 5

	// reset when off screen
	if e.py > ScreenHeight {
		e.px, e.py = e.randomPosition()
	}
}

func (e *Food) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(e.px), float64(e.py))
	screen.DrawImage(e.image, op)
}

func (e *Food) randomPosition() (int, int) {
	imageWidth, imageHeight := e.image.Size()
	px := rand.Intn(ScreenWidth-imageWidth) + imageWidth
	py := -1*rand.Intn(ScreenHeight-imageHeight) + imageHeight
	return px, py
}

func NewFood(name string, imageBytes []byte) *Food {
	img, _ := utils.LoadImage(imageBytes)
	food := &Food{
		name:  name,
		image: ebiten.NewImageFromImage(img),
	}
	food.px, food.py = food.randomPosition()
	return food
}
