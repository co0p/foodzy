package component

type FoodSpawner struct {
	CoolDown int
	Rate     int
	Variance float64
	Velocity struct {
		X float64
		Y float64
	}
}

const FoodSpawnerType ComponentType = "FoodSpawner"

func (f *FoodSpawner) Type() ComponentType {
	return FoodSpawnerType
}
