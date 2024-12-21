package system

import (
	"kar"
	"kar/res"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var hotbarPositionX = kar.ScreenW / 3
var hotbarPositionY = 8.0

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
		for x := range 9 {
			quantity := PlayerInventory.Slots[x].Quantity
			SlotOffsetX := float64(x) * 34
			SlotOffsetX += hotbarPositionX
			kar.GlobalDIO.GeoM.Reset()
			kar.GlobalDIO.GeoM.Scale(2, 2)
			kar.GlobalDIO.GeoM.Translate(SlotOffsetX, hotbarPositionY)
			if PlayerInventory.Slots[x].Quantity > 0 {
				colorm.DrawImage(kar.Screen, GetSprite(PlayerInventory.Slots[x].ID), kar.GlobalColorM, kar.GlobalDIO)
			}

			if x == PlayerInventory.SelectedSlot {
				kar.GlobalDIO.GeoM.Translate(-2, -2)
				colorm.DrawImage(kar.Screen, res.SelectionBar, kar.GlobalColorM, kar.GlobalDIO)
			}

			gui.itemQuantityTextDO.GeoM.Reset()
			gui.itemQuantityTextDO.GeoM.Translate(SlotOffsetX+10, hotbarPositionY+13)
			if quantity > 0 {

				num := strconv.FormatUint(uint64(quantity), 10)
				if quantity < 10 {
					num = " " + num
				}
				text.Draw(kar.Screen, num, res.Font, gui.itemQuantityTextDO)
			}
		}
	}
}
