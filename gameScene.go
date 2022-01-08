package foodzy

import (
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/co0p/foodzy/internal/scene"
	"github.com/co0p/foodzy/internal/sound"
	"github.com/hajimehoshi/ebiten/v2"
)

const GameSceneName string = "game"

type GameScene struct {
	scene.GameScene
	gameoverAction ActionType
	pauseAction    ActionType
	entityManager  *ecs.EntityManager
}

func NewGameScene(soundManager *sound.SoundManager, gameoverAction ActionType, pauseAction ActionType) *GameScene {
	entityManager := ecs.EntityManager{}
	entityManager.AddEntity(NewBackground())
	entityManager.AddEntity(NewPlayer())
	entityManager.AddEntity(NewFoodSpawner(40))
	entityManager.AddEntity(NewHealth())

	g := &GameScene{
		entityManager:  &entityManager,
		gameoverAction: gameoverAction,
		pauseAction:    pauseAction,
	}

	g.Systems = append(g.Systems,
		NewHealthSystem(&entityManager, gameoverAction),
		NewMovementSystem(&entityManager),
		NewControllerSystem(&entityManager),
		NewFoodSpawningSystem(&entityManager),
		NewCollisionSystem(&entityManager),
		NewSpriteRenderSystem(&entityManager),
		NewTextRenderSystem(&entityManager),
		NewSoundSystem(&entityManager, soundManager),
		NewCleanupSystem(&entityManager, 100),
	)

	return g
}

func (g *GameScene) Name() string {
	return GameSceneName
}

func (g *GameScene) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		g.pauseAction(nil)
	}

	return g.GameScene.Update()
}
