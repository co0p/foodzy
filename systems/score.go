package systems

import (
	"fmt"
	"github.com/co0p/foodzy/components"
	"github.com/co0p/foodzy/entities"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
	"image/color"
)

const NutrientsConsumptionRate = -0.05

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
	change := components.Nutrient{
		Water:         NutrientsConsumptionRate,
		Carbohydrates: NutrientsConsumptionRate,
		Protein:       NutrientsConsumptionRate,
		Fat:           NutrientsConsumptionRate,
		Vitamins:      NutrientsConsumptionRate,
		Minerals:      NutrientsConsumptionRate,
	}
	currentNutrients.Add(change)

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

	if level >= 10 {
		c = colornames.Red
	}

	if level >= 20 {
		c = colornames.Darkorange
	}

	if level >= 30 {
		c = colornames.Blue
	}

	if level >= 40 {
		c = colornames.Green
	}

	if level >= 50 {
		c = colornames.Darkgreen
	}

	return c
}
