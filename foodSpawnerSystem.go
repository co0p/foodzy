package foodzy

import (
	"github.com/co0p/foodzy/asset"
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
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
	{name: "Corn_baguette", asset: asset.Corn_baguette, nutrient: component.Nutrient{Corn: 2, KCal: 100}},
	{name: "Corn_bread", asset: asset.Corn_bread, nutrient: component.Nutrient{Corn: 3, KCal: 100}},
	{name: "Corn_rice", asset: asset.Corn_rice, nutrient: component.Nutrient{Corn: 2, KCal: 100}},
	{name: "Dairy_cheese", asset: asset.Dairy_cheese, nutrient: component.Nutrient{Dairy: 2, KCal: 100}},
	{name: "Dairy_milk", asset: asset.Dairy_milk, nutrient: component.Nutrient{Dairy: 2, KCal: 100}},
	{name: "Drink_beer", asset: asset.Drink_beer, nutrient: component.Nutrient{Drink: 3, KCal: 100}},
	{name: "Drink_coffee", asset: asset.Drink_coffee, nutrient: component.Nutrient{Drink: 1, KCal: 100}},
	{name: "Drink_juice", asset: asset.Drink_juice, nutrient: component.Nutrient{Drink: 1, KCal: 100}},
	{name: "Drink_tea", asset: asset.Drink_tea, nutrient: component.Nutrient{Drink: 1, KCal: 100}},
	{name: "Drink_water", asset: asset.Drink_water, nutrient: component.Nutrient{Drink: 1, KCal: 100}},
	{name: "Fish_crab", asset: asset.Fish_crab, nutrient: component.Nutrient{Fish: 3, KCal: 100}},
	{name: "Fish_sushi", asset: asset.Fish_sushi, nutrient: component.Nutrient{Fish: 2, KCal: 100}},
	{name: "Fruit_apple", asset: asset.Fruit_apple, nutrient: component.Nutrient{Fruit: 2, KCal: 100}},
	{name: "Fruit_banana", asset: asset.Fruit_banana, nutrient: component.Nutrient{Fruit: 2, KCal: 100}},
	{name: "Fruit_grapes", asset: asset.Fruit_grapes, nutrient: component.Nutrient{Fruit: 1, KCal: 100}},
	{name: "Fruit_orange", asset: asset.Fruit_orange, nutrient: component.Nutrient{Fruit: 2, KCal: 100}},
	{name: "Fruit_strawberry", asset: asset.Fruit_strawberry, nutrient: component.Nutrient{Fruit: 1, KCal: 100}},
	{name: "Meat_steak", asset: asset.Meat_steak, nutrient: component.Nutrient{Meat: 3, KCal: 100}},
	{name: "Treat_cupcake", asset: asset.Treat_cupcake, nutrient: component.Nutrient{Treat: 1, KCal: 100}},
	{name: "Treat_donut", asset: asset.Treat_donut, nutrient: component.Nutrient{Treat: 1, KCal: 100}},
	{name: "Vegetable_carrot", asset: asset.Vegetable_carrot, nutrient: component.Nutrient{Vegetable: 1, KCal: 100}},
	{name: "Vegetable_eggplant", asset: asset.Vegetable_eggplant, nutrient: component.Nutrient{Vegetable: 1, KCal: 100}},
	{name: "Vegetable_potato", asset: asset.Vegetable_potato, nutrient: component.Nutrient{Vegetable: 2, KCal: 100}},
	{name: "Vegetable_tomato", asset: asset.Vegetable_tomato, nutrient: component.Nutrient{Vegetable: 1, KCal: 100}},
}

// FoodSpawningSystem is responsible for spawning random food
type FoodSpawningSystem struct {
	manager *ecs.Manager
}

func NewFoodSpawningSystem(manager *ecs.Manager) *FoodSpawningSystem {
	return &FoodSpawningSystem{
		manager: manager,
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
	posX := rand.Intn(ScreenWidth-padding*3) + padding
	posY := -200

	sprite := component.NewSprite(candidate.name, candidate.asset)
	position := component.Transform{X: float64(posX), Y: float64(posY), Scale: 0.7}
	velocity := component.Velocity{X: foodSpawner.Velocity.X, Y: foodSpawner.Velocity.Y}

	food := NewFood(&candidate.nutrient, sprite, &velocity, &position)
	s.manager.AddEntity(food)

	foodSpawner.CoolDown = foodSpawner.Rate
	return nil
}
