package system

import (
	"fmt"
	"kar"
	"kar/items"
	"kar/res"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var hotbarPositionX = 8.
var hotbarPositionY = 8.
var itemQuantityTextDO = &text.DrawOptions{}

type DrawHotbar struct {
	itemQuantityTextDO *text.DrawOptions
}

func (gui *DrawHotbar) Init() {
	gui.itemQuantityTextDO = &text.DrawOptions{}
}

func (gui *DrawHotbar) Update() {
}
func (gui *DrawHotbar) Draw() {
	if kar.WorldECS.Alive(PlayerEntity) {
		// Background
		kar.GlobalDIO.GeoM.Reset()
		kar.GlobalDIO.GeoM.Scale(2, 2)
		kar.GlobalDIO.GeoM.Translate(hotbarPositionX, hotbarPositionY)
		colorm.DrawImage(kar.Screen, res.Hotbar, kar.GlobalColorM, kar.GlobalDIO)

		for x := range 9 {
			// draw item
			slotID := PlayerInventory.Slots[x].ID
			quantity := PlayerInventory.Slots[x].Quantity
			SlotOffsetX := float64(x) * 34
			SlotOffsetX += hotbarPositionX
			kar.GlobalDIO.GeoM.Reset()
			kar.GlobalDIO.GeoM.Scale(2, 2)
			kar.GlobalDIO.GeoM.Translate(SlotOffsetX+8, hotbarPositionY+8)
			if slotID != items.Air && PlayerInventory.Slots[x].Quantity > 0 {
				colorm.DrawImage(kar.Screen, res.Icon8[slotID], kar.GlobalColorM, kar.GlobalDIO)
			}

			if x == PlayerInventory.SelectedSlotIndex {
				// draw border
				kar.GlobalDIO.GeoM.Translate(-10, -10)
				colorm.DrawImage(kar.Screen, res.SelectionBar, kar.GlobalColorM, kar.GlobalDIO)

				// draw display name
				itemQuantityTextDO.GeoM.Reset()
				itemQuantityTextDO.GeoM.Translate(SlotOffsetX, hotbarPositionY+32)
				text.Draw(kar.Screen, items.Property[slotID].DisplayName, res.Font, itemQuantityTextDO)
			}

			itemQuantityTextDO.GeoM.Reset()
			itemQuantityTextDO.GeoM.Translate(SlotOffsetX+10, hotbarPositionY+13)

			// draw text
			if quantity > 0 {
				num := strconv.FormatUint(uint64(quantity), 10)
				if quantity < 10 {
					num = " " + num
				}
				text.Draw(kar.Screen, num, res.Font, itemQuantityTextDO)
			}
		}

		// Draw debug info
		if kar.DrawDebugTextEnabled {
			ebitenutil.DebugPrintAt(kar.Screen, fmt.Sprintf(stats,
				PlayerController.CurrentState,
				PlayerController.InputAxisLast,
				targetBlockPos,
				blockHealth,
			), 10, 50)
		}
	}
}

var stats = `state %v
inputAxisLast %v
targeBlock %v
blockHealth %v
`
