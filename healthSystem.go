package foodzy

import (
	"fmt"
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type HealthSystem struct {
	manager        *ecs.EntityManager
	gameOverAction func()
}

func NewHealthSystem(manager *ecs.EntityManager, gameOverAction func()) *HealthSystem {
	return &HealthSystem{manager: manager, gameOverAction: gameOverAction}
}

func (s *HealthSystem) Draw(screen *ebiten.Image) {}

func (s *HealthSystem) Update() error {

	// update player consumption
	player := s.manager.QueryFirstByTag("player")
	health := player.GetComponent(component.HealthType).(*component.Health)
	health.Consume()

	// update the score
	score := s.manager.QueryFirstByTag("score")
	text := score.GetComponent(component.TextType).(*component.Text)
	text.Value = fmt.Sprintf("%3d%%", health.CurrentHealth())

	if health.CurrentHealth() == 0 {
		s.gameOverAction()
	}

	return nil
}
