package system

import (
	"kar"
	"kar/arc"
)

type Player struct {
}

func (c *Player) Update() {
	q := arc.FilterMovement.Query(&kar.WorldECS)
	for q.Next() {

		rect, anim, dop := q.Get()

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

	}
}

func (c *Player) Draw() {

}
func (c *Player) Init() {

}
