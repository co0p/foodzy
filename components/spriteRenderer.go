package components

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type SpriteRenderer struct {
	Container *Element
	image     *ebiten.Image
}

func NewSpriteRenderer(element *Element, imageBytes []byte) *SpriteRenderer {
	image, err := newImageFromBytes(imageBytes)
	if err != nil {
		panic(fmt.Errorf("failed to create image from bytes: %s", err))
	}
	return &SpriteRenderer{
		Container: element,
		image:     ebiten.NewImageFromImage(image),
	}
}

func (c *SpriteRenderer) OnDraw(screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	/*
		w, h := c.image.Size()
		sw, sh := screen.Size()

		//tx := float64(sw/2+w/h-w/2) + c.Container.Position.X
		//ty := float64(sh-h-h/2) + c.Container.Position.Y
	*/

	tx := c.Container.Position.X
	ty := c.Container.Position.Y

	op.GeoM.Translate(tx, ty)
	op.GeoM.Rotate(c.Container.Rotation)
	screen.DrawImage(c.image, op)

	return nil
}

func (c *SpriteRenderer) OnUpdate() error {
	// nothing to do
	return nil
}

func newImageFromBytes(input []byte) (image.Image, error) {
	image, _, err := image.Decode(bytes.NewReader(input))
	if err != nil {
		return image, err
	}
	return image, nil
}

func (c *SpriteRenderer) size() (int, int) {
	return c.image.Size()
}
