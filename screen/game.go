package screen

import (
	"github.com/co0p/foodzy/entity"
	"github.com/co0p/foodzy/system"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const GameScreenName Name = "game"

type GameScreen struct {
	ScreenWidth   int
	ScreenHeight  int
	entityManager *entity.Manager
	systems       []system.System
}

func NewGameScreen(ScreenWidth int, ScreenHeight int) *GameScreen {
	return &GameScreen{
		ScreenWidth:   ScreenWidth,
		ScreenHeight:  ScreenHeight,
		entityManager: &entity.Manager{},
		systems:       []system.System{},
	}
}

func (g *GameScreen) Name() Name {
	return GameScreenName
}

func (g *GameScreen) Init() {
	log.Printf("[screen:%s] initializing\n", g.Name())

	g.entityManager.AddEntity(entity.NewBackground())
	g.entityManager.AddEntity(entity.NewPlayer(g.ScreenWidth, g.ScreenHeight))
	g.entityManager.AddEntity(entity.NewFoodSpawner(40))

	scores := entity.ConstructScores(g.ScreenWidth, g.ScreenHeight)
	for _, v := range scores {
		g.entityManager.AddEntity(v)
	}

	g.systems = append(g.systems,
		system.NewMovementSystem(g.entityManager),
		system.NewControllerSystem(g.entityManager, g.ScreenWidth, g.ScreenHeight),
		system.NewSoundSystem(),
		system.NewFoodSpawningSystem(g.entityManager, g.ScreenWidth),
		system.NewCollisionSystem(g.entityManager),
		system.NewScoreSystem(g.entityManager),
		system.NewSpriteRenderSystem(g.entityManager),
		system.NewTextRenderSystem(g.entityManager),
		system.NewCleanupSystem(g.entityManager, 100, g.ScreenHeight),
	)
}

func (g *GameScreen) Exit() {
	log.Printf("[screen:%s] exit\n", g.Name())
}

func (g GameScreen) Update() error {

	for _, s := range g.systems {
		if err := s.Update(); err != nil {
			return err
		}
	}

	return nil
}

func (g GameScreen) Draw(screen *ebiten.Image) {
	for _, s := range g.systems {
		s.Draw(screen)
	}
}
