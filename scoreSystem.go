package foodzy

import (
	"fmt"
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/internal/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
	"image/color"
)

type ScoreSystem struct {
	manager      *ecs.Manager
	ScreenHeight int
	ScreenWidth  int
}

func NewScoreSystem(manager *ecs.Manager) *ScoreSystem {
	return &ScoreSystem{manager: manager}
}

func (s *ScoreSystem) Draw(screen *ebiten.Image) {}

func (s *ScoreSystem) Update() error {

	// update player stats
	e := s.manager.QueryByTag("player")
	if len(e) != 1 {
		panic(fmt.Sprintf("expected one entity with tag 'player', got %v", len(e)))
	}
	player := e[0]

	currentNutrients := player.GetComponent(component.NutrientType).(*component.Nutrient)
	consumption := player.GetComponent(component.ConsumptionType).(*component.Consumption)
	change := component.Nutrient{
		Corn:      consumption.Corn,
		Dairy:     consumption.Dairy,
		Drink:     consumption.Drink,
		Fish:      consumption.Fish,
		Meat:      consumption.Meat,
		Treat:     consumption.Treat,
		Fruit:     consumption.Fruit,
		Vegetable: consumption.Vegetable,
		KCal:      consumption.KCal,
	}
	currentNutrients.Add(&change)

	// update colors of score
	scores := s.manager.QueryByComponents(component.NutrientType, component.TextType)
	for _, v := range scores {
		scoreNutrient := v.GetComponent(component.NutrientType).(*component.Nutrient)
		newScoreValue := currentNutrients.MatMul(scoreNutrient)

		text := v.GetComponent(component.TextType).(*component.Text)
		text.Color = deriveFontColor(newScoreValue)
	}

	return nil
}

func deriveFontColor(level float64) color.Color {

	c := colornames.Darkred

	if level >= 1 {
		c = colornames.Red
	}

	if level >= 2 {
		c = colornames.Darkorange
	}

	if level >= 3 {
		c = colornames.Blue
	}

	if level >= 4 {
		c = colornames.Green
	}

	if level >= 5 {
		c = colornames.Darkgreen
	}

	return c
}
