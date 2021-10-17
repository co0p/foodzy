package components

import (
	"image/color"
)

type Text struct {
	Value string
	Color color.Color
}

func (p Text) ID() string {
	return ""
}
