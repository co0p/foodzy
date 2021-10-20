package components

const (
	FoodFruit = iota
	FoodVegetable
	FoodDrink
	FoodMeat
	FoodFish
	FoodCorn
	FoodDairy
	FoodTreat
)

type FoodType int

type FoodSpawner struct {
	CoolDown int
	Rate     int
	Variance float64
	Velocity struct {
		X float64
		Y float64
	}
	Types []FoodType
}

func (f FoodSpawner) ID() string {
	return ""
}
