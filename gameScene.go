package foodzy

import (
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/co0p/foodzy/internal/scene"
	"github.com/co0p/foodzy/internal/sound"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const GameScreenName string = "game"

type GameScreen struct {
	entityManager *ecs.Manager
	soundManager  *sound.SoundManager
	systems       []ecs.System
	actions       map[scene.Action]func()
	initialized   bool
}

func NewGameScreen(soundManager *sound.SoundManager) *GameScreen {
	return &GameScreen{
		entityManager: &ecs.Manager{},
		soundManager:  soundManager,
		systems:       []ecs.System{},
		actions:       make(map[scene.Action]func()),
		initialized:   false,
	}
}

func (g *GameScreen) Name() string {
	return GameScreenName
}

func (g *GameScreen) Init() {
	if g.initialized {
		log.Printf("[screen:%s] resuming\n", g.Name())
		return
	}

	log.Printf("[screen:%s] initializing\n", g.Name())

	g.entityManager.AddEntity(NewBackground())
	g.entityManager.AddEntity(NewPlayer(ScreenWidth, ScreenHeight))
	g.entityManager.AddEntity(NewFoodSpawner(40))

	scores := ConstructScores(ScreenWidth, ScreenHeight)
	for _, v := range scores {
		g.entityManager.AddEntity(v)
	}

	g.systems = append(g.systems,
		NewMovementSystem(g.entityManager),
		NewControllerSystem(g.entityManager),
		NewFoodSpawningSystem(g.entityManager),
		NewCollisionSystem(g.entityManager),
		NewSpriteRenderSystem(g.entityManager),
		NewTextRenderSystem(g.entityManager),
		NewSoundSystem(g.entityManager, g.soundManager),
		NewCleanupSystem(g.entityManager, 100),
	)

	g.initialized = true
}

func (g *GameScreen) Exit() {
	log.Printf("[screen:%s] exit\n", g.Name())
}

func (g *GameScreen) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		g.actions[scene.ActionActivateStartScreen]()
	}

	for _, s := range g.systems {
		if err := s.Update(); err != nil {
			return err
		}
	}

	return nil
}

func (g *GameScreen) Draw(screen *ebiten.Image) {
	for _, s := range g.systems {
		s.Draw(screen)
	}
}

func (g *GameScreen) AddAction(action scene.Action, callback func()) {
	g.actions[action] = callback
}
