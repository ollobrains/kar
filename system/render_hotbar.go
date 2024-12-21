package system

import (
	"kar"
	"kar/res"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var hotbarPositionX = 16.0
var hotbarPositionY = 16.0

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
		for y := range 9 {
			// draw item
			quantity := PlayerInventory.Slots[y].Quantity
			SlotOffsetY := float64(y) * 24
			SlotOffsetY += hotbarPositionY
			kar.GlobalDIO.GeoM.Reset()
			kar.GlobalDIO.GeoM.Scale(2, 2)
			// kar.GlobalDIO.GeoM.Translate(SlotOffsetX, hotbarPositionY)
			kar.GlobalDIO.GeoM.Translate(hotbarPositionX, SlotOffsetY)
			if PlayerInventory.Slots[y].Quantity > 0 {
				colorm.DrawImage(kar.Screen, GetSprite(PlayerInventory.Slots[y].ID), kar.GlobalColorM, kar.GlobalDIO)
			}

			// draw border
			if y == PlayerInventory.SelectedSlot {
				kar.GlobalDIO.GeoM.Translate(-10, -2)
				colorm.DrawImage(kar.Screen, res.SelectionBar, kar.GlobalColorM, kar.GlobalDIO)
			}

			// draw text
			gui.itemQuantityTextDO.GeoM.Reset()
			gui.itemQuantityTextDO.GeoM.Translate(hotbarPositionX+24, SlotOffsetY)
			// text.Draw(kar.Screen, strconv.FormatUint(uint64(quantity), 10), res.Font, gui.itemQuantityTextDO)
			if quantity > 0 {
				num := strconv.FormatUint(uint64(quantity), 10)
				// if quantity < 10 {
				// 	num = " " + num
				// }
				text.Draw(kar.Screen, num, res.Font, gui.itemQuantityTextDO)
			}
		}
	}
}
