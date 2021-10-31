package foodzy

import (
	"github.com/co0p/foodzy/components"
	"github.com/co0p/foodzy/entities"
	"github.com/co0p/foodzy/systems"
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
	"time"
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
	foodSystem       *systems.FoodSpawningSystem
	controllerSystem *systems.ControllerSystem
	collisionSystem  *systems.CollisionSystem
	scoreSystem      *systems.ScoreSystem

	spriteRenderSystem *systems.spriteRenderSystem
	textRenderSystem   *systems.TextRenderSystem

	soundSystem   *systems.SoundSystem
	cleanupSystem *systems.CleanupSystem
}

func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())

	entityManager := entities.Manager{}
	entityManager.AddEntity(entities.NewBackground())
	entityManager.AddEntity(entities.NewPlayer(ScreenWidth, ScreenHeight))
	entityManager.AddEntity(entities.NewFoodSpawner(SpawnFrequency))

	scores := constructScores()
	for _, v := range scores {
		entityManager.AddEntity(v)
	}

	return &Game{
		entityManager:      &entityManager,
		movementSystem:     systems.NewMovementSystem(&entityManager),
		controllerSystem:   systems.NewControllerSystem(&entityManager, ScreenWidth, ScreenHeight),
		soundSystem:        systems.NewSoundSystem(),
		foodSystem:         systems.NewFoodSpawningSystem(&entityManager, ScreenWidth),
		collisionSystem:    systems.NewCollisionSystem(&entityManager),
		scoreSystem:        systems.NewScoreSystem(&entityManager),
		spriteRenderSystem: systems.NewSpriteRenderSystem(&entityManager),
		textRenderSystem:   systems.NewTextRenderSystem(&entityManager),

		cleanupSystem: systems.NewCleanupSystem(&entityManager, 100, ScreenHeight),
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
	g.spriteRenderSystem.Draw(screen)
	g.textRenderSystem.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func constructScores() []*entities.Entity {
	scores := []*entities.Entity{}
	scoreCount := 8

	padding := 20
	xOffset := float64((ScreenWidth/scoreCount - padding) - padding)
	yPosition := float64(ScreenHeight - padding)
	scores = append(scores, entities.NewScore("Corn", components.Nutrient{Corn: 1.0}, xOffset, yPosition))
	scores = append(scores, entities.NewScore("Dairy", components.Nutrient{Dairy: 1.0}, xOffset*2, yPosition))
	scores = append(scores, entities.NewScore("Drink", components.Nutrient{Drink: 1.0}, xOffset*3, yPosition))
	scores = append(scores, entities.NewScore("Fish", components.Nutrient{Fish: 1.0}, xOffset*4, yPosition))
	scores = append(scores, entities.NewScore("Meat", components.Nutrient{Meat: 1.0}, xOffset*5, yPosition))
	scores = append(scores, entities.NewScore("Treat", components.Nutrient{Treat: 1.0}, xOffset*6, yPosition))
	scores = append(scores, entities.NewScore("Fruit", components.Nutrient{Fruit: 1.0}, xOffset*7, yPosition))
	scores = append(scores, entities.NewScore("Vegetable", components.Nutrient{Vegetable: 1.0}, xOffset*8, yPosition))
	return scores
}
