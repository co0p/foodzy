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
	{name: "Bread", asset: assets.Bread, nutrient: components.Nutrient{Carbohydrates: 15, Fat: 2, Protein: 2, Minerals: 4, Water: 1}, toxic: components.Toxic{}},
	{name: "Carrot", asset: assets.Carrot, nutrient: components.Nutrient{Carbohydrates: 10, Protein: 1, Fat: 1, Water: 5}, toxic: components.Toxic{}},
	{name: "Beer", asset: assets.Beer, nutrient: components.Nutrient{Carbohydrates: 2, Minerals: 5, Vitamins: 5, Water: 40}, toxic: components.Toxic{}},
	{name: "Cheese", asset: assets.Cheese, nutrient: components.Nutrient{Carbohydrates: 1, Fat: 40, Vitamins: 1, Minerals: 5, Water: 6}, toxic: components.Toxic{}},
	{name: "Fish", asset: assets.Fish, nutrient: components.Nutrient{Protein: 26, Fat: 12, Vitamins: 2, Minerals: 3, Water: 12}, toxic: components.Toxic{}},
	{name: "Meat", asset: assets.Meat, nutrient: components.Nutrient{Protein: 19, Fat: 17, Minerals: 5, Water: 5}, toxic: components.Toxic{}},
	{name: "Lemon", asset: assets.Lemon, nutrient: components.Nutrient{Carbohydrates: 21, Vitamins: 20, Minerals: 6, Water: 10}, toxic: components.Toxic{}},
	{name: "Tomato", asset: assets.Tomato, nutrient: components.Nutrient{Carbohydrates: 10, Protein: 2, Vitamins: 5, Minerals: 5, Water: 4}, toxic: components.Toxic{}},
	{name: "Strawberry", asset: assets.Strawberry, nutrient: components.Nutrient{Carbohydrates: 12, Protein: 1, Vitamins: 4, Minerals: 2, Water: 10}, toxic: components.Toxic{}},
}

func NewRandomFood(screenWidth int) Entity {

	idx := rand.Intn(len(existingFoods))
	f := existingFoods[idx]
	posX := rand.Intn(screenWidth)
	entity := Entity{Active: true}

	sprite := components.NewSprite("food", f.asset)
	spriteWidth, spriteHeight := sprite.Image.Size()

	entity.AddComponent(sprite)
	entity.AddComponent(&components.Position{X: float64(posX), Y: -100})
	entity.AddComponent(&components.Dimension{Width: float64(spriteWidth), Height: float64(spriteHeight)})
	entity.AddComponent(&components.Velocity{X: 0, Y: 3})
	entity.AddComponent(&components.Collision{})
	entity.AddComponent(&f.nutrient)
	entity.AddComponent(&f.toxic)

	return entity
}
