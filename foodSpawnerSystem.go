package foodzy

import (
	"github.com/co0p/foodzy/asset"
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"math/rand"
)

const initialYPosition float64 = -200

var foods = []struct {
	food  component.Food
	asset []byte
}{
	{component.Food{Value: -4}, asset.Corn_baguette},
	{component.Food{Value: -3}, asset.Corn_bread},
	{component.Food{Value: 1}, asset.Corn_rice},
	{component.Food{Value: 2}, asset.Dairy_cheese},
	{component.Food{Value: 3}, asset.Dairy_milk},
	{component.Food{Value: -10}, asset.Drink_beer},
	{component.Food{Value: 1}, asset.Drink_coffee},
	{component.Food{Value: 3}, asset.Drink_juice},
	{component.Food{Value: 5}, asset.Drink_tea},
	{component.Food{Value: 25}, asset.Drink_water},
	{component.Food{Value: 2}, asset.Fish_crab},
	{component.Food{Value: 4}, asset.Fish_sushi},
	{component.Food{Value: 10}, asset.Fruit_apple},
	{component.Food{Value: 10}, asset.Fruit_banana},
	{component.Food{Value: 8}, asset.Fruit_grapes},
	{component.Food{Value: 6}, asset.Fruit_orange},
	{component.Food{Value: 2}, asset.Fruit_strawberry},
	{component.Food{Value: -5}, asset.Meat_steak},
	{component.Food{Value: -20}, asset.Treat_cupcake},
	{component.Food{Value: -20}, asset.Treat_donut},
	{component.Food{Value: 8}, asset.Vegetable_carrot},
	{component.Food{Value: 6}, asset.Vegetable_eggplant},
	{component.Food{Value: 2}, asset.Vegetable_potato},
	{component.Food{Value: 6}, asset.Vegetable_tomato},
}

// FoodSpawningSystem is responsible for spawning random food
type FoodSpawningSystem struct {
	manager *ecs.EntityManager
}

func NewFoodSpawningSystem(manager *ecs.EntityManager) *FoodSpawningSystem {
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
	idx := rand.Intn(len(foods))
	candidate := foods[idx]

	posX := float64(rand.Intn(ScreenWidth-padding*4) + padding)
	posY := initialYPosition

	position := component.Transform{X: posX, Y: posY, Scale: 0.7}
	velocity := component.Velocity{X: foodSpawner.Velocity.X, Y: foodSpawner.Velocity.Y}
	sprite := component.NewSprite("", candidate.asset)

	food := NewFood(&candidate.food, sprite, &velocity, &position)
	s.manager.AddEntity(food)

	foodSpawner.CoolDown = foodSpawner.Rate
	return nil
}
