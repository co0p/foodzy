package components_test

import (
	"github.com/co0p/foodzy/components"
	"github.com/hajimehoshi/ebiten/v2"
	"testing"
)

func TestElement_AddComponent(t *testing.T) {

	myElement := components.Element{}
	myElement.AddComponent(&dummyComponent{})
	assertPanic(t, func() {
		myElement.AddComponent(&dummyComponent{})
	})
}

func TestElement_GetComponent(t *testing.T) {

	myElement := components.Element{}
	myElement.AddComponent(&dummyComponent{})

	component := myElement.GetComponent(&dummyComponent{})
	if component == nil {
		t.Errorf("expected component not to be nil, got nil")
	}

}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}

type dummyComponent struct{}

func (d dummyComponent) OnUpdate() error {
	return nil
}

func (d dummyComponent) OnDraw(screen *ebiten.Image) error {
	return nil
}
