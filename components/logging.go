package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

type LoggingComponent struct {
	Container *Element
}

func (l *LoggingComponent) OnUpdate() error {
	log.Printf("[%s] OnUpdate\n", l.Container.Id)
	return nil
}

func (l *LoggingComponent) OnDraw(screen *ebiten.Image) error {
	log.Printf("[%s] OnDraw - [%v,%v]\n", l.Container.Id, l.Container.Position.X, l.Container.Position.Y)
	return nil
}

func NewLoggingComponent(container *Element) *LoggingComponent {
	return &LoggingComponent{
		Container: container,
	}
}
