package foodzy

/*
this file is used to generate asset files to be loaded by the game
*/

// images
//go:generate file2byteslice -package=assets -input=assets/background.png -output=assets/background.go -var=Background
//go:generate file2byteslice -package=assets -input=assets/corn_baguette.png -output=assets/corn_baguette.go -var=Corn_baguette
//go:generate file2byteslice -package=assets -input=assets/corn_bread.png -output=assets/corn_bread.go -var=Corn_bread
//go:generate file2byteslice -package=assets -input=assets/corn_rice.png -output=assets/corn_rice.go -var=Corn_rice
//go:generate file2byteslice -package=assets -input=assets/dairy_cheese.png -output=assets/dairy_cheese.go -var=Dairy_cheese
//go:generate file2byteslice -package=assets -input=assets/dairy_milk.png -output=assets/dairy_milk.go -var=Dairy_milk
//go:generate file2byteslice -package=assets -input=assets/drink_beer.png -output=assets/drink_beer.go -var=Drink_beer
//go:generate file2byteslice -package=assets -input=assets/drink_coffee.png -output=assets/drink_coffee.go -var=Drink_coffee
//go:generate file2byteslice -package=assets -input=assets/drink_juice.png -output=assets/drink_juice.go -var=Drink_juice
//go:generate file2byteslice -package=assets -input=assets/drink_tea.png -output=assets/drink_tea.go -var=Drink_tea
//go:generate file2byteslice -package=assets -input=assets/drink_water.png -output=assets/drink_water.go -var=Drink_water
//go:generate file2byteslice -package=assets -input=assets/fish_crab.png -output=assets/fish_crab.go -var=Fish_crab
//go:generate file2byteslice -package=assets -input=assets/fish_sushi.png -output=assets/fish_sushi.go -var=Fish_sushi
//go:generate file2byteslice -package=assets -input=assets/fruit_apple.png -output=assets/fruit_apple.go -var=Fruit_apple
//go:generate file2byteslice -package=assets -input=assets/fruit_banana.png -output=assets/fruit_banana.go -var=Fruit_banana
//go:generate file2byteslice -package=assets -input=assets/fruit_grapes.png -output=assets/fruit_grapes.go -var=Fruit_grapes
//go:generate file2byteslice -package=assets -input=assets/fruit_orange.png -output=assets/fruit_orange.go -var=Fruit_orange
//go:generate file2byteslice -package=assets -input=assets/fruit_strawberry.png -output=assets/fruit_strawberry.go -var=Fruit_strawberry
//go:generate file2byteslice -package=assets -input=assets/meat_steak.png -output=assets/meat_steak.go -var=Meat_steak
//go:generate file2byteslice -package=assets -input=assets/plate.png -output=assets/plate.go -var=Plate
//go:generate file2byteslice -package=assets -input=assets/treat_cupcake.png -output=assets/treat_cupcake.go -var=Treat_cupcake
//go:generate file2byteslice -package=assets -input=assets/treat_donut.png -output=assets/treat_donut.go -var=Treat_donut
//go:generate file2byteslice -package=assets -input=assets/vegetable_carrot.png -output=assets/vegetable_carrot.go -var=Vegetable_carrot
//go:generate file2byteslice -package=assets -input=assets/vegetable_eggplant.png -output=assets/vegetable_eggplant.go -var=Vegetable_eggplant
//go:generate file2byteslice -package=assets -input=assets/vegetable_potato.png -output=assets/vegetable_potato.go -var=Vegetable_potato
//go:generate file2byteslice -package=assets -input=assets/vegetable_tomato.png -output=assets/vegetable_tomato.go -var=Vegetable_tomato

// sounds
//go:generate file2byteslice -package=assets -input=assets/soundtrack.mp3 -output=assets/soundtrack.go -var=Soundtrack
