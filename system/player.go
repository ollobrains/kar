package system

import (
	"image"
	"kar"
	"kar/arc"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var targetBlock image.Point
var playerCenterX, playerCenterY float64

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
		rect.X += dx
		rect.Y += dy
		PlayerController.UpdateState()

		if PlayerController.VelX > 0.01 {
			dop.FlipX = false
		} else if PlayerController.VelX < -0.01 {
			dop.FlipX = true
		}

		playerTile := Map.GetTileCoords(playerCenterX, playerCenterY)
		targetBlock = Map.Raycast(playerTile, PlayerController.InputAxisLast)
		if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
			if !targetBlock.Eq(image.Point{}) {
				placePos := targetBlock.Sub(PlayerController.InputAxisLast)
				if !placePos.Eq(playerTile) {
					Map.SetTile(placePos, 1)
				}
			}
		}
	}
}

func (c *Player) Draw() {

}
func (c *Player) Init() {

}
