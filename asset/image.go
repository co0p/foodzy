package asset

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
)

func LoadImage(input []byte) (*ebiten.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(input))
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(img), nil
}

type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

func LoadSpriteSheet(input []byte, width int, height int, count int) ([]*ebiten.Image, error) {
	sprites := []*ebiten.Image{}

	img, _, err := image.Decode(bytes.NewReader(input))
	if err != nil {
		return nil, err
	}

	subImage := img.(SubImager)
	for i := 0; i < count; i++ {
		rect := image.Rect(i*width, 0, (i+1)*width, height)
		sprite := subImage.SubImage(rect)

		fromImage := ebiten.NewImageFromImage(sprite)
		sprites = append(sprites, fromImage)
	}

	return sprites, nil
}
