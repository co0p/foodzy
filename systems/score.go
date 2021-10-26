package systems

import (
	"fmt"
	"github.com/co0p/foodzy/components"
	"github.com/co0p/foodzy/entities"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
	"image/color"
)

type ScoreSystem struct {
	manager      *entities.Manager
	ScreenHeight int
	ScreenWidth  int
}

func NewScoreSystem(manager *entities.Manager) *ScoreSystem {
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

	currentNutrients := player.GetComponent(&components.Nutrient{}).(*components.Nutrient)
	consumption := player.GetComponent(&components.Consumption{}).(*components.Consumption)
	change := components.Nutrient{
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
	scores := s.manager.QueryByComponents(&components.Nutrient{}, &components.Text{})
	for _, v := range scores {
		scoreNutrient := v.GetComponent(&components.Nutrient{}).(*components.Nutrient)
		newScoreValue := currentNutrients.MatMul(scoreNutrient)

		text := v.GetComponent(&components.Text{}).(*components.Text)
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
