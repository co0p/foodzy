package foodzy

import (
	"github.com/co0p/foodzy/assets"
	"github.com/co0p/foodzy/components"
	"github.com/hajimehoshi/ebiten/v2"
)

const playerSize = 50

type Player struct {
	image      *ebiten.Image
	posX, posY float64
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	w, h := p.image.Size()
	sw, sh := screen.Size()

	tx := float64(sw/2+w/h-w/2) + p.posX
	ty := float64(sh-h-h/2) + p.posY
	op.GeoM.Translate(tx, ty)
	screen.DrawImage(p.image, op)
}

func (p *Player) Update(game *Game) {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.moveLeft()
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.moveRight()
	}
}

func (p *Player) moveLeft() {
	p.posX -= 5
}

func (p *Player) moveRight() {
	p.posX += 5
}

func NewPlayerElement() *components.Element {

	pos := components.Vector{
		X: float64(ScreenWidth / 2.0),
		Y: float64(ScreenHeight - playerSize*2.0),
	}
	playerElement := components.NewElement("player", true, pos, 0)

	spriteRenderer := components.NewSpriteRenderer(playerElement, assets.Plate)

	playerElement.AddComponent(components.NewLoggingComponent(playerElement))
	playerElement.AddComponent(spriteRenderer)

	return playerElement
}
