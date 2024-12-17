package system

import (
	"kar"
	"kar/res"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

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

		kar.GlobalDIO.GeoM.Reset()
		kar.GlobalDIO.GeoM.Translate(-64, -11)
		kar.GlobalDIO.GeoM.Scale(2, 2)
		kar.GlobalDIO.GeoM.Translate((kar.ScreenW/2)-1, kar.ScreenH-40)
		kar.Screen.DrawImage(res.Hotbar, kar.GlobalDIO)

		// Draw hotbar selected border
		kar.GlobalDIO.GeoM.Translate(-2, -2)
		selectedOffsetX := float64(PlayerInventory.SelectedSlot) * 40
		kar.GlobalDIO.GeoM.Translate(selectedOffsetX, 0)
		kar.Screen.DrawImage(res.HotbarSelection, kar.GlobalDIO)

		// Draw hotbar slots
		for x := range 9 {
			quantity := PlayerInventory.Slots[x].Quantity
			offsetX := (float64(x) * 40) + 320
			kar.GlobalDIO.GeoM.Reset()
			kar.GlobalDIO.GeoM.Translate(-8, -8)
			kar.GlobalDIO.GeoM.Scale(2, 2)
			kar.GlobalDIO.GeoM.Translate(offsetX, kar.ScreenH-40)
			if PlayerInventory.Slots[x].Quantity > 0 {
				kar.Screen.DrawImage(GetSprite(PlayerInventory.Slots[x].ID), kar.GlobalDIO)
			}
			gui.itemQuantityTextDO.GeoM.Reset()
			gui.itemQuantityTextDO.GeoM.Translate(offsetX-8, kar.ScreenH-45)
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
