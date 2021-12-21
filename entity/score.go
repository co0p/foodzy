package entity

import (
	"github.com/co0p/foodzy/component"
	"image/color"
)

func NewScore(text string, nutrient component.Nutrient, posX float64, posY float64) *entity {
	bg := NewEntity("score", true)

	bg.AddComponent(&component.Position{X: posX, Y: posY})
	bg.AddComponent(&component.Text{Value: text, Color: color.White})
	bg.AddComponent(&nutrient)

	return bg
}

func ConstructScores(ScreenWidth int, ScreenHeight int) []*entity {
	scoreCount := 8
	padding := 20
	xOffset := float64((ScreenWidth/scoreCount - padding) - padding)
	yPosition := float64(ScreenHeight - padding)

	var scores []*entity
	scores = append(scores, NewScore("Corn", component.Nutrient{Corn: 1.0}, xOffset, yPosition))
	scores = append(scores, NewScore("Dairy", component.Nutrient{Dairy: 1.0}, xOffset*2, yPosition))
	scores = append(scores, NewScore("Drink", component.Nutrient{Drink: 1.0}, xOffset*3, yPosition))
	scores = append(scores, NewScore("Fish", component.Nutrient{Fish: 1.0}, xOffset*4, yPosition))
	scores = append(scores, NewScore("Meat", component.Nutrient{Meat: 1.0}, xOffset*5, yPosition))
	scores = append(scores, NewScore("Treat", component.Nutrient{Treat: 1.0}, xOffset*6, yPosition))
	scores = append(scores, NewScore("Fruit", component.Nutrient{Fruit: 1.0}, xOffset*7, yPosition))
	scores = append(scores, NewScore("Vegetable", component.Nutrient{Vegetable: 1.0}, xOffset*8, yPosition))
	return scores
}
