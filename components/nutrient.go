package components

type Nutrient struct {
	Water         float64
	Carbohydrates float64
	Protein       float64
	Fat           float64
	Vitamins      float64
	Minerals      float64
}

func (p Nutrient) ID() string {
	return ""
}
