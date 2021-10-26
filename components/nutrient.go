package components

type Nutrient struct {
	Corn      float64
	Dairy     float64
	Drink     float64
	Fish      float64
	Meat      float64
	Treat     float64
	Fruit     float64
	Vegetable float64
	KCal      float64
}

func (n Nutrient) ID() string {
	panic("")
}

func (n *Nutrient) Add(other *Nutrient) {
	n.Corn = Floor(n.Corn + other.Corn)
	n.Dairy = Floor(n.Dairy + other.Dairy)
	n.Drink = Floor(n.Drink + other.Drink)
	n.Fish = Floor(n.Fish + other.Fish)
	n.Meat = Floor(n.Meat + other.Meat)
	n.Treat = Floor(n.Treat + other.Treat)
	n.Fruit = Floor(n.Fruit + other.Fruit)
	n.Vegetable = Floor(n.Vegetable + other.Vegetable)
	n.KCal = Floor(n.KCal + other.KCal)
}

func (n *Nutrient) MatMul(other *Nutrient) float64 {
	return n.Corn*other.Corn +
		n.Dairy*other.Dairy +
		n.Drink*other.Drink +
		n.Fish*other.Fish +
		n.Meat*other.Meat +
		n.Treat*other.Treat +
		n.Fruit*other.Fruit +
		n.Vegetable*other.Vegetable +
		n.KCal*other.KCal
}

func Floor(x float64) float64 {
	if x < 0 {
		return 0
	}
	return x
}
