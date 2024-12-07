package system

import (
	"kar/engine/mathutil"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/setanarut/anim"
	"github.com/setanarut/tilecollider"
)

type Controller struct {
	CurrentState string
	Collider     *tilecollider.Collider[uint16]
	AnimPlayer   *anim.AnimationPlayer

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

	SkiddingJumpEnabled bool

	// Input durumları
	IsLeftKeyPressed     bool
	IsRightKeyPressed    bool
	IsJumpKeyPressed     bool
	IsJumpKeyJustPressed bool
	IsRunKeyPressed      bool
	WalkAcceleration     float64
	WalkDeceleration     float64
	RunAcceleration      float64
	RunDeceleration      float64

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
	// c.JumpHoldTime *= s
	c.JumpBoost *= s
	c.MinSpeedThresForJumpBoostMultiplier *= s
	// c.JumpBoostMultiplier *= s
	// c.SpeedJumpFactor *= s
	c.ShortJumpVelocity *= s
	// c.JumpReleaseTimer *= s
	c.MaxWalkSpeed *= s
	c.MaxRunSpeed *= s
	c.WalkAcceleration *= s
	c.WalkDeceleration *= s
	c.RunAcceleration *= s
	c.RunDeceleration *= s
}
func (c *Controller) UpdateInput() {
	c.IsRunKeyPressed = ebiten.IsKeyPressed(ebiten.KeyShift)
	c.IsLeftKeyPressed = ebiten.IsKeyPressed(ebiten.KeyA)
	c.IsRightKeyPressed = ebiten.IsKeyPressed(ebiten.KeyD)
	c.IsJumpKeyPressed = ebiten.IsKeyPressed(ebiten.KeySpace)
	c.IsJumpKeyJustPressed = inpututil.IsKeyJustPressed(ebiten.KeySpace)
}

func (c *Controller) UpdatePhysics(x, y, w, h float64) (dx, dy float64) {
	c.IsSkidding = (c.VelX > 0 && c.IsLeftKeyPressed) || (c.VelX < 0 && c.IsRightKeyPressed)
	maxSpeed := c.MaxWalkSpeed
	currentAccel := c.WalkAcceleration
	currentDecel := c.WalkDeceleration
	c.HorizontalVelocity = math.Abs(c.VelX)

	if !c.IsSkidding {
		// Koşma durumunda maksimum hızı ve ivmelenmeyi ayarla
		if c.IsRunKeyPressed {
			maxSpeed = c.MaxRunSpeed
			currentAccel = c.RunAcceleration
			currentDecel = c.RunDeceleration
		} else if c.HorizontalVelocity > c.MaxWalkSpeed {
			// Koşma tuşu bırakıldığında ve hız yürüme hızından fazlaysa
			// RunDeceleration kullan
			currentDecel = c.RunDeceleration
		}
	}

	if c.IsRightKeyPressed {
		if c.VelX > maxSpeed {
			// Hız maksimumun üzerindeyse yavaşla
			c.VelX = max(maxSpeed, c.VelX-currentDecel)
		} else {
			// Normal ivmelenme
			c.VelX = min(maxSpeed, c.VelX+currentAccel)
		}
	} else if c.IsLeftKeyPressed {
		if c.VelX < -maxSpeed {
			// Hız maksimumun üzerindeyse yavaşla
			c.VelX = min(-maxSpeed, c.VelX+currentDecel)
		} else {
			// Normal ivmelenme
			c.VelX = max(-maxSpeed, c.VelX-currentAccel)
		}
	} else {
		// Hareket tuşları basılı değilse
		if c.VelX > 0 {
			c.VelX = max(0, c.VelX-currentDecel)
		} else if c.VelX < 0 {
			c.VelX = min(0, c.VelX+currentDecel)
		}
	}

	c.VelY += c.Gravity
	c.VelY = min(c.MaxFallSpeed, c.VelY)

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

func (c *Controller) UpdateState() {
	switch c.CurrentState {
	case "idle":
		if c.IsJumpKeyJustPressed {
			c.changeState("jumping")
			c.VelY = c.JumpPower
			c.JumpTimer = 0
		} else if c.HorizontalVelocity > 0.01 {
			if c.HorizontalVelocity > c.MaxWalkSpeed {
				c.changeState("running")
			} else {
				c.changeState("walking")
			}
		}

	case "walking":
		c.AnimPlayer.Animations["walkRight"].FPS = mathutil.MapRange(c.HorizontalVelocity, 0, c.MaxRunSpeed, 4, 23)

		// Kayma durumu kontrolü
		if c.IsSkidding {
			c.changeState("skidding")
			break
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

	case "running":
		c.AnimPlayer.Animations["walkRight"].FPS = mathutil.MapRange(c.HorizontalVelocity, 0, c.MaxRunSpeed, 4, 23)

		// Kayma durumu kontrolü
		if c.IsSkidding {
			c.changeState("skidding")
			break
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

	case "jumping":
		if !c.IsJumpKeyPressed && c.JumpTimer < c.JumpReleaseTimer {
			c.VelY = c.ShortJumpVelocity
			c.JumpTimer = c.JumpHoldTime // Zıplama süresini bitir
		} else if c.IsJumpKeyPressed && c.JumpTimer < c.JumpHoldTime {
			speedFactor := (c.HorizontalVelocity / c.MaxRunSpeed) * c.SpeedJumpFactor
			c.VelY += c.JumpBoost * (1 + speedFactor)
			c.JumpTimer++
		} else if c.VelY >= 0 {
			c.changeState("falling")
		}

		if c.IsLeftKeyPressed && c.VelX > 0 {

			c.VelX -= c.Deceleration
		} else if c.IsRightKeyPressed && c.VelX < 0 {
			c.VelX += c.Deceleration
		}

	case "falling":
		if c.IsOnFloor {
			if c.HorizontalVelocity <= 0 {
				c.changeState("idle")
			} else if c.IsRunKeyPressed {
				c.changeState("running")
			} else {
				c.changeState("walking")
			}
		}

	case "skidding":
		if c.HorizontalVelocity < 0.01 {
			c.changeState("idle")
		} else if !c.IsSkidding {
			if c.HorizontalVelocity > c.MaxWalkSpeed {
				c.changeState("running")
			} else {
				c.changeState("walking")
			}
		}

		if c.SkiddingJumpEnabled {
			if c.IsJumpKeyJustPressed {
				c.changeState("jumping")
				if c.HorizontalVelocity > c.MinSpeedThresForJumpBoostMultiplier {
					c.VelY = c.JumpPower * c.JumpBoostMultiplier
				} else {
					c.VelY = c.JumpPower
				}
				c.JumpTimer = 0
			}
		}
	}

	// Düşme durumunda da IsOnFloor'u false yap
	if c.VelY > 0 {
		c.IsOnFloor = false
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

	// // Mevcut durumdan çık
	// switch c.CurrentState {
	// case "idle":
	// 	c.exitIdle()
	// case "walking":
	// 	c.exitWalking()
	// case "running":
	// 	c.exitRunning()
	// case "jumping":
	// 	c.exitJumping()
	// case "falling":
	// 	c.exitFalling()
	// }

	c.previousState = c.CurrentState
	c.CurrentState = newState

	// Yeni duruma gir
	switch newState {
	case "idle":
		c.enterIdle()
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
