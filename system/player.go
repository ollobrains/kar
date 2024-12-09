package system

import (
	"kar"
	"kar/arc"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var targetBlockX, targetBlockY int
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

		// if ebiten.IsKeyPressed(ebiten.KeySpace) {

		// }

		targetBlockX, targetBlockY = Map.GetTileCoords(playerCenterX, playerCenterY)
		targetBlockX += int(PlayerController.InputAxisLast.X)
		targetBlockY += int(PlayerController.InputAxisLast.Y)
		if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
			Map.SetTile(targetBlockX, targetBlockY, 1)
		}

	}
}

func (c *Player) Draw() {

}
func (c *Player) Init() {

}
