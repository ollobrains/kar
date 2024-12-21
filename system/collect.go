package system

import (
	"kar"
	"kar/arc"
	"kar/engine/mathutil"
	"math"

	"github.com/mlange-42/arche/ecs"
)

var (
	itemGravity float64 = 3
	toRemove    []ecs.Entity
	sinspace    = mathutil.SinSpace(0, 2*math.Pi, 3, 60)
	sinspaceLen = len(sinspace) - 1
)

type Collect struct {
}

func (c *Collect) Init() {}
func (c *Collect) Update() {
	if kar.WorldECS.Alive(PlayerEntity) {
		collisionQuery := arc.FilterCollision.Query(&kar.WorldECS)
		for collisionQuery.Next() {
			PlayerRect := arc.MapRect.GetUnchecked(PlayerEntity)
			rect, itemID, timers := collisionQuery.Get()
			if timers.CollisionCountdown != 0 {
				timers.CollisionCountdown--
			} else {
				if PlayerRect.OverlapsRect(rect) {
					ok := PlayerInventory.AddItemIfEmpty(itemID.ID)
					if ok {
						toRemove = append(toRemove, collisionQuery.Entity())
					}
				}
			}
			dy := Collider.CollideY(rect.X, rect.Y+8, rect.W, rect.H, itemGravity)
			rect.Y += dy
			// rect.Y += sinspace[timers.AnimationIndex]
			timers.AnimationIndex = (timers.AnimationIndex + 1) % sinspaceLen
		}
		for _, e := range toRemove {
			kar.WorldECS.RemoveEntity(e)
		}
		toRemove = toRemove[:0]
	}
}
func (itm *Collect) Draw() {}
