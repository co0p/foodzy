package systems

import (
	"fmt"
	"github.com/co0p/foodzy/components"
	"github.com/co0p/foodzy/entities"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
)

type ScoreSystem struct {
	manager      *entities.Manager
	font         font.Face
	ScreenHeight int
	ScreenWidth  int
}

func NewScoreSystem(manager *entities.Manager, width int, height int) *ScoreSystem {

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	font, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	return &ScoreSystem{manager: manager, font: font, ScreenWidth: width, ScreenHeight: height}
}

func (s *ScoreSystem) Draw(screen *ebiten.Image) {
	e := s.manager.QueryByTag("player")
	if len(e) != 1 {
		panic(fmt.Sprintf("expected one entity with tag 'player', got %v", len(e)))
	}
	player := e[0]
	currentNutrients := player.GetComponent(&components.Nutrient{}).(*components.Nutrient)

	candidates := s.manager.QueryByComponents(&components.Text{}, &components.Nutrient{})
	padding := 20
	xOffset := s.ScreenWidth/len(candidates) - padding
	yPosition := s.ScreenHeight - padding
	for idx, entity := range candidates {
		label := entity.GetComponent(&components.Text{}).(*components.Text)
		nutrients := entity.GetComponent(&components.Nutrient{}).(*components.Nutrient)

		msg := fmt.Sprintf("%s: %0.f", label.Text, getNutrientScore(currentNutrients, nutrients))
		text.Draw(screen, msg, s.font, xOffset*idx+padding, yPosition, color.White)
	}
}

func (s *ScoreSystem) Update() error {
	e := s.manager.QueryByTag("player")
	if len(e) != 1 {
		panic(fmt.Sprintf("expected one entity with tag 'player', got %v", len(e)))
	}
	player := e[0]
	currentNutrients := player.GetComponent(&components.Nutrient{}).(*components.Nutrient)
	currentNutrients.Fat = Floor(currentNutrients.Fat - 0.01)
	currentNutrients.Water = Floor(currentNutrients.Water - 0.01)
	currentNutrients.Carbohydrates = Floor(currentNutrients.Carbohydrates - 0.01)
	currentNutrients.Minerals = Floor(currentNutrients.Minerals - 0.01)
	currentNutrients.Protein = Floor(currentNutrients.Protein - 0.01)
	currentNutrients.Vitamins = Floor(currentNutrients.Vitamins - 0.01)

	return nil
}

func getNutrientScore(playerNutrients *components.Nutrient, componentNutrients *components.Nutrient) float64 {
	return playerNutrients.Fat*componentNutrients.Fat +
		playerNutrients.Water*componentNutrients.Water +
		playerNutrients.Carbohydrates*componentNutrients.Carbohydrates +
		playerNutrients.Minerals*componentNutrients.Minerals +
		playerNutrients.Protein*componentNutrients.Protein +
		playerNutrients.Vitamins*componentNutrients.Vitamins
}

func Floor(x float64) float64 {
	if x < 0 {
		return 0
	}
	return x
}
