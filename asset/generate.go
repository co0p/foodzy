package asset

/*
this file is used to generate asset files to be loaded by the game
*/

// images
//go:generate file2byteslice -package=asset -input=background.png -output=background.go -var=Background
//go:generate file2byteslice -package=asset -input=corn_baguette.png -output=corn_baguette.go -var=Corn_baguette
//go:generate file2byteslice -package=asset -input=corn_bread.png -output=corn_bread.go -var=Corn_bread
//go:generate file2byteslice -package=asset -input=corn_rice.png -output=corn_rice.go -var=Corn_rice
//go:generate file2byteslice -package=asset -input=dairy_cheese.png -output=dairy_cheese.go -var=Dairy_cheese
//go:generate file2byteslice -package=asset -input=dairy_milk.png -output=dairy_milk.go -var=Dairy_milk
//go:generate file2byteslice -package=asset -input=drink_beer.png -output=drink_beer.go -var=Drink_beer
//go:generate file2byteslice -package=asset -input=drink_coffee.png -output=drink_coffee.go -var=Drink_coffee
//go:generate file2byteslice -package=asset -input=drink_juice.png -output=drink_juice.go -var=Drink_juice
//go:generate file2byteslice -package=asset -input=drink_tea.png -output=drink_tea.go -var=Drink_tea
//go:generate file2byteslice -package=asset -input=drink_water.png -output=drink_water.go -var=Drink_water
//go:generate file2byteslice -package=asset -input=fish_crab.png -output=fish_crab.go -var=Fish_crab
//go:generate file2byteslice -package=asset -input=fish_sushi.png -output=fish_sushi.go -var=Fish_sushi
//go:generate file2byteslice -package=asset -input=fruit_apple.png -output=fruit_apple.go -var=Fruit_apple
//go:generate file2byteslice -package=asset -input=fruit_banana.png -output=fruit_banana.go -var=Fruit_banana
//go:generate file2byteslice -package=asset -input=fruit_grapes.png -output=fruit_grapes.go -var=Fruit_grapes
//go:generate file2byteslice -package=asset -input=fruit_orange.png -output=fruit_orange.go -var=Fruit_orange
//go:generate file2byteslice -package=asset -input=fruit_strawberry.png -output=fruit_strawberry.go -var=Fruit_strawberry
//go:generate file2byteslice -package=asset -input=meat_steak.png -output=meat_steak.go -var=Meat_steak
//go:generate file2byteslice -package=asset -input=plate.png -output=plate.go -var=Plate
//go:generate file2byteslice -package=asset -input=treat_cupcake.png -output=treat_cupcake.go -var=Treat_cupcake
//go:generate file2byteslice -package=asset -input=treat_donut.png -output=treat_donut.go -var=Treat_donut
//go:generate file2byteslice -package=asset -input=vegetable_carrot.png -output=vegetable_carrot.go -var=Vegetable_carrot
//go:generate file2byteslice -package=asset -input=vegetable_eggplant.png -output=vegetable_eggplant.go -var=Vegetable_eggplant
//go:generate file2byteslice -package=asset -input=vegetable_potato.png -output=vegetable_potato.go -var=Vegetable_potato
//go:generate file2byteslice -package=asset -input=vegetable_tomato.png -output=vegetable_tomato.go -var=Vegetable_tomato

// sounds
//go:generate file2byteslice -package=asset -input=soundtrack.mp3 -output=soundtrack.go -var=Soundtrack
//go:generate file2byteslice -package=asset -input=eating.mp3 -output=eating.go -var=Eating

// menu
//go:generate file2byteslice -package=asset -input=title.png -output=title.go -var=Title
//go:generate file2byteslice -package=asset -input=menu_start.png -output=menuStart.go -var=MenuStart
//go:generate file2byteslice -package=asset -input=menu_start_active.png -output=menuStartActive.go -var=MenuStartActive
//go:generate file2byteslice -package=asset -input=menu_quit.png -output=menuQuit.go -var=MenuQuit
//go:generate file2byteslice -package=asset -input=menu_quit_active.png -output=menuQuitActive.go -var=MenuQuitActive
