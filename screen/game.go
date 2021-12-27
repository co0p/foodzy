package screen

import (
	"github.com/co0p/foodzy/entity"
	"github.com/co0p/foodzy/system"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const GameScreenName string = "game"

type GameScreen struct {
	ScreenWidth   int
	ScreenHeight  int
	entityManager *entity.Manager
	systems       []system.System
	actions       map[Action]func()
	initialized   bool
}

func NewGameScreen(ScreenWidth int, ScreenHeight int) *GameScreen {
	return &GameScreen{
		ScreenWidth:   ScreenWidth,
		ScreenHeight:  ScreenHeight,
		entityManager: &entity.Manager{},
		systems:       []system.System{},
		actions:       make(map[Action]func()),
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
		system.NewFoodSpawningSystem(g.entityManager, g.ScreenWidth),
		system.NewCollisionSystem(g.entityManager),
		system.NewScoreSystem(g.entityManager),
		system.NewSpriteRenderSystem(g.entityManager),
		system.NewTextRenderSystem(g.entityManager),
		system.NewCleanupSystem(g.entityManager, 100, g.ScreenHeight),
	)

	g.initialized = true
}

func (g *GameScreen) Exit() {
	log.Printf("[screen:%s] exit\n", g.Name())
}

func (g *GameScreen) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		g.actions[ActionActivateStartScreen]()
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

func (g *GameScreen) AddAction(action Action, callback func()) {
	g.actions[action] = callback
}
