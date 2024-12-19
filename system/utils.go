package system

import (
	"kar/items"
	"kar/res"

	"github.com/hajimehoshi/ebiten/v2"
)

func GetSprite(id uint16) *ebiten.Image {
	im, ok := res.Images[id]
	if ok {
		return im
	} else {
		if len(res.Frames[id]) > 0 {
			return res.Frames[id][0]
		} else {
			return res.Images[items.Air]
		}
	}
}
