package utils

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
