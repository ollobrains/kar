package system

import (
	"log"

	"kar"
	"kar/arc"
	"kar/engine/mathutil"
	"kar/res"

	// Possibly you’re using a camera transform library from Ebiten or a custom camera
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"
)

// RenderSystem handles camera movement, tile rendering, entity rendering,
// debug bounding boxes, and optionally, fall damage logic.
type RenderSystem struct {
	// If you want to configure fall damage thresholds:
	FallDamageThreshold float64
	FallDamageMultiplier float64
}

// Init can set up any resources needed. If none, we keep it empty.
func (rn *RenderSystem) Init() {
	// Example: set default values if not manually assigned
	if rn.FallDamageThreshold == 0 {
		rn.FallDamageThreshold = 10.0
	}
	if rn.FallDamageMultiplier == 0 {
		rn.FallDamageMultiplier = 1.25
	}
}

// Update handles camera scrolling, player animations, and calls to your ECS queries.
func (rn *RenderSystem) Update() {
	// Example: camera logic
	if playerCenterX < kar.Camera.TopLeftX {
		kar.Camera.TopLeftX -= kar.Camera.Width()
	}
	if playerCenterX > kar.Camera.Right() {
		kar.Camera.TopLeftX += kar.Camera.Width()
	}
	if playerCenterY < kar.Camera.TopLeftY {
		kar.Camera.TopLeftY -= kar.Camera.Height()
	}
	if playerCenterY > kar.Camera.Bottom() {
		kar.Camera.TopLeftY += kar.Camera.Height()
	}

	// Animate players (or other entities) from arc.FilterAnimPlayer
	q := arc.FilterAnimPlayer.Query(&kar.WorldECS)
	for q.Next() {
		animPlayer := q.Get()
		animPlayer.Update()
	}

	// Optional: call a function to check fall damage on any relevant entities
	rn.checkFallDamage()
}

// Draw handles rendering tilemaps, bounding boxes, items, player sprites, etc.
func (rn *RenderSystem) Draw() {
	// 1. Draw the tilemap. This example references a "Map" object not shown in your snippet.
	camMin := Map.WorldToTile(kar.Camera.TopLeft())
	camMin.X = min(max(camMin.X, 0), Map.W)
	camMin.Y = min(max(camMin.Y, 0), Map.H)
	camMaxX := min(max(camMin.X+kar.RenderArea.X, 0), Map.W)
	camMaxY := min(max(camMin.Y+kar.RenderArea.Y, 0), Map.H)

	for y := camMin.Y; y < camMaxY; y++ {
		for x := camMin.X; x < camMaxX; x++ {
			tileID := Map.Grid[y][x]
			if tileID != 0 {
				px, py := float64(x*Map.TileW), float64(y*Map.TileH)

				kar.GlobalDIO.GeoM.Reset()
				kar.GlobalDIO.GeoM.Scale(2, 2)
				kar.GlobalDIO.GeoM.Translate(px, py)

				// Example: if x,y is the "target" block, do some special logic
				if x == targetBlockPos.X && y == targetBlockPos.Y {
					i := mathutil.MapRange(blockHealth, 0, 180, 0, 5)
					if res.BlockCrackFrames[tileID] != nil {
						frameIndex := int(i)
						kar.Camera.DrawWithColorM(
							res.BlockCrackFrames[tileID][frameIndex],
							kar.GlobalColorM,
							kar.GlobalDIO,
							kar.Screen,
						)
					}
				} else {
					kar.Camera.DrawWithColorM(
						res.BlockCrackFrames[tileID][0],
						kar.GlobalColorM,
						kar.GlobalDIO,
						kar.Screen,
					)
				}
			}
		}
	}

	// 2. Draw the highlight on the target block if "IsRayHit" is set.
	if IsRayHit {
		kar.GlobalDIO.GeoM.Reset()
		kar.GlobalDIO.GeoM.Scale(2, 2)
		kar.GlobalDIO.GeoM.Translate(
			float64(targetBlockPos.X*Map.TileW)-2,
			float64(targetBlockPos.Y*Map.TileH)-2,
		)
		kar.Camera.DrawWithColorM(res.SelectionBlock, kar.GlobalColorM, kar.GlobalDIO, kar.Screen)
	}

	// 3. Draw player(s)
	playerQuery := arc.FilterPlayer.Query(&kar.WorldECS)
	for playerQuery.Next() {
		anim, _, dop, rect, _ := playerQuery.Get()
		sclX := dop.Scale

		kar.GlobalDIO.GeoM.Reset()
		if dop.FlipX {
			sclX *= -1
			kar.GlobalDIO.GeoM.Scale(sclX, dop.Scale)
			kar.GlobalDIO.GeoM.Translate(rect.X+rect.W, rect.Y)
		} else {
			kar.GlobalDIO.GeoM.Scale(sclX, dop.Scale)
			kar.GlobalDIO.GeoM.Translate(rect.X, rect.Y)
		}
		kar.Camera.DrawWithColorM(anim.CurrentFrame, kar.GlobalColorM, kar.GlobalDIO, kar.Screen)
	}

	// 4. Debug: draw bounding boxes or "Rect" components, plus a dot for center
	if kar.DrawDebugHitboxesEnabled {
		rectQ := arc.FilterRect.Query(&kar.WorldECS)
		for rectQ.Next() {
			r := rectQ.Get()
			x, y := kar.Camera.ApplyCameraTransformToPoint(r.X, r.Y)
			vector.StrokeRect(
				kar.Screen,
				float32(x),
				float32(y),
				float32(r.W),
				float32(r.H),
				1,
				colornames.Magenta,
				false,
			)
			// draw center dot
			vector.DrawFilledCircle(
				kar.Screen,
				float32(x+(r.W*0.5)),
				float32(y+(r.H*0.5)),
				2,
				colornames.Magenta,
				false,
			)
		}
	}

	// 5. Draw items
	itemQuery := arc.FilterItem.Query(&kar.WorldECS)
	for itemQuery.Next() {
		id, rect, timers, _ := itemQuery.Get()

		kar.GlobalDIO.GeoM.Reset()
		kar.GlobalDIO.GeoM.Scale(kar.ItemScale, kar.ItemScale)

		// If you had a sin array for item bobbing, consider referencing c.sinspace[timers.AnimationIndex] as well
		dy := c.sinspace[timers.AnimationIndex % len(c.sinspace)]
		kar.GlobalDIO.GeoM.Translate(rect.X, rect.Y+dy)

		if res.Icon8[id.ID] == nil {
			log.Fatal("item image not found for ID=", id.ID)
		}

		kar.Camera.DrawWithColorM(res.Icon8[id.ID], kar.GlobalColorM, kar.GlobalDIO, kar.Screen)
	}
}

// checkFallDamage is an example method to detect vertical velocity and apply damage if landing hard.
func (rn *RenderSystem) checkFallDamage() {
	fallQuery := arc.FilterPlayer.Query(&kar.WorldECS)
	for fallQuery.Next() {
		_, collider, dop, rect, stats := fallQuery.Get() // adjust if your ECS data differs

		// If 'stats' has a 'Health' or 'HP' field, we can reduce it if velocity is big.
		// Example: check if collider.VelY is beyond threshold AND we're on the ground or just collided with ground
		if collider.OnGround && collider.VelY > rn.FallDamageThreshold {
			// The difference from threshold could define the actual damage
			var fallSpeedExcess = collider.VelY - rn.FallDamageThreshold
			var damage = int(fallSpeedExcess * rn.FallDamageMultiplier)

			// Subtract HP
			stats.Health -= damage

			// Potential clamp
			if stats.Health < 0 {
				stats.Health = 0
			}
			// Possibly do: if stats.Health <= 0, kill the entity or trigger a "died" event
		}
	}
}
