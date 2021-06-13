package foodzy

import (
	"github.com/co0p/foodzy/assets"
	"github.com/co0p/foodzy/components"
	"github.com/hajimehoshi/ebiten/v2"
)

const playerSize = 50

type Player struct {
	image      *ebiten.Image
	posX, posY float64
}

func NewPlayerElement() *components.Element {

	pos := components.Vector{
		X: float64(ScreenWidth / 2.0),
		Y: float64(ScreenHeight - playerSize*2.0),
	}
	playerElement := components.NewElement("player", true, pos, 0)

	playerElement.AddComponent(components.NewSpriteRenderer(playerElement, assets.Plate))
	playerElement.AddComponent(components.NewKeyboardMover(playerElement, 5, float64(ScreenWidth)))
	playerElement.AddComponent(components.NewLoggingComponent(playerElement))

	return playerElement
}
