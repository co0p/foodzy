package components

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"reflect"
)

type Vector struct {
	X, Y float64
}

// Component describes functionality elements will use
type Component interface {
	OnUpdate() error
	OnDraw(screen *ebiten.Image) error
}

// Element is the basic build block of the Component based architecture holding the components
type Element struct {
	Position   Vector
	Rotation   float64
	Active     bool
	components []Component
	Id         string
}

func NewElement(name string, active bool, pos Vector, rotation float64) *Element {
	return &Element{
		Id:       name,
		Active:   active,
		Position: pos,
		Rotation: rotation,
	}
}

// AddComponent adds a Component to the Element; panics if there is a Component of same type already
func (elem *Element) AddComponent(new Component) {
	newType := reflect.TypeOf(new)
	for _, existing := range elem.components {
		if newType == reflect.TypeOf(existing) {
			panic(fmt.Sprintf("attempted to add a Component of type %v that already has been added", new))
		}
	}

	elem.components = append(elem.components, new)
}

// GetComponent gets a Component from the Element; panics if Component does not exist
func (elem *Element) GetComponent(withType Component) Component {
	askingType := reflect.TypeOf(withType)
	for _, existing := range elem.components {
		if askingType == reflect.TypeOf(existing) {
			return existing
		}
	}
	panic(fmt.Sprintf("attempted to get omponent of type %v that does not exist", withType))
}

func (elem *Element) OnUpdate() {
	for _, element := range elem.components {
		if err := element.OnUpdate(); err != nil {
			panic("should not happen")
		}
	}
}

func (elem *Element) OnDraw(screen *ebiten.Image) {
	for _, element := range elem.components {
		if err := element.OnDraw(screen); err != nil {
			panic("should not happen")
		}
	}
}
