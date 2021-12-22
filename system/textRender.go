package system

import (
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/entity"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"log"
)

type TextRenderSystem struct {
	manager *entity.Manager
	font    font.Face
}

func NewTextRenderSystem(manager *entity.Manager) *TextRenderSystem {
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
	candidates := s.manager.QueryByComponents(component.TextType, component.TransformType)

	for _, e := range candidates {

		if !e.Active {
			continue
		}

		txt := e.GetComponent(component.TextType).(*component.Text)
		pos := e.GetComponent(component.TransformType).(*component.Transform)
		text.Draw(screen, txt.Value, s.font, int(pos.X), int(pos.Y), txt.Color)
	}
}

func (s *TextRenderSystem) Update() error {
	return nil
}
