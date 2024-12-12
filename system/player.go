package system

import (
	"image"
	"kar"
	"kar/arc"
	"kar/items"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var targetBlock image.Point
var targetBlockID uint16
var placeBlock image.Point
var playerCenterX, playerCenterY float64
var isBlockPlaceable bool
var damage float64 = 0.05

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
		targetBlockID = Map.TileID(targetBlock)
		if hit {
			placeBlock = targetBlock.Sub(PlayerController.InputAxisLast)
			isBlockPlaceable = !rect.Overlaps(Map.GetTileRect(placeBlock))
		} else {
			isBlockPlaceable = false
			blockHealth = 0
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
			if hit && isBlockPlaceable {
				Map.SetTile(placeBlock, items.RandomBlock())
			}
		}

		if hit {
			if ebiten.IsKeyPressed(ebiten.KeyRight) {
				if items.IsBreakable(targetBlockID) {
					blockHealth += damage
				}
				if blockHealth >= items.Property[targetBlockID].MaxHealth {
					Map.SetTile(targetBlock, 0)
					blockHealth = 0
				}
			} else {
				blockHealth = 0
			}
		}
	}
}

var blockHealth float64

func (c *Player) Draw() {

}
func (c *Player) Init() {

}
