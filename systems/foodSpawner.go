package systems

import (
	"github.com/co0p/foodzy/assets"
	"github.com/co0p/foodzy/components"
	"github.com/co0p/foodzy/entities"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"math/rand"
)

type food struct {
	name     string
	asset    []byte
	foodType components.FoodType
}

var existingFood = []food{
	{name: "Corn_baguette", asset: assets.Corn_baguette, foodType: components.FoodCorn},
	{name: "Corn_bread", asset: assets.Corn_bread, foodType: components.FoodCorn},
	{name: "Corn_rice", asset: assets.Corn_rice, foodType: components.FoodCorn},
	{name: "Dairy_cheese", asset: assets.Dairy_cheese, foodType: components.FoodDairy},
	{name: "Dairy_milk", asset: assets.Dairy_milk, foodType: components.FoodDairy},
	{name: "Drink_beer", asset: assets.Drink_beer, foodType: components.FoodDrink},
	{name: "Drink_coffee", asset: assets.Drink_coffee, foodType: components.FoodDrink},
	{name: "Drink_juice", asset: assets.Drink_juice, foodType: components.FoodDrink},
	{name: "Drink_tea", asset: assets.Drink_tea, foodType: components.FoodDrink},
	{name: "Drink_water", asset: assets.Drink_water, foodType: components.FoodDrink},
	{name: "Fish_crab", asset: assets.Fish_crab, foodType: components.FoodFish},
	{name: "Fish_sushi", asset: assets.Fish_sushi, foodType: components.FoodFish},
	{name: "Fruit_apple", asset: assets.Fruit_apple, foodType: components.FoodFruit},
	{name: "Fruit_banana", asset: assets.Fruit_banana, foodType: components.FoodFruit},
	{name: "Fruit_grapes", asset: assets.Fruit_grapes, foodType: components.FoodFruit},
	{name: "Fruit_orange", asset: assets.Fruit_orange, foodType: components.FoodFruit},
	{name: "Fruit_strawberry", asset: assets.Fruit_strawberry, foodType: components.FoodFruit},
	{name: "Meat_steak", asset: assets.Meat_steak, foodType: components.FoodMeat},
	{name: "Treat_cupcake", asset: assets.Treat_cupcake, foodType: components.FoodTreat},
	{name: "Treat_donut", asset: assets.Treat_donut, foodType: components.FoodTreat},
	{name: "Vegetable_carrot", asset: assets.Vegetable_carrot, foodType: components.FoodVegetable},
	{name: "Vegetable_eggplant", asset: assets.Vegetable_eggplant, foodType: components.FoodVegetable},
	{name: "Vegetable_potato", asset: assets.Vegetable_potato, foodType: components.FoodVegetable},
	{name: "Vegetable_tomato", asset: assets.Vegetable_tomato, foodType: components.FoodVegetable},
}

// FoodSpawningSystem is responsible for spawning random food
type FoodSpawningSystem struct {
	manager     *entities.Manager
	windowWidth int
}

func NewFoodSpawningSystem(manager *entities.Manager, windowWidth int) *FoodSpawningSystem {
	return &FoodSpawningSystem{
		manager:     manager,
		windowWidth: windowWidth,
	}
}

func (s *FoodSpawningSystem) Draw(image *ebiten.Image) {}

func (s *FoodSpawningSystem) Update() error {
	e := s.manager.QueryByComponents(&components.FoodSpawner{})

	if len(e) != 1 {
		log.Panic("expected to find exactly one spawner")
	}
	foodSpawner := e[0].GetComponent(&components.FoodSpawner{}).(*components.FoodSpawner)

	if foodSpawner.CoolDown > 0 {
		foodSpawner.CoolDown--
		return nil
	}

	food := newFood(foodSpawner.Types, s.windowWidth, foodSpawner.Velocity.Y)
	s.manager.AddEntity(food)

	foodSpawner.CoolDown = foodSpawner.Rate
	return nil
}

func newFood(types []components.FoodType, width int, yVelocity float64) *entities.Entity {

	// seriously? no better way ?
	candidates := []food{}
	for _, v := range existingFood {
		for _, t := range types {
			if v.foodType == t {
				candidates = append(candidates, v)
			}
		}
	}

	padding := 10
	idx := rand.Intn(len(candidates))
	candidate := candidates[idx]
	posX := rand.Intn(width-padding*3) + padding
	entity := entities.Entity{Active: true}

	sprite := components.NewSprite(candidate.name, candidate.asset)
	spriteWidth, spriteHeight := sprite.Image.Size()
	entity.AddComponent(sprite)

	entity.AddComponent(&components.Position{X: float64(posX), Y: -100})
	entity.AddComponent(&components.Dimension{Width: float64(spriteWidth), Height: float64(spriteHeight)})
	entity.AddComponent(&components.Velocity{X: 0, Y: yVelocity})
	entity.AddComponent(&components.Collision{})

	return &entity
}
