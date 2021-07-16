package entities

import (
	"github.com/co0p/foodzy/assets"
	"github.com/co0p/foodzy/components"
	"math/rand"
)

func NewRandomFood(width int) Entity {
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
	posX := rand.Intn(width)
	entity := Entity{Active: true}

	entity.AddComponent(components.NewSprite("food", foodAssets[idx]))
	entity.AddComponent(&components.Position{X: float64(posX), Y: -100})
	entity.AddComponent(&components.Velocity{X: 0, Y: 3})

	return entity
}
