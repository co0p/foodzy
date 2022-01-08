package main

import (
	"github.com/co0p/foodzy"
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

type TextDemo struct {
	textRenderer *foodzy.TextRenderSystem
}

func (g *TextDemo) Update() error {
	return g.textRenderer.Update()
}

func (g *TextDemo) Draw(screen *ebiten.Image) {
	g.textRenderer.Draw(screen)
}

func (g *TextDemo) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 400, 400
}

func main() {

	// E - Entity
	label := ecs.NewEntity("label", true)
	entityManager := ecs.EntityManager{}
	entityManager.AddEntity(label)

	// C - components
	text := component.Text{Color: colornames.Cyan, Value: "Hi Gophers", Font: &foodzy.FontMedium}
	transform := component.Transform{X: 100, Y: 100, Z: 1, Scale: 1}
	label.AddComponents(&transform, &text)

	// S - System
	textRenderer := foodzy.NewTextRenderSystem(&entityManager)

	game := TextDemo{textRenderer: textRenderer}

	if err := ebiten.RunGame(&game); err != nil {
		panic(err)
	}
}
