package entity

import (
	"github.com/co0p/foodzy/component"
	"image/color"
)

func NewScore(text string, nutrient component.Nutrient, posX float64, posY float64) *Entity {
	e := NewEntity("score", true)

	e.AddComponent(&component.Transform{X: posX, Y: posY, Scale: 1})
	e.AddComponent(&component.Text{Value: text, Color: color.White})
	e.AddComponent(&nutrient)

	return e
}

func ConstructScores(ScreenWidth int, ScreenHeight int) []*Entity {
	scoreCount := 8
	padding := 20
	xOffset := float64((ScreenWidth/scoreCount - padding) - padding)
	yPosition := float64(ScreenHeight - padding)

	var scores []*Entity
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
