package systems

import (
	"fmt"
	"github.com/co0p/foodzy/components"
	"github.com/co0p/foodzy/entities"
	"github.com/hajimehoshi/ebiten/v2"
)

type CollisionSystem struct {
	manager *entities.Manager
}

func NewCollisionSystem(manager *entities.Manager) *CollisionSystem {
	return &CollisionSystem{manager: manager}
}

func (c *CollisionSystem) Draw(image *ebiten.Image) {}

func (c *CollisionSystem) Update() error {

	e := c.manager.QueryByTag("player")
	if len(e) != 1 {
		panic(fmt.Sprintf("expected one entity with tag 'player', got %v", len(e)))
	}
	player := e[0]
	playerPos := player.GetComponent(&components.Position{}).(*components.Position)
	playerDim := player.GetComponent(&components.Dimension{}).(*components.Dimension)
	playerBox := boundingBox{
		x:      playerPos.X,
		y:      playerPos.Y,
		width:  playerDim.Width,
		height: playerDim.Height,
	}

	playerNutrients := player.GetComponent(&components.Nutrient{}).(*components.Nutrient)

	otherEntities := c.manager.QueryByComponents(&components.Nutrient{}, &components.Collision{}, &components.Position{})
	for _, entity := range otherEntities {
		if player == entity {
			continue
		}

		entityNutrients := entity.GetComponent(&components.Nutrient{}).(*components.Nutrient)

		entityPos := entity.GetComponent(&components.Position{}).(*components.Position)
		entityDim := entity.GetComponent(&components.Dimension{}).(*components.Dimension)
		entityBox := boundingBox{
			x:      entityPos.X,
			y:      entityPos.Y,
			width:  entityDim.Width,
			height: entityDim.Height,
		}

		if c.AABBCollision(playerBox, entityBox) && entity.Active {
			entity.Active = false
			playerNutrients.Fat += entityNutrients.Fat
			playerNutrients.Water += entityNutrients.Water
			playerNutrients.Carbohydrates += entityNutrients.Carbohydrates
			playerNutrients.Minerals += entityNutrients.Minerals
			playerNutrients.Protein += entityNutrients.Protein
			playerNutrients.Vitamins += entityNutrients.Vitamins
		}
	}

	return nil
}

type boundingBox struct {
	x      float64
	y      float64
	width  float64
	height float64
}

func (c *CollisionSystem) AABBCollision(rect1 boundingBox, rect2 boundingBox) bool {

	return rect1.x < rect2.x+rect2.width &&
		rect1.x+rect1.width > rect2.x &&
		rect1.y < rect2.y+rect2.height &&
		rect1.y+rect1.height > rect2.y
}
