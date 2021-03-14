package foodzy

import (
	"bytes"
	"image"
	_ "image/png"
)

func LoadImage(input []byte) (image.Image, error) {
	image, _, err := image.Decode(bytes.NewReader(input))
	if err != nil {
		return image, err
	}
	return image, nil
}
