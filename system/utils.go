package system

import (
	"kar/items"
	"kar/res"

	"github.com/hajimehoshi/ebiten/v2"
)

func GetSprite(id uint16) *ebiten.Image {
	im, ok := res.ItemIcons[id]
	if ok {
		return im
	} else {
		if len(res.BlockCrackFrames[id]) > 0 {
			return res.BlockCrackFrames[id][0]
		} else {
			return res.ItemIcons[items.Air]
		}
	}
}
