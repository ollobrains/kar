package system

import (
	"kar"
	"kar/arc"

	"github.com/mlange-42/arche/ecs"
)

var itemGravity float64 = 3
var toRemove []ecs.Entity

type Collect struct {
}

func (c *Collect) Init() {}
func (c *Collect) Update() {
	if kar.WorldECS.Alive(PlayerEntity) {
		collisionQuery := arc.FilterCollision.Query(&kar.WorldECS)
		for collisionQuery.Next() {
			PlayerRect := arc.MapRect.GetUnchecked(PlayerEntity)
			rect, itemID := collisionQuery.Get()
			if PlayerRect.OverlapsRect(rect) {
				ok := PlayerInventory.AddItemIfEmpty(itemID.ID)
				if ok {
					toRemove = append(toRemove, collisionQuery.Entity())
				}
			}

			dy := Collider.CollideY(rect.X, rect.Y, rect.W, rect.H, itemGravity)
			rect.Y += dy
		}
		for _, e := range toRemove {
			kar.WorldECS.RemoveEntity(e)
		}
		toRemove = toRemove[:0]
	}
}
func (itm *Collect) Draw() {}
