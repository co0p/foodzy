package components

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// KeyboardMover moves the element via keyboard controls
type KeyboardMover struct {
	container *Element
	speed     float64

	sr          *SpriteRenderer
	screenWidth float64
}

func NewKeyboardMover(container *Element, speed float64, screenWidth float64) *KeyboardMover {
	return &KeyboardMover{
		speed:       speed,
		container:   container,
		screenWidth: screenWidth,
		sr:          container.GetComponent(&SpriteRenderer{}).(*SpriteRenderer),
	}
}

func (c *KeyboardMover) OnUpdate() error {

	w, _ := containerSize(c)

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		newLeft := c.container.Position.X - c.speed
		if newLeft > 0 {
			c.container.Position.X = newLeft
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		newRight := c.container.Position.X + c.speed
		if newRight < (c.screenWidth - float64(w)) {
			c.container.Position.X = newRight
		}
	}

	return nil
}

func containerSize(c *KeyboardMover) (int, int) {
	renderer := c.container.GetComponent(&SpriteRenderer{}).(*SpriteRenderer)
	return renderer.size()
}

func (c *KeyboardMover) OnDraw(_ *ebiten.Image) error {
	return nil
}
