package systems

import (
	"github.com/co0p/foodzy/components"
	"github.com/co0p/foodzy/entities"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"log"
)

type TextRenderSystem struct {
	manager *entities.Manager
	font    font.Face
}

func NewTextRenderSystem(manager *entities.Manager) *TextRenderSystem {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	f, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    20,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	return &TextRenderSystem{manager: manager, font: f}
}

func (s *TextRenderSystem) Draw(screen *ebiten.Image) {
	candidates := s.manager.QueryByComponents(&components.Text{}, &components.Position{})

	for _, e := range candidates {

		if !e.Active {
			continue
		}

		pos := e.GetComponent(&components.Position{}).(*components.Position)
		txt := e.GetComponent(&components.Text{}).(*components.Text)
		text.Draw(screen, txt.Value, s.font, int(pos.X), int(pos.Y), txt.Color)
	}
}

func (s *TextRenderSystem) Update() error {
	return nil
}
