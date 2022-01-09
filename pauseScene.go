package foodzy

import (
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/co0p/foodzy/internal/scene"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const PauseSceneName = "pause"

type PauseScene struct {
	scene.GameScene

	resumeAction ActionType
}

func NewPauseScene(resumeAction ActionType) *PauseScene {
	txt := component.Text{Value: "PAUSE", Color: PrimaryColor, Font: &FontHuge}

	posX, posY := txt.RelativeCenter(ScreenWidth, ScreenHeight)
	transform := component.Transform{X: posX, Y: posY, Z: 1, Scale: 1}
	pauseText := ecs.NewEntity("pause", true)
	pauseText.AddComponents(&transform, &txt)

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

	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		s.resumeAction(nil)
		return nil
	}

	return s.GameScene.Update()
}

func (s *PauseScene) Name() string {
	return PauseSceneName
}
