package entities

import (
	"github.com/co0p/foodzy/assets"
	"github.com/co0p/foodzy/components"
	"math/rand"
)

type food struct {
	name     string
	asset    []byte
	nutrient components.Nutrient
	toxic    components.Toxic
}

var existingFoods = []food{
	{name: "Bread", asset: assets.Bread, nutrient: components.Nutrient{Carbohydrates: 100}, toxic: components.Toxic{}},
	{name: "Carrot", asset: assets.Carrot, nutrient: components.Nutrient{Vitamins: 123}, toxic: components.Toxic{}},
	{name: "Beer", asset: assets.Beer, nutrient: components.Nutrient{Carbohydrates: 80}, toxic: components.Toxic{}},
	{name: "Cheese", asset: assets.Cheese, nutrient: components.Nutrient{Fat: 20}, toxic: components.Toxic{}},
	{name: "Fish", asset: assets.Fish, nutrient: components.Nutrient{Fat: 15}, toxic: components.Toxic{}},
	{name: "Meat", asset: assets.Meat, nutrient: components.Nutrient{Fat: 20}, toxic: components.Toxic{}},
	{name: "Lemon", asset: assets.Lemon, nutrient: components.Nutrient{Vitamins: 20}, toxic: components.Toxic{}},
	{name: "Tomato", asset: assets.Tomato, nutrient: components.Nutrient{Vitamins: 10}, toxic: components.Toxic{}},
	{name: "Strawberry", asset: assets.Strawberry, nutrient: components.Nutrient{Vitamins: 1}, toxic: components.Toxic{}},
}

func NewRandomFood(screenWidth int) Entity {

	idx := rand.Intn(len(existingFoods))
	food := existingFoods[idx]
	posX := rand.Intn(screenWidth)
	entity := Entity{Active: true}

	sprite := components.NewSprite("food", food.asset)
	spriteWidth, spriteHeight := sprite.Image.Size()

	entity.AddComponent(sprite)
	entity.AddComponent(&components.Position{X: float64(posX), Y: -100})
	entity.AddComponent(&components.Dimension{Width: float64(spriteWidth), Height: float64(spriteHeight)})
	entity.AddComponent(&components.Velocity{X: 0, Y: 3})
	entity.AddComponent(&components.Collision{})
	entity.AddComponent(&food.nutrient)
	entity.AddComponent(&food.toxic)

	return entity
}
