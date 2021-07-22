package entities

import (
	"github.com/co0p/foodzy/assets"
	"github.com/co0p/foodzy/components"
	"math/rand"
)

func NewRandomFood(screenWidth int) Entity {
	foodAssets := [][]byte{
		assets.Bread,
		assets.Carrot,
		assets.Beer,
		assets.Cheese,
		assets.Fish,
		assets.Meat,
		assets.Lemon,
		assets.Tomato,
		assets.Strawberry,
	}

	idx := rand.Intn(len(foodAssets))
	posX := rand.Intn(screenWidth)
	entity := Entity{Active: true}

	sprite := components.NewSprite("food", foodAssets[idx])
	spriteWidth, spriteHeight := sprite.Image.Size()

	entity.AddComponent(sprite)
	entity.AddComponent(&components.Position{X: float64(posX), Y: -100})
	entity.AddComponent(&components.Dimension{Width: float64(spriteWidth), Height: float64(spriteHeight)})
	entity.AddComponent(&components.Velocity{X: 0, Y: 3})
	entity.AddComponent(&components.Collision{})

	return entity
}
