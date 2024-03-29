package foodzy

import (
	"github.com/co0p/foodzy/asset"
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
)

var (
	FontSmall  font.Face
	FontMedium font.Face
	FontBig    font.Face
	FontHuge   font.Face
)

var (
	PrimaryColor   = color.RGBA{R: 0, G: 20, B: 27, A: 240}
	SecondaryColor = color.RGBA{R: 0, G: 20, B: 27, A: 200}
)

func init() {
	FontSmall = loadFont(asset.StopBullyingFont, 14)
	FontMedium = loadFont(asset.StopBullyingFont, 24)
	FontBig = loadFont(asset.StopBullyingFont, 46)
	FontHuge = loadFont(asset.StopBullyingFont, 96)
}

type TextRenderSystem struct {
	manager *ecs.EntityManager
}

func NewTextRenderSystem(manager *ecs.EntityManager) *TextRenderSystem {
	return &TextRenderSystem{manager: manager}
}

func (s *TextRenderSystem) Draw(screen *ebiten.Image) {
	candidates := s.manager.QueryByComponents(component.TextType, component.TransformType)

	for _, e := range candidates {

		if !e.Active {
			continue
		}

		txt := e.GetComponent(component.TextType).(*component.Text)
		pos := e.GetComponent(component.TransformType).(*component.Transform)

		// we have to adjust the rendering due to the dot-positioning
		_, h := txt.Dimensions()
		text.Draw(screen, txt.Value, *txt.Font, int(pos.X), int(pos.Y+h), txt.Color)
	}
}

func (s *TextRenderSystem) Update() error {
	return nil
}

func loadFont(ttfFont []byte, size float64) font.Face {

	tt, err := opentype.Parse(ttfFont)
	if err != nil {
		log.Fatal(err)
	}

	f, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	return f
}
