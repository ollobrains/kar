package system

import (
	"image"
	"kar"
	"kar/arc"
	"kar/engine/mathutil"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var targetBlock image.Point
var playerCenterX, playerCenterY float64
var isBlockPlaceable bool

type Player struct {
}

func (c *Player) Update() {
	q := arc.FilterMovement.Query(&kar.WorldECS)
	for q.Next() {

		rect, anim, dop := q.Get()

		playerCenterX, playerCenterY = rect.X+rect.W/2, rect.Y+rect.H/2

		PlayerController.AnimPlayer = anim
		PlayerController.UpdateInput()

		dx, dy := PlayerController.UpdatePhysics(rect.X, rect.Y, rect.W, rect.H)

		if PlayerController.VelX > 0.01 {
			dop.FlipX = false // sağa gidiyor
			PlayerController.InputAxisLast.X = 1
		} else if PlayerController.VelX < -0.01 {
			dop.FlipX = true // sola gidiyor
			PlayerController.InputAxisLast.X = -1
		}

		rect.X += dx
		rect.Y += dy
		PlayerController.UpdateState()

		var hit bool
		playerTile := Map.GetTileCoords(playerCenterX, playerCenterY)
		targetBlock, hit = Map.Raycast(playerTile, PlayerController.InputAxisLast, kar.BlockPlacementDistance)

		if hit {
			targetBlock = targetBlock.Sub(PlayerController.InputAxisLast)
			isBlockPlaceable = !rect.Overlaps(Map.GetTileRect(targetBlock))
		} else {
			isBlockPlaceable = false
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
			if hit && isBlockPlaceable {
				Map.SetTile(targetBlock, uint16(mathutil.RandRangeInt(1, 100)))
			}
		}
	}
}

func (c *Player) Draw() {

}
func (c *Player) Init() {

}
