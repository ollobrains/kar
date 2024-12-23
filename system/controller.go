package system

import (
	"image"
	"kar/arc"
	"kar/engine/mathutil"
	"kar/items"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/setanarut/anim"
	"github.com/setanarut/tilecollider"
)

type Controller struct {
	DOP          *arc.DrawOptions
	AnimPlayer   *anim.AnimationPlayer
	CurrentState string
	Collider     *tilecollider.Collider[uint16]

	VelX                                float64
	VelY                                float64
	JumpPower                           float64
	Gravity                             float64
	MaxFallSpeed                        float64
	MaxRunSpeed                         float64
	MaxWalkSpeed                        float64
	Acceleration                        float64
	Deceleration                        float64
	JumpHoldTime                        float64
	JumpBoost                           float64
	JumpTimer                           float64
	MinSpeedThresForJumpBoostMultiplier float64 // Yüksek zıplama için gereken minimum hız
	JumpBoostMultiplier                 float64 // Yüksek zıplamada kullanılacak çarpan
	SpeedJumpFactor                     float64 // Yatay hızın zıplama yüksekliğine etkisini kontrol eden çarpan
	ShortJumpVelocity                   float64 // Kısa zıplama için hız
	JumpReleaseTimer                    float64 // Zıplama tuşu bırakıldığında geçen süre

	IsOnFloor  bool
	IsSkidding bool
	IsFalling  bool

	SkiddingJumpEnabled bool

	// Input durumları
	IsBreakKeyPressed     bool
	IsPlaceKeyJustPressed bool
	IsJumpKeyPressed      bool
	IsJumpKeyJustPressed  bool
	IsRunKeyPressed       bool
	InputAxis             image.Point
	InputAxisLast         image.Point

	WalkAcceleration float64
	WalkDeceleration float64
	RunAcceleration  float64
	RunDeceleration  float64

	HorizontalVelocity float64
	// Durum değişikliği için yeni alan
	previousState string
}

func NewController(velX, velY float64, tc *tilecollider.Collider[uint16]) *Controller {
	return &Controller{
		CurrentState:                        "falling",
		Collider:                            tc,
		VelX:                                velX,
		VelY:                                velY,
		JumpPower:                           -3.7,
		Gravity:                             0.19,
		MaxFallSpeed:                        6.0,
		Acceleration:                        0.08,
		Deceleration:                        0.1,
		JumpHoldTime:                        20.0,
		JumpBoost:                           -0.1,
		MinSpeedThresForJumpBoostMultiplier: 0.1,
		JumpBoostMultiplier:                 1.01,
		SpeedJumpFactor:                     0.3,
		ShortJumpVelocity:                   -2.0,
		JumpReleaseTimer:                    5,
		MaxWalkSpeed:                        2.0,
		MaxRunSpeed:                         3.0,
		WalkAcceleration:                    0.04,
		WalkDeceleration:                    0.04,
		RunAcceleration:                     0.04,
		RunDeceleration:                     0.04,
	}
}

func (c *Controller) SetScale(s float64) {
	c.JumpPower *= s
	c.Gravity *= s
	c.MaxFallSpeed *= s
	c.Acceleration *= s
	c.Deceleration *= s
	c.JumpBoost *= s
	c.MinSpeedThresForJumpBoostMultiplier *= s
	c.ShortJumpVelocity *= s
	c.MaxWalkSpeed *= s
	c.MaxRunSpeed *= s
	c.WalkAcceleration *= s
	c.WalkDeceleration *= s
	c.RunAcceleration *= s
	c.RunDeceleration *= s
}
func (c *Controller) UpdateInput() {
	c.IsBreakKeyPressed = ebiten.IsKeyPressed(ebiten.KeyRight)
	c.IsRunKeyPressed = ebiten.IsKeyPressed(ebiten.KeyShift)
	c.IsJumpKeyPressed = ebiten.IsKeyPressed(ebiten.KeySpace)
	c.IsPlaceKeyJustPressed = inpututil.IsKeyJustPressed(ebiten.KeyLeft)
	c.IsJumpKeyJustPressed = inpututil.IsKeyJustPressed(ebiten.KeySpace)
	c.InputAxis = image.Point{}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		c.InputAxis.Y -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		c.InputAxis.Y += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		c.InputAxis.X -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		c.InputAxis.X += 1
	}
	if !c.InputAxis.Eq(image.Point{}) {
		c.InputAxisLast = c.InputAxis
	}
}

func (c *Controller) UpdatePhysics(x, y, w, h float64) (dx, dy float64) {
	maxSpeed := c.MaxWalkSpeed
	currentAccel := c.WalkAcceleration
	currentDecel := c.WalkDeceleration
	c.HorizontalVelocity = math.Abs(c.VelX)

	if !c.IsSkidding {
		if c.IsRunKeyPressed {
			maxSpeed = c.MaxRunSpeed
			currentAccel = c.RunAcceleration
			currentDecel = c.RunDeceleration
		} else if c.HorizontalVelocity > c.MaxWalkSpeed {
			currentDecel = c.RunDeceleration
		}
	}

	if c.InputAxis.X > 0 {
		if c.VelX > maxSpeed {
			c.VelX = max(maxSpeed, c.VelX-currentDecel)
		} else {
			c.VelX = min(maxSpeed, c.VelX+currentAccel)
		}
	} else if c.InputAxis.X < 0 {
		if c.VelX < -maxSpeed {
			c.VelX = min(-maxSpeed, c.VelX+currentDecel)
		} else {
			c.VelX = max(-maxSpeed, c.VelX-currentAccel)
		}
	} else {
		if c.VelX > 0 {
			c.VelX = max(0, c.VelX-currentDecel)
		} else if c.VelX < 0 {
			c.VelX = min(0, c.VelX+currentDecel)
		}
	}

	c.IsSkidding = (c.VelX > 0 && c.InputAxis.X == -1) || (c.VelX < 0 && c.InputAxis.X == 1)
	c.VelY += c.Gravity
	c.VelY = min(c.MaxFallSpeed, c.VelY)

	if c.VelX > 0.01 {
		c.DOP.FlipX = false // sağa gidiyor
		c.InputAxisLast.X = 1
	} else if c.VelX < -0.01 {
		c.DOP.FlipX = true // sola gidiyor
		c.InputAxisLast.X = -1
	}

	return c.Collider.Collide(math.Round(x), y, w, h, c.VelX, c.VelY, c.handleCollision)

}

func (c *Controller) handleCollision(ci []tilecollider.CollisionInfo[uint16], dx, dy float64) {
	c.IsOnFloor = false
	for _, v := range ci {
		if v.Normal[1] == -1 {
			c.VelY = 0
			c.IsOnFloor = true
		}
		if v.Normal[1] == 1 {
			c.VelY = 0
		}
		if v.Normal[0] == -1 {
			c.VelX = 0
		}
		if v.Normal[0] == 1 {
			c.VelX = 0
		}
	}
}

func (c *Controller) ResetVelocity() {
	c.VelX = 0
	c.HorizontalVelocity = 0
}

func (c *Controller) Skidding() {
	if c.SkiddingJumpEnabled && c.IsJumpKeyJustPressed {
		c.ResetVelocity()

		// Yeni yöne doğru çok küçük sabit değerle başla
		if c.InputAxis.X > 0 {
			c.VelX = 0.3
		} else if c.InputAxis.X < 0 {
			c.VelX = -0.3
		}

		c.VelY = c.JumpPower * 0.7 // Zıplama gücünü azalt
		c.JumpTimer = 0
		c.changeState("jumping")
		return
	}

	// Mevcut mantık devam eder...
	if c.HorizontalVelocity < 0.01 {
		c.changeState("idle")
	} else if !c.IsSkidding {
		if c.HorizontalVelocity > c.MaxWalkSpeed {
			c.changeState("running")
		} else {
			c.changeState("walking")
		}
	}
}

func (c *Controller) Falling() {
	if c.IsOnFloor {
		if c.HorizontalVelocity <= 0 {
			c.changeState("idle")
		} else if c.IsRunKeyPressed {
			c.changeState("running")
		} else {
			c.changeState("walking")
		}
	}
}

func (c *Controller) Attacking() {
	if IsRayHit {
		targetBlockID := Map.TileID(targetBlockPos)
		if items.IsBreakable(targetBlockID) {
			blockHardness := items.Property[targetBlockID].MaxHealth
			if blockHealth < blockHardness/4 {
				blockHealth += blockHardness / 4
			}
			blockHealth += damage
		}
		// Destroy block
		if blockHealth >= items.Property[targetBlockID].MaxHealth {
			blockHealth = 0
			Map.SetTile(targetBlockPos, items.Air)
			// spawn drop item
			x, y := Map.TileToWorld(targetBlockPos)
			AppendToSpawnList(x, y, items.Property[targetBlockID].DropID)
		}
	}

	if !IsRayHit {
		c.changeState("idle")
	}

	if !c.IsOnFloor && c.VelY > 0.01 {
		c.changeState("falling")
	} else if !c.IsBreakKeyPressed && c.IsOnFloor {
		c.changeState("idle")
	} else if !c.IsBreakKeyPressed && c.IsJumpKeyJustPressed {
		c.changeState("jumping")
		c.VelY = c.JumpPower
		c.JumpTimer = 0
	}
}

func (c *Controller) Jumping() {
	// Skidding'den geldiyse özel durum
	if c.previousState == "skidding" {
		if !c.IsJumpKeyPressed && c.JumpTimer < c.JumpReleaseTimer {
			c.VelY = c.ShortJumpVelocity * 0.7 // Kısa zıplama gücünü azalt
			c.JumpTimer = c.JumpHoldTime
		} else if c.IsJumpKeyPressed && c.JumpTimer < c.JumpHoldTime {
			c.VelY += c.JumpBoost * 0.7 // Boost gücünü azalt
			c.JumpTimer++
		} else if c.VelY >= 0 {
			c.changeState("falling")
		}
	} else {
		// Normal jumping mantığı aynen devam eder
		if !c.IsJumpKeyPressed && c.JumpTimer < c.JumpReleaseTimer {
			c.VelY = c.ShortJumpVelocity
			c.JumpTimer = c.JumpHoldTime
		} else if c.IsJumpKeyPressed && c.JumpTimer < c.JumpHoldTime {
			speedFactor := (c.HorizontalVelocity / c.MaxRunSpeed) * c.SpeedJumpFactor
			c.VelY += c.JumpBoost * (1 + speedFactor)
			c.JumpTimer++
		} else if c.VelY >= 0 {
			c.changeState("falling")
		}
	}

	// Yatay hareket kontrolü
	if c.InputAxis.X < 0 && c.VelX > 0 {
		c.VelX -= c.Deceleration
	} else if c.InputAxis.X > 0 && c.VelX < 0 {
		c.VelX += c.Deceleration
	}
}

func (c *Controller) Running() {
	c.AnimPlayer.Animations["walkRight"].FPS = mathutil.MapRange(c.HorizontalVelocity, 0, c.MaxRunSpeed, 4, 23)

	// Kayma durumu kontrolü
	if c.IsSkidding {
		c.changeState("skidding")
		return
	}

	if c.VelY > 0 && !c.IsOnFloor {
		c.changeState("falling")
	}

	if c.IsJumpKeyJustPressed {
		c.changeState("jumping")
		if c.HorizontalVelocity > c.MinSpeedThresForJumpBoostMultiplier {
			c.VelY = c.JumpPower * c.JumpBoostMultiplier
		} else {
			c.VelY = c.JumpPower
		}
		c.JumpTimer = 0
	} else if c.HorizontalVelocity < 0.01 {
		c.changeState("idle")
	} else if c.HorizontalVelocity <= c.MaxWalkSpeed {
		c.changeState("walking")
	}
}

func (c *Controller) Walking() {
	c.AnimPlayer.Animations["walkRight"].FPS = mathutil.MapRange(c.HorizontalVelocity, 0, c.MaxRunSpeed, 4, 23)

	// Kayma durumu kontrolü
	if c.IsSkidding {
		c.changeState("skidding")
		return
	}

	if c.VelY > 0 && !c.IsOnFloor {
		c.changeState("falling")
	}

	if c.IsJumpKeyJustPressed {
		c.changeState("jumping")
		if c.HorizontalVelocity > c.MinSpeedThresForJumpBoostMultiplier {
			c.VelY = c.JumpPower * c.JumpBoostMultiplier
		} else {
			c.VelY = c.JumpPower
		}
		c.JumpTimer = 0
	} else if c.HorizontalVelocity <= 0 {
		c.changeState("idle")
	} else if c.HorizontalVelocity > c.MaxWalkSpeed {
		c.changeState("running")
	}
}

func (c *Controller) Idle() {
	if c.InputAxisLast.Y == -1 {
		c.AnimPlayer.SetStateAndReset("idleUp")
	} else if c.InputAxisLast.Y == 1 {
		c.AnimPlayer.SetStateAndReset("idleDown")
	}
	if c.IsJumpKeyJustPressed {
		c.changeState("jumping")
		c.VelY = c.JumpPower
		c.JumpTimer = 0
	} else if c.IsOnFloor && c.HorizontalVelocity > 0.01 {
		if c.HorizontalVelocity > c.MaxWalkSpeed {
			c.changeState("running")
		} else {
			c.changeState("walking")
		}
	} else if !c.IsOnFloor && c.VelY > 0.01 {
		c.changeState("falling")
	} else if c.IsBreakKeyPressed && IsRayHit {
		c.changeState("attacking")
	}
}

func (c *Controller) UpdateState() {
	switch c.CurrentState {
	case "idle":
		c.Idle()
	case "walking":
		c.Walking()
	case "running":
		c.Running()
	case "jumping":
		c.Jumping()
	case "falling":
		c.Falling()
	case "attacking":
		c.Attacking()
	case "skidding":
		c.Skidding()
	}
}

// func (c *Controller) exitRunning()  {}
// func (c *Controller) exitJumping()  {}
// func (c *Controller) exitFalling()  {}

func (c *Controller) enterWalking() {
	c.AnimPlayer.SetStateAndReset("walkRight")
}
func (c *Controller) enterRunning() {
	c.AnimPlayer.SetStateAndReset("walkRight")
}

func (c *Controller) enterIdle() {
	c.AnimPlayer.SetStateAndReset("idleRight")
}
func (c *Controller) enterAttacking() {

	if c.InputAxisLast.X == 1 {
		c.AnimPlayer.SetStateAndReset("attackRight")
	} else if c.InputAxisLast.X == -1 {
		c.AnimPlayer.SetStateAndReset("attackRight")
		c.DOP.FlipX = true
	} else if c.InputAxisLast.Y == 1 {
		c.AnimPlayer.SetStateAndReset("attack")
	} else if c.InputAxisLast.Y == -1 {
		c.AnimPlayer.SetStateAndReset("attackUp")
	}
}
func (c *Controller) exitAttacking() {
	blockHealth = 0
}

func (c *Controller) enterJumping() {
	c.AnimPlayer.SetStateAndReset("jump")
}

func (c *Controller) enterFalling() {
	c.AnimPlayer.SetStateAndReset("jump")
}

func (c *Controller) enterSkidding() {
	c.AnimPlayer.SetStateAndReset("skidding")
}

func (c *Controller) changeState(newState string) {
	if c.CurrentState == newState {
		return
	}

	// Mevcut durumdan çık
	switch c.CurrentState {
	case "attacking":
		c.exitAttacking()
		// case "idle":
		// c.exitIdle()
		// case "walking":
		// 	c.exitWalking()
		// case "running":
		// 	c.exitRunning()
		// case "jumping":
		// 	c.exitJumping()
		// case "falling":
		// 	c.exitFalling()
	}

	c.previousState = c.CurrentState
	c.CurrentState = newState

	// Yeni duruma gir
	switch newState {
	case "idle":
		c.enterIdle()
	case "attacking":
		c.enterAttacking()
	case "walking":
		c.enterWalking()
	case "running":
		c.enterRunning()
	case "jumping":
		c.enterJumping()
	case "falling":
		c.enterFalling()
	case "skidding":
		c.enterSkidding()
	}
}
