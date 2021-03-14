package foodzy

import (
	"github.com/co0p/foodzy/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type Player struct {
	image *ebiten.Image
	posX, posY  float64
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	w, h := p.image.Size()
	sw, sh := screen.Size()

	tx := float64(sw/2 + w/h - w/2) + p.posX
	ty := float64(sh - h - h/2) + p.posY
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

func NewPlayer() *Player {
	img, _ := LoadImage(assets.Plate)

	player := &Player{
		image: ebiten.NewImageFromImage(img),
	}

	log.Printf("created player")
	return player
}
