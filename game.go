package foodzy

import (
	"github.com/co0p/foodzy/entities"
	"github.com/co0p/foodzy/systems"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  int    = 800
	ScreenHeight int    = 600
	GameName     string = "Foodzy"
)

type Game struct {
	entityManager *entities.Manager

	movementSystem   *systems.MovementSystem
	controllerSystem *systems.ControllerSystem
	renderSystem     *systems.RenderSystem

	soundSystem *systems.SoundSystem
}

func NewGame() *Game {

	entityManager := entities.Manager{}
	entityManager.AddEntity(entities.NewBackground())
	entityManager.AddEntity(entities.NewPlayer(ScreenWidth, ScreenHeight))

	return &Game{
		entityManager:    &entityManager,
		movementSystem:   systems.NewMovementSystem(&entityManager),
		controllerSystem: systems.NewControllerSystem(&entityManager, ScreenWidth, ScreenHeight),
		renderSystem:     systems.NewRenderSystem(&entityManager),
		soundSystem:      systems.NewSoundSystem(),
	}
}

func (g *Game) Update() error {
	err := g.controllerSystem.Update()
	err = g.movementSystem.Update()
	return err
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderSystem.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
