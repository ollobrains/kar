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
	// damage                       float64 = 1
	blockHealth                  float64
	playerHealth                 float64
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
		anim, hlt, dop, rect, _ := q.Get()
		playerHealth = hlt.Health
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
		PlayerController.UpdateState()

		// Drop Item
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
			currentSlot := PlayerInventory.SelectedSlot()
			if currentSlot.ID != items.Air {
				AppendToSpawnList(playerCenterX, playerCenterY, currentSlot.ID, currentSlot.Durability)
				PlayerInventory.RemoveItemFromSelectedSlot()
			}
		}

		// Place block
		if PlayerController.IsPlaceKeyJustPressed {
			anyItemOverlapsWithPlaceCoords := false
			if IsRayHit && items.HasTag(PlayerInventory.SelectedSlot().ID, items.Block) {
				placeBlock = targetBlockPos.Sub(PlayerController.InputAxisLast)
				queryItem := arc.FilterItem.Query(&kar.WorldECS)
				for queryItem.Next() {
					_, itemRect, _, _ := queryItem.Get()
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

		if inpututil.IsKeyJustPressed(ebiten.KeyV) {
			kar.DrawDebugTextEnabled = !kar.DrawDebugTextEnabled
		}
	}

}

func (c *PlayerSys) Draw() {

}
func (c *PlayerSys) Init() {

}
