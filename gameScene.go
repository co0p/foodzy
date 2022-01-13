package foodzy

import (
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/co0p/foodzy/internal/scene"
	"github.com/co0p/foodzy/internal/sound"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const GameSceneName string = "game"

type GameScene struct {
	scene.GameScene
	gameoverAction ActionType
	pauseAction    ActionType
	entityManager  *ecs.EntityManager
	soundManager   *sound.SoundManager
}

func NewGameScene(soundManager *sound.SoundManager, gameoverAction ActionType, pauseAction ActionType) *GameScene {
	entityManager := ecs.EntityManager{}
	entityManager.AddEntity(NewBackground())
	entityManager.AddEntity(NewPlayer())
	entityManager.AddEntity(NewFoodSpawner(40))
	entityManager.AddEntity(NewHealth())

	g := &GameScene{
		entityManager:  &entityManager,
		soundManager:   soundManager,
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
		NewDebugRendererSystem(&entityManager),
		NewCleanupSystem(&entityManager, 100),
	)

	return g
}

func (g *GameScene) Name() string {
	return GameSceneName
}

func (g *GameScene) Init() {
	g.Stop()
	g.entityManager.Clear()

	g.soundManager.Volume(SoundBackground, 1)

	g.entityManager.AddEntity(NewBackground())
	g.entityManager.AddEntity(NewPlayer())
	g.entityManager.AddEntity(NewFoodSpawner(40))
	g.entityManager.AddEntity(NewHealth())
}

func (g *GameScene) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.pauseAction(nil)
	}

	return g.GameScene.Update()
}
