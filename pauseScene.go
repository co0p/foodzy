package foodzy

import (
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/co0p/foodzy/internal/scene"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

const PauseSceneName = "pause"

type PauseScene struct {
	scene.GameScene

	resumeAction ActionType
}

func NewPauseScene(resumeAction ActionType) *PauseScene {

	transform := component.Transform{X: 100, Y: 100, Z: 1, Scale: 1}
	text := component.Text{Value: "PAUSE", Color: colornames.Beige, Font: &FontHuge}

	pauseText := ecs.NewEntity("pause", true)
	pauseText.AddComponents(&transform, &text)

	entityManager := ecs.EntityManager{}
	entityManager.AddEntity(NewBackground())
	entityManager.AddEntity(pauseText)

	s := &PauseScene{resumeAction: resumeAction}
	s.Systems = append(s.Systems,
		NewSpriteRenderSystem(&entityManager),
		NewTextRenderSystem(&entityManager),
	)

	return s
}

func (s *PauseScene) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		s.resumeAction(nil)
	}
	return s.GameScene.Update()
}

func (s *PauseScene) Name() string {
	return PauseSceneName
}
