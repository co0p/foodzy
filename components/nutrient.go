package components

type Nutrient struct {
	Water         float64
	Carbohydrates float64
	Protein       float64
	Fat           float64
	Vitamins      float64
	Minerals      float64
}

func (n Nutrient) ID() string {
	return ""
}

func (n *Nutrient) Add(other Nutrient) {
	n.Protein = Floor(n.Protein + other.Protein)
	n.Fat = Floor(n.Fat + other.Fat)
	n.Water = Floor(n.Water + other.Water)
	n.Minerals = Floor(n.Minerals + other.Minerals)
	n.Vitamins = Floor(n.Vitamins + other.Vitamins)
	n.Carbohydrates = Floor(n.Carbohydrates + other.Carbohydrates)
}

func (n *Nutrient) MatMul(other *Nutrient) float64 {
	return n.Fat*other.Fat +
		n.Water*other.Water +
		n.Carbohydrates*other.Carbohydrates +
		n.Minerals*other.Minerals +
		n.Protein*other.Protein +
		n.Vitamins*other.Vitamins
}

func Floor(x float64) float64 {
	if x < 0 {
		return 0
	}
	return x
}
