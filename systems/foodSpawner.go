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
	nutrient components.Nutrient
}

var existingFoods = []food{
	{name: "Corn_baguette", asset: assets.Corn_baguette, nutrient: components.Nutrient{Corn: 2, KCal: 100}},
	{name: "Corn_bread", asset: assets.Corn_bread, nutrient: components.Nutrient{Corn: 3, KCal: 100}},
	{name: "Corn_rice", asset: assets.Corn_rice, nutrient: components.Nutrient{Corn: 2, KCal: 100}},
	{name: "Dairy_cheese", asset: assets.Dairy_cheese, nutrient: components.Nutrient{Dairy: 2, KCal: 100}},
	{name: "Dairy_milk", asset: assets.Dairy_milk, nutrient: components.Nutrient{Dairy: 2, KCal: 100}},
	{name: "Drink_beer", asset: assets.Drink_beer, nutrient: components.Nutrient{Drink: 3, KCal: 100}},
	{name: "Drink_coffee", asset: assets.Drink_coffee, nutrient: components.Nutrient{Drink: 1, KCal: 100}},
	{name: "Drink_juice", asset: assets.Drink_juice, nutrient: components.Nutrient{Drink: 1, KCal: 100}},
	{name: "Drink_tea", asset: assets.Drink_tea, nutrient: components.Nutrient{Drink: 1, KCal: 100}},
	{name: "Drink_water", asset: assets.Drink_water, nutrient: components.Nutrient{Drink: 1, KCal: 100}},
	{name: "Fish_crab", asset: assets.Fish_crab, nutrient: components.Nutrient{Fish: 3, KCal: 100}},
	{name: "Fish_sushi", asset: assets.Fish_sushi, nutrient: components.Nutrient{Fish: 2, KCal: 100}},
	{name: "Fruit_apple", asset: assets.Fruit_apple, nutrient: components.Nutrient{Fruit: 2, KCal: 100}},
	{name: "Fruit_banana", asset: assets.Fruit_banana, nutrient: components.Nutrient{Fruit: 2, KCal: 100}},
	{name: "Fruit_grapes", asset: assets.Fruit_grapes, nutrient: components.Nutrient{Fruit: 1, KCal: 100}},
	{name: "Fruit_orange", asset: assets.Fruit_orange, nutrient: components.Nutrient{Fruit: 2, KCal: 100}},
	{name: "Fruit_strawberry", asset: assets.Fruit_strawberry, nutrient: components.Nutrient{Fruit: 1, KCal: 100}},
	{name: "Meat_steak", asset: assets.Meat_steak, nutrient: components.Nutrient{Meat: 3, KCal: 100}},
	{name: "Treat_cupcake", asset: assets.Treat_cupcake, nutrient: components.Nutrient{Treat: 1, KCal: 100}},
	{name: "Treat_donut", asset: assets.Treat_donut, nutrient: components.Nutrient{Treat: 1, KCal: 100}},
	{name: "Vegetable_carrot", asset: assets.Vegetable_carrot, nutrient: components.Nutrient{Vegetable: 1, KCal: 100}},
	{name: "Vegetable_eggplant", asset: assets.Vegetable_eggplant, nutrient: components.Nutrient{Vegetable: 1, KCal: 100}},
	{name: "Vegetable_potato", asset: assets.Vegetable_potato, nutrient: components.Nutrient{Vegetable: 2, KCal: 100}},
	{name: "Vegetable_tomato", asset: assets.Vegetable_tomato, nutrient: components.Nutrient{Vegetable: 1, KCal: 100}},
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

	food := newFood(s.windowWidth, foodSpawner.Velocity.Y)
	s.manager.AddEntity(food)

	foodSpawner.CoolDown = foodSpawner.Rate
	return nil
}

func newFood(width int, yVelocity float64) *entities.Entity {

	padding := 10
	idx := rand.Intn(len(existingFoods))
	candidate := existingFoods[idx]
	posX := rand.Intn(width-padding*3) + padding
	entity := entities.Entity{Active: true}

	sprite := components.NewSprite(candidate.name, candidate.asset)
	spriteWidth, spriteHeight := sprite.Image.Size()

	entity.AddComponent(sprite)
	entity.AddComponent(&candidate.nutrient)

	entity.AddComponent(&components.Position{X: float64(posX), Y: -100})
	entity.AddComponent(&components.Dimension{Width: float64(spriteWidth), Height: float64(spriteHeight)})
	entity.AddComponent(&components.Velocity{X: 0, Y: yVelocity})
	entity.AddComponent(&components.Collision{})

	return &entity
}
