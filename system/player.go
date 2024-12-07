package system

import (
	"kar"
	"kar/arc"
)

var CurrentState string

type Player struct {
	controller *Controller
}

func (c *Player) Init() {
	c.controller = NewController(0, 5, Collider)
	c.controller.Collider = Collider
	c.controller.SetScale(3)
}

func (c *Player) Update() {
	q := arc.FilterMovement.Query(&kar.WorldECS)
	for q.Next() {
		rect, anim, dop := q.Get()
		c.controller.AnimPlayer = anim
		c.controller.UpdateInput()
		dx, dy := c.controller.UpdatePhysics(rect.X, rect.Y, rect.W, rect.H)
		rect.X += dx
		rect.Y += dy

		if c.controller.VelX > 0.01 {
			dop.FlipX = false
		} else if c.controller.VelX < -0.01 {
			dop.FlipX = true
		}

		c.controller.UpdateState()
		CurrentState = c.controller.CurrentState
	}
}

func (c *Player) Draw() {

}
