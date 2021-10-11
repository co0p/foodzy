package foodzy

import (
	"github.com/co0p/foodzy/components"
	"github.com/co0p/foodzy/entities"
	"github.com/co0p/foodzy/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth    int    = 800
	ScreenHeight   int    = 600
	GameName       string = "Foodzy"
	SpawnFrequency int    = 40
)

type Game struct {
	entityManager *entities.Manager

	movementSystem   *systems.MovementSystem
	foodSystem       *systems.FoodManager
	controllerSystem *systems.ControllerSystem
	renderSystem     *systems.RenderSystem
	collisionSystem  *systems.CollisionSystem
	scoreSystem      *systems.ScoreSystem

	soundSystem   *systems.SoundSystem
	cleanupSystem *systems.CleanupSystem
}

func NewGame() *Game {

	entityManager := entities.Manager{}
	entityManager.AddEntity(entities.NewBackground())
	entityManager.AddEntity(entities.NewPlayer(ScreenWidth, ScreenHeight))

	entityManager.AddEntity(entities.NewScore("Water", components.Nutrient{Water: 1}))
	entityManager.AddEntity(entities.NewScore("Carbs", components.Nutrient{Carbohydrates: 1}))

	return &Game{
		entityManager:    &entityManager,
		movementSystem:   systems.NewMovementSystem(&entityManager),
		controllerSystem: systems.NewControllerSystem(&entityManager, ScreenWidth, ScreenHeight),
		renderSystem:     systems.NewRenderSystem(&entityManager),
		soundSystem:      systems.NewSoundSystem(),
		foodSystem:       systems.NewFoodSystem(&entityManager, SpawnFrequency, ScreenWidth),
		collisionSystem:  systems.NewCollisionSystem(&entityManager),
		scoreSystem:      systems.NewScoreSystem(&entityManager, ScreenWidth, ScreenHeight),
		cleanupSystem:    systems.NewCleanupSystem(&entityManager, 100, ScreenHeight),
	}
}

func (g *Game) Update() error {
	err := g.controllerSystem.Update()
	err = g.foodSystem.Update()
	err = g.movementSystem.Update()
	err = g.collisionSystem.Update()
	err = g.scoreSystem.Update()

	err = g.cleanupSystem.Update()
	return err
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderSystem.Draw(screen)
	g.scoreSystem.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
