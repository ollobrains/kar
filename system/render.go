package system

import (
	"kar"
	"kar/arc"
	"kar/engine/mathutil"
	"kar/items"
	"kar/res"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/setanarut/kamera/v2"
	"golang.org/x/image/colornames"
)

type Render struct{}

func (rn *Render) Init() {
	kar.Camera.SmoothType = kamera.None
}

func (rn *Render) Update() {

	// kar.Camera.LookAt(400, 450)
	kar.Camera.LookAt(playerCenterX, playerCenterY)
	q := arc.FilterAnimPlayer.Query(&kar.WorldECS)

	for q.Next() {
		a := q.Get()
		a.Update()
	}

}

func (rn *Render) Draw() {

	// Draw tilemap
	camMin := Map.WorldToTile(kar.Camera.TopLeft())
	camMin.X = min(max(camMin.X, 0), Map.W)
	camMin.Y = min(max(camMin.Y, 0), Map.H)
	camMaxX := min(max(camMin.X+19, 0), Map.W)
	camMaxY := min(max(camMin.Y+11, 0), Map.H)

	for y := camMin.Y; y < camMaxY; y++ {
		for x := camMin.X; x < camMaxX; x++ {
			tileID := Map.Grid[y][x]
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

	playerQuery := arc.FilterPlayer.Query(&kar.WorldECS)
	for playerQuery.Next() {
		anim, _, dop, rect, _ := playerQuery.Get()

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

	// Draw Items
	itemQuery := arc.FilterItem.Query(&kar.WorldECS)
	for itemQuery.Next() {
		id, _, rect, _ := itemQuery.Get()
		kar.GlobalDIO.GeoM.Reset()
		kar.GlobalDIO.GeoM.Scale(kar.ItemScale, kar.ItemScale)
		kar.GlobalDIO.GeoM.Translate(rect.X, rect.Y)
		kar.Camera.DrawWithColorM(GetSprite(id.ID), kar.GlobalColorM, kar.GlobalDIO, kar.Screen)
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
