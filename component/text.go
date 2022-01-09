package component

import (
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"image/color"
)

type Text struct {
	Value string
	Color color.Color
	Font  *font.Face
}

const TextType ecs.ComponentType = "Text"

func (t Text) Type() ecs.ComponentType {
	return TextType
}

func (t *Text) Dimensions() (float64, float64) {

	if t.Value == "" || t.Font == nil {
		return 0, 0
	}
	width := text.BoundString(*t.Font, t.Value).Dx()
	height := text.BoundString(*t.Font, t.Value).Dy()
	return float64(width), float64(height)
}

func (t *Text) RelativeCenter(screenWidth, screenHeight int) (float64, float64) {

	if t.Value == "" || t.Font == nil {
		return 0, 0
	}
	width := text.BoundString(*t.Font, t.Value).Dx()
	height := text.BoundString(*t.Font, t.Value).Dy()
	return float64(screenWidth-width) / 2, float64(screenHeight-height) / 2
}
