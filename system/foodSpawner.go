package system

import (
	"github.com/co0p/foodzy/assets"
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/entity"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"math/rand"
)

type FoodDefintion struct {
	name     string
	asset    []byte
	nutrient component.Nutrient
}

var foodList = []FoodDefintion{
	{name: "Corn_baguette", asset: assets.Corn_baguette, nutrient: component.Nutrient{Corn: 2, KCal: 100}},
	{name: "Corn_bread", asset: assets.Corn_bread, nutrient: component.Nutrient{Corn: 3, KCal: 100}},
	{name: "Corn_rice", asset: assets.Corn_rice, nutrient: component.Nutrient{Corn: 2, KCal: 100}},
	{name: "Dairy_cheese", asset: assets.Dairy_cheese, nutrient: component.Nutrient{Dairy: 2, KCal: 100}},
	{name: "Dairy_milk", asset: assets.Dairy_milk, nutrient: component.Nutrient{Dairy: 2, KCal: 100}},
	{name: "Drink_beer", asset: assets.Drink_beer, nutrient: component.Nutrient{Drink: 3, KCal: 100}},
	{name: "Drink_coffee", asset: assets.Drink_coffee, nutrient: component.Nutrient{Drink: 1, KCal: 100}},
	{name: "Drink_juice", asset: assets.Drink_juice, nutrient: component.Nutrient{Drink: 1, KCal: 100}},
	{name: "Drink_tea", asset: assets.Drink_tea, nutrient: component.Nutrient{Drink: 1, KCal: 100}},
	{name: "Drink_water", asset: assets.Drink_water, nutrient: component.Nutrient{Drink: 1, KCal: 100}},
	{name: "Fish_crab", asset: assets.Fish_crab, nutrient: component.Nutrient{Fish: 3, KCal: 100}},
	{name: "Fish_sushi", asset: assets.Fish_sushi, nutrient: component.Nutrient{Fish: 2, KCal: 100}},
	{name: "Fruit_apple", asset: assets.Fruit_apple, nutrient: component.Nutrient{Fruit: 2, KCal: 100}},
	{name: "Fruit_banana", asset: assets.Fruit_banana, nutrient: component.Nutrient{Fruit: 2, KCal: 100}},
	{name: "Fruit_grapes", asset: assets.Fruit_grapes, nutrient: component.Nutrient{Fruit: 1, KCal: 100}},
	{name: "Fruit_orange", asset: assets.Fruit_orange, nutrient: component.Nutrient{Fruit: 2, KCal: 100}},
	{name: "Fruit_strawberry", asset: assets.Fruit_strawberry, nutrient: component.Nutrient{Fruit: 1, KCal: 100}},
	{name: "Meat_steak", asset: assets.Meat_steak, nutrient: component.Nutrient{Meat: 3, KCal: 100}},
	{name: "Treat_cupcake", asset: assets.Treat_cupcake, nutrient: component.Nutrient{Treat: 1, KCal: 100}},
	{name: "Treat_donut", asset: assets.Treat_donut, nutrient: component.Nutrient{Treat: 1, KCal: 100}},
	{name: "Vegetable_carrot", asset: assets.Vegetable_carrot, nutrient: component.Nutrient{Vegetable: 1, KCal: 100}},
	{name: "Vegetable_eggplant", asset: assets.Vegetable_eggplant, nutrient: component.Nutrient{Vegetable: 1, KCal: 100}},
	{name: "Vegetable_potato", asset: assets.Vegetable_potato, nutrient: component.Nutrient{Vegetable: 2, KCal: 100}},
	{name: "Vegetable_tomato", asset: assets.Vegetable_tomato, nutrient: component.Nutrient{Vegetable: 1, KCal: 100}},
}

// FoodSpawningSystem is responsible for spawning random food
type FoodSpawningSystem struct {
	manager     *entity.Manager
	windowWidth int
}

func NewFoodSpawningSystem(manager *entity.Manager, windowWidth int) *FoodSpawningSystem {
	return &FoodSpawningSystem{
		manager:     manager,
		windowWidth: windowWidth,
	}
}

func (s *FoodSpawningSystem) Draw(image *ebiten.Image) {}

func (s *FoodSpawningSystem) Update() error {
	e := s.manager.QueryByComponents(component.FoodSpawnerType)

	if len(e) != 1 {
		log.Panic("expected to find exactly one spawner")
	}
	foodSpawner := e[0].GetComponent(component.FoodSpawnerType).(*component.FoodSpawner)

	if foodSpawner.CoolDown > 0 {
		foodSpawner.CoolDown--
		return nil
	}

	padding := 10
	idx := rand.Intn(len(foodList))
	candidate := foodList[idx]
	posX := rand.Intn(s.windowWidth-padding*3) + padding
	posY := -200

	sprite := component.NewSprite(candidate.name, candidate.asset)
	position := component.Transform{X: float64(posX), Y: float64(posY), Scale: 0.7}
	velocity := component.Velocity{X: foodSpawner.Velocity.X, Y: foodSpawner.Velocity.Y}

	food := entity.NewFood(&candidate.nutrient, sprite, &velocity, &position)
	s.manager.AddEntity(food)

	foodSpawner.CoolDown = foodSpawner.Rate
	return nil
}
