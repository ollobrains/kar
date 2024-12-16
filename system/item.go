package system

import (
	"fmt"
	"kar"
	"kar/arc"

	"github.com/mlange-42/arche/ecs"
)

var toRemove []ecs.Entity

type Item struct {
}

func (itm *Item) Init() {}
func (itm *Item) Update() {

	// Draw Item
	itemQuery := arc.FilterItem.Query(&kar.WorldECS)

	for itemQuery.Next() {
		_, _, _, PlayerRect, _ := arc.MapPlayer.GetUnchecked(PlayerEntity)
		_, _, _, itemRect := itemQuery.Get()

		if itemRect.OverlapsR(PlayerRect) {
			toRemove = append(toRemove, itemQuery.Entity())
		}

	}

	// The world is unlocked again.
	// Actually remove the collected entities.
	for _, e := range toRemove {
		fmt.Println("çarptı")
		kar.WorldECS.RemoveEntity(e)
	}

	// Empty the slice, so we can reuse it in the next time step.
	toRemove = toRemove[:0]

}
func (itm *Item) Draw() {}
