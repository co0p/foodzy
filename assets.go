package foodzy

import (
	"bytes"
	"image"
	_ "image/png"
)

func LoadImage(input []byte) (image.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(input))
	if err != nil {
		return img, err
	}
	return img, nil
}
