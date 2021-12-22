package foodzy

import (
	"github.com/co0p/foodzy/entity"
	"github.com/co0p/foodzy/system"
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
	entityManager *entity.Manager

	movementSystem   *system.MovementSystem
	foodSystem       *system.FoodSpawningSystem
	controllerSystem *system.ControllerSystem
	collisionSystem  *system.CollisionSystem
	scoreSystem      *system.ScoreSystem

	spriteRenderSystem *system.SpriteRenderSystem
	textRenderSystem   *system.TextRenderSystem

	soundSystem   *system.SoundSystem
	cleanupSystem *system.CleanupSystem
}

func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())

	entityManager := entity.Manager{}
	entityManager.AddEntity(entity.NewBackground())
	entityManager.AddEntity(entity.NewPlayer(ScreenWidth, ScreenHeight))
	entityManager.AddEntity(entity.NewFoodSpawner(SpawnFrequency))

	scores := entity.ConstructScores(ScreenWidth, ScreenHeight)
	for _, v := range scores {
		entityManager.AddEntity(v)
	}

	return &Game{
		entityManager:      &entityManager,
		movementSystem:     system.NewMovementSystem(&entityManager),
		controllerSystem:   system.NewControllerSystem(&entityManager, ScreenWidth, ScreenHeight),
		soundSystem:        system.NewSoundSystem(),
		foodSystem:         system.NewFoodSpawningSystem(&entityManager, ScreenWidth),
		collisionSystem:    system.NewCollisionSystem(&entityManager),
		scoreSystem:        system.NewScoreSystem(&entityManager),
		spriteRenderSystem: system.NewSpriteRenderSystem(&entityManager),
		textRenderSystem:   system.NewTextRenderSystem(&entityManager),

		cleanupSystem: system.NewCleanupSystem(&entityManager, 100, ScreenHeight),
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
