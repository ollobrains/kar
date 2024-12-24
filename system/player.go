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
	damage                       float64 = 0.3
	blockHealth                  float64
	targetBlockPos               image.Point
	placeBlock                   image.Point
	playerTile                   image.Point
	playerCenterX, playerCenterY float64
	IsRayHit                     bool
)

type PlayerSys struct {
}

func (c *PlayerSys) Update() {
	q := arc.FilterPlayer.Query(&kar.WorldECS)
	for q.Next() {
		anim, _, dop, rect, _ := q.Get()
		PlayerController.DOP = dop
		PlayerController.AnimPlayer = anim
		playerCenterX, playerCenterY = rect.X+rect.W/2, rect.Y+rect.H/2
		PlayerController.UpdateInput()

		dx, dy := PlayerController.UpdatePhysics(rect.X, rect.Y, rect.W, rect.H)

		playerTile = Map.WorldToTile(playerCenterX, playerCenterY)
		targetBlockTemp := targetBlockPos
		targetBlockPos, IsRayHit = Map.Raycast(playerTile, PlayerController.InputAxisLast, kar.RaycastDist)
		// eğer block odağı değiştiyse saldırıyı sıfırla
		if !targetBlockPos.Eq(targetBlockTemp) || !IsRayHit {
			blockHealth = 0
		}
		rect.X += dx
		rect.Y += dy

		if PlayerController.CurrentState != "idle" {
			if PlayerController.InputAxisLast.Y == -1 {
				anim.SetStateAndReset("attackUp")
			}
			if PlayerController.InputAxisLast.Y == 1 {
				anim.SetStateAndReset("attackDown")
			}
		}
		PlayerController.UpdateState()

		if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
			PlayerInventory.SelectPrevSlot()
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyE) {
			PlayerInventory.SelectNextSlot()
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
			PlayerInventory.ClearSelectedSlot()
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			PlayerInventory.RandomFillAllSlots()

		}

		// Drop Item
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
			if PlayerInventory.SelectedSlotID() != items.Air {
				AppendToSpawnList(playerCenterX, playerCenterY, PlayerInventory.SelectedSlotID())
				PlayerInventory.RemoveItemFromSelectedSlot()
			}
		}

		// Place block
		if PlayerController.IsPlaceKeyJustPressed {
			anyItemOverlapsWithPlaceCoords := false
			if IsRayHit && items.IsBlock(PlayerInventory.SelectedSlotID()) {
				placeBlock = targetBlockPos.Sub(PlayerController.InputAxisLast)
				queryItem := arc.FilterItem.Query(&kar.WorldECS)
				for queryItem.Next() {
					_, _, itemRect, _ := queryItem.Get()
					anyItemOverlapsWithPlaceCoords = itemRect.Overlaps(Map.GetTileRect(placeBlock))
					if anyItemOverlapsWithPlaceCoords {
						queryItem.Close()
						break
					}
				}
				if !anyItemOverlapsWithPlaceCoords {
					if !rect.Overlaps(Map.GetTileRect(placeBlock)) {
						Map.SetTile(placeBlock, PlayerInventory.SelectedSlotID())
						PlayerInventory.RemoveItemFromSelectedSlot()
					}
				}
			}
		}

	}

}

func (c *PlayerSys) Draw() {

}
func (c *PlayerSys) Init() {

}
