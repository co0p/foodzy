package foodzy

/*
this file is used to generate asset files to be loaded by the game
*/

// images
//go:generate file2byteslice -package=assets -input=assets/plate.png -output=assets/plate.go -var=Plate
//go:generate file2byteslice -package=assets -input=assets/background.png -output=assets/background.go -var=Background
//go:generate file2byteslice -package=assets -input=assets/knive.png -output=assets/knive.go -var=Knife
//go:generate file2byteslice -package=assets -input=assets/fork.png -output=assets/fork.go -var=Fork
//go:generate file2byteslice -package=assets -input=assets/beer.png -output=assets/beer.go -var=Beer
//go:generate file2byteslice -package=assets -input=assets/bread.png -output=assets/bread.go -var=Bread
//go:generate file2byteslice -package=assets -input=assets/carrot.png -output=assets/carrot.go -var=Carrot
//go:generate file2byteslice -package=assets -input=assets/cheese.png -output=assets/cheese.go -var=Cheese
//go:generate file2byteslice -package=assets -input=assets/fish.png -output=assets/fish.go -var=Fish
//go:generate file2byteslice -package=assets -input=assets/lemon.png -output=assets/lemon.go -var=Lemon
//go:generate file2byteslice -package=assets -input=assets/meat.png -output=assets/meat.go -var=Meat
//go:generate file2byteslice -package=assets -input=assets/strawberry.png -output=assets/strawberry.go -var=Strawberry
//go:generate file2byteslice -package=assets -input=assets/tomato.png -output=assets/tomato.go -var=Tomato
//go:generate file2byteslice -package=assets -input=assets/soundtrack.wav -output=assets/soundtrack.go -var=BackgroundMusic

