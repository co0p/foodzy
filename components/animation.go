package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"math"
)

type Animation struct {
	frames       []*ebiten.Image
	currentFrame *ebiten.Image
	count        float64
	currentIndex int
	Speed        float64
}

func (a *Animation) GetCurrentFrame() *ebiten.Image {
	return a.currentFrame
}

func (a Animation) ID() string {
	return ""
}

func (a *Animation) Step() {
	a.count += a.Speed
	a.currentIndex = int(math.Floor(a.count))

	if a.currentIndex >= len(a.frames) {
		a.count = 0
		a.currentIndex = 0
	}

	a.currentFrame = a.frames[a.currentIndex]
}

func NewAnimation(speed float64, frames []*ebiten.Image) *Animation {

	if len(frames) == 0 {
		log.Fatal("animation component cannot have empty frames")
	}

	return &Animation{
		frames:       frames,
		currentFrame: frames[0],
		Speed:        speed,
	}
}
