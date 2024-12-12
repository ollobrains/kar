package system

import (
	"image"
	"kar"
	"kar/arc"
	"kar/items"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	targetBlock                  image.Point
	targetBlockID                uint16
	placeBlock                   image.Point
	playerCenterX, playerCenterY float64
	isBlockPlaceable             bool
	raycastHit                   bool
	damage                       float64 = 0.05
)

type PlayerSys struct {
}

func (c *PlayerSys) Update() {
	q := arc.FilterPlayer.Query(&kar.WorldECS)
	for q.Next() {

		_, dop, anim, rect, _ := q.Get()

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

		playerTile := Map.GetTileCoords(playerCenterX, playerCenterY)
		targetBlock, raycastHit = Map.Raycast(playerTile, PlayerController.InputAxisLast, kar.BlockPlacementDistance)
		targetBlockID = Map.TileID(targetBlock)
		if raycastHit {
			placeBlock = targetBlock.Sub(PlayerController.InputAxisLast)
			isBlockPlaceable = !rect.Overlaps(Map.GetTileRect(placeBlock))
		} else {
			isBlockPlaceable = false
			blockHealth = 0
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
			if raycastHit && isBlockPlaceable && PlayerInventory.SelectedQuantity() > 0 {
				Map.SetTile(placeBlock, PlayerInventory.SelectedID())
				PlayerInventory.RemoveItemFromSelected()
			}
		}

		if raycastHit {
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

	if inpututil.IsKeyJustPressed(ebiten.KeyTab) {
		PlayerInventory.SelectNextSlot()
	}
}

var blockHealth float64

func (c *PlayerSys) Draw() {

}
func (c *PlayerSys) Init() {

}
