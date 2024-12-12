package system

import (
	"kar"
	"kar/res"
	"strconv"

	eb "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type DrawHotbar struct {
	hotbarDIO, itemsDIO *eb.DrawImageOptions
	itemQuantityTextDO  *text.DrawOptions
}

func (gui *DrawHotbar) Init() {
	gui.hotbarDIO = &eb.DrawImageOptions{}
	gui.itemsDIO = &eb.DrawImageOptions{}
	gui.itemQuantityTextDO = &text.DrawOptions{}
}

func (gui *DrawHotbar) Update() {
}
func (gui *DrawHotbar) Draw() {

	if kar.WorldECS.Alive(PlayerEntity) {

		gui.hotbarDIO.GeoM.Reset()
		gui.hotbarDIO.GeoM.Translate(-64, -11)
		gui.hotbarDIO.GeoM.Scale(2, 2)
		gui.hotbarDIO.GeoM.Translate((kar.ScreenW/2)-1, kar.ScreenH-40)
		kar.Screen.DrawImage(res.Hotbar, gui.hotbarDIO)

		// Draw hotbar selected border
		gui.hotbarDIO.GeoM.Translate(-2, -2)
		selectedOffsetX := float64(PlayerInventory.SelectedSlot) * 40
		gui.hotbarDIO.GeoM.Translate(selectedOffsetX, 0)
		kar.Screen.DrawImage(res.HotbarSelection, gui.hotbarDIO)

		// Draw hotbar slots
		for x := range 9 {
			quantity := PlayerInventory.Slots[x].Quantity
			offsetX := (float64(x) * 40) + 320
			gui.itemsDIO.GeoM.Reset()
			gui.itemsDIO.GeoM.Translate(-8, -8)
			gui.itemsDIO.GeoM.Scale(2, 2)
			gui.itemsDIO.GeoM.Translate(offsetX, kar.ScreenH-40)
			if PlayerInventory.Slots[x].Quantity > 0 {
				kar.Screen.DrawImage(GetSprite(PlayerInventory.Slots[x].ID), gui.itemsDIO)
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
