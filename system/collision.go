package system

import (
	"fmt"
	"github.com/co0p/foodzy/component"
	"github.com/co0p/foodzy/entity"
	"github.com/hajimehoshi/ebiten/v2"
)

type CollisionSystem struct {
	manager *entity.Manager
}

func NewCollisionSystem(manager *entity.Manager) *CollisionSystem {
	return &CollisionSystem{manager: manager}
}

func (c *CollisionSystem) Draw(image *ebiten.Image) {}

func (c *CollisionSystem) Update() error {

	e := c.manager.QueryByTag("player")
	if len(e) != 1 {
		panic(fmt.Sprintf("expected one entity with tag 'player', got %v", len(e)))
	}
	player := e[0]
	playerTransform := player.GetComponent(component.TransformType).(*component.Transform)
	playerDim := player.GetComponent(component.CollisionType).(*component.Collision)
	playerBox := boundingBox{
		x:      playerTransform.X,
		y:      playerTransform.Y,
		width:  playerDim.Width * playerTransform.Scale,
		height: playerDim.Height * playerTransform.Scale,
	}

	otherEntities := c.manager.QueryByComponents(component.NutrientType, component.CollisionType, component.TransformType)
	for _, entity := range otherEntities {
		if player == entity || !entity.Active {
			continue
		}

		entityPos := entity.GetComponent(component.TransformType).(*component.Transform)
		entityDim := entity.GetComponent(component.CollisionType).(*component.Collision)
		entityBox := boundingBox{
			x:      entityPos.X,
			y:      entityPos.Y,
			width:  entityDim.Width * entityPos.Scale,
			height: entityDim.Height * entityPos.Scale,
		}

		if playerBox.AABBCollision(entityBox) {
			entity.Active = false
			playerNutrients := player.GetComponent(component.NutrientType).(*component.Nutrient)
			entityNutrients := entity.GetComponent(component.NutrientType).(*component.Nutrient)
			playerNutrients.Add(entityNutrients)
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

func (rect1 *boundingBox) AABBCollision(rect2 boundingBox) bool {
	return rect1.x < rect2.x+rect2.width &&
		rect1.x+rect1.width > rect2.x &&
		rect1.y < rect2.y+rect2.height &&
		rect1.y+rect1.height > rect2.y
}
