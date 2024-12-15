package system

import (
	"image"
	"kar"
	"kar/arc"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	damage                       float64 = 0.1
	blockHealth                  float64
	targetBlock                  image.Point
	placeBlock                   image.Point
	playerTile                   image.Point
	targetBlockID                uint16
	playerCenterX, playerCenterY float64
	isBlockPlaceable             bool
	IsRaycastHit                 bool
	DOP                          *arc.DrawOptions
)

type PlayerSys struct {
}

func (c *PlayerSys) Update() {
	q := arc.FilterPlayer.Query(&kar.WorldECS)
	for q.Next() {

		_, dop, anim, rect, _ := q.Get()

		playerCenterX, playerCenterY = rect.X+rect.W/2, rect.Y+rect.H/2

		PlayerController.AnimPlayer = anim
		DOP = dop
		PlayerController.UpdateInput()

		dx, dy := PlayerController.UpdatePhysics(rect.X, rect.Y, rect.W, rect.H)

		playerTile = Map.GetTileCoords(playerCenterX, playerCenterY)
		targetBlock, IsRaycastHit = Map.Raycast(playerTile, PlayerController.InputAxisLast, kar.BlockPlacementDistance)
		targetBlockID = Map.TileID(targetBlock)

		if IsRaycastHit {
			placeBlock = targetBlock.Sub(PlayerController.InputAxisLast)
			isBlockPlaceable = !rect.Overlaps(Map.GetTileRect(placeBlock))
		} else {
			isBlockPlaceable = false
			blockHealth = 0
		}

		rect.X += dx
		rect.Y += dy
		PlayerController.UpdateState()

		if PlayerController.IsPlaceKeyJustPressed {
			if IsRaycastHit && isBlockPlaceable && PlayerInventory.SelectedSlotQuantity() > 0 {
				Map.SetTile(placeBlock, PlayerInventory.SelectedSlotID())
				PlayerInventory.RemoveItemFromSelectedSlot()
			}
		}

	}

	if inpututil.IsKeyJustPressed(ebiten.KeyTab) {
		PlayerInventory.SelectNextSlot()
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		PlayerInventory.ClearSlot(PlayerInventory.SelectedSlot)
	}
}

func (c *PlayerSys) Draw() {

}
func (c *PlayerSys) Init() {

}
