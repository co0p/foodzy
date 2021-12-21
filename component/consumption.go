package component

var ConsumptionDefaultRate = -0.001

type Consumption struct {
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

const ConsumptionType ComponentType = "Consumption"

func (c *Consumption) Type() ComponentType {
	return ConsumptionType
}
