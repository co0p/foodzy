package component

import (
	"image/color"
)

type Text struct {
	Value string
	Color color.Color
}

const TextType ComponentType = "Text"

func (t Text) Type() ComponentType {
	return TextType
}
