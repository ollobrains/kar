package system

import (
	"kar"
	"kar/arc"
	"kar/engine/mathutil"
	"kar/items"
	"kar/res"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"
)

type Render struct{}

func (rn *Render) Init() {
}

func (rn *Render) Update() {

	kar.Camera.LookAt(playerCenterX, playerCenterY)
	q := arc.FilterAnimPlayer.Query(&kar.WorldECS)

	for q.Next() {
		a := q.Get()
		a.Update()
	}

}

func (rn *Render) Draw() {

	// Draw tilemap
	for y, row := range Map.Grid {
		for x, tileID := range row {

			if tileID != 0 {
				px, py := float64(x*Map.TileW), float64(y*Map.TileH)
				kar.GlobalDIO.GeoM.Reset()
				kar.GlobalDIO.GeoM.Scale(3, 3)
				kar.GlobalDIO.GeoM.Translate(px, py)
				if x == targetBlock.X && y == targetBlock.Y {
					i := mathutil.MapRange(blockHealth, 0, items.Property[tileID].MaxHealth, 0, 5)
					if res.Frames[tileID] != nil {
						kar.Camera.DrawWithColorM(res.Frames[tileID][int(i)], kar.GlobalColorM, kar.GlobalDIO, kar.Screen)
					}
				} else {
					kar.Camera.DrawWithColorM(GetSprite(tileID), kar.GlobalColorM, kar.GlobalDIO, kar.Screen)

				}
			}
		}
	}

	// Draw Items
	itemQuery := arc.FilterItem.Query(&kar.WorldECS)
	for itemQuery.Next() {
		id, _, dop, rect := itemQuery.Get()
		kar.GlobalDIO.GeoM.Reset()
		kar.GlobalDIO.GeoM.Scale(dop.Scale, dop.Scale)
		kar.GlobalDIO.GeoM.Translate(rect.X, rect.Y)
		kar.Camera.DrawWithColorM(GetSprite(id.ID), kar.GlobalColorM, kar.GlobalDIO, kar.Screen)
	}

	playerQuery := arc.FilterPlayer.Query(&kar.WorldECS)
	for playerQuery.Next() {
		_, dop, anim, rect, _ := playerQuery.Get()

		// Draw player
		sclX := dop.Scale
		kar.GlobalDIO.GeoM.Reset()
		if dop.FlipX {
			sclX *= -1
			kar.GlobalDIO.GeoM.Scale(sclX, dop.Scale)
			kar.GlobalDIO.GeoM.Translate(rect.X+rect.W, rect.Y)
		} else {
			kar.GlobalDIO.GeoM.Scale(sclX, dop.Scale)
			kar.GlobalDIO.GeoM.Translate(rect.X, rect.Y)
		}
		kar.Camera.DrawWithColorM(anim.CurrentFrame, kar.GlobalColorM, kar.GlobalDIO, kar.Screen)
	}

	// Draw all rects for debug
	if kar.DrawDebugHitboxes {
		rectQ := arc.FilterRect.Query(&kar.WorldECS)
		for rectQ.Next() {
			rect := rectQ.Get()
			x, y := kar.Camera.ApplyCameraTransformToPoint(rect.X, rect.Y)
			vector.StrokeRect(
				kar.Screen,
				float32(x),
				float32(y),
				float32(rect.W),
				float32(rect.H),
				1,
				colornames.Magenta,
				false,
			)
			// Draw player center
			vector.DrawFilledCircle(
				kar.Screen,
				float32(x+rect.W*0.5),
				float32(y+rect.H*0.5),
				2,
				colornames.Magenta,
				false,
			)
		}
	}

	// // Draw target tile border
	kar.GlobalDIO.GeoM.Translate(float64(targetBlock.X*Map.TileW), float64(targetBlock.Y*Map.TileH))
	kar.Camera.DrawWithColorM(res.Border, kar.GlobalColorM, kar.GlobalDIO, kar.Screen)

	// Draw debug info
	ebitenutil.DebugPrintAt(kar.Screen, PlayerController.CurrentState, 10, 10)
	ebitenutil.DebugPrintAt(kar.Screen, "InputAxis"+PlayerController.InputAxisLast.String(), 10, 30)
	ebitenutil.DebugPrintAt(kar.Screen, "Target Block"+targetBlock.String(), 10, 50)
	ebitenutil.DebugPrintAt(kar.Screen, "Place Block"+placeBlock.String(), 10, 50+20)
}
