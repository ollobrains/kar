package system

import (
	"kar"
	"kar/arc"

	"github.com/mlange-42/arche/ecs"
)

var itemGravity float64 = 5
var toRemove []ecs.Entity

type Collect struct {
}

func (c *Collect) Init() {}
func (c *Collect) Update() {

	collisionQuery := arc.FilterCollision.Query(&kar.WorldECS)

	for collisionQuery.Next() {
		PlayerRect := arc.MapRect.GetUnchecked(PlayerEntity)
		rect, itemID := collisionQuery.Get()
		if PlayerRect.OverlapsRect(rect) {
			PlayerInventory.AddItemIfEmpty(itemID.ID)
			toRemove = append(toRemove, collisionQuery.Entity())
		}

		dx, dy := Collider.Collide(rect.X, rect.Y, rect.W, rect.H, 0, itemGravity, nil)
		rect.X += dx
		rect.Y += dy
	}

	for _, e := range toRemove {
		kar.WorldECS.RemoveEntity(e)
	}
	toRemove = toRemove[:0]

}
func (itm *Collect) Draw() {}
