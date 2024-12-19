package arc

import (
	"kar"
	"kar/res"

	"github.com/mlange-42/arche/ecs"
	gn "github.com/mlange-42/arche/generic"
	"github.com/setanarut/anim"
)

var (
	MapInventory = gn.NewMap1[Inventory](&kar.WorldECS)
	MapRect      = gn.NewMap1[Rect](&kar.WorldECS)
	MapItem      = gn.NewMap4[ItemID, Health, Rect, CollisionCountdown](&kar.WorldECS)
	MapPlayer    = gn.NewMap5[
		anim.AnimationPlayer,
		Health,
		DrawOptions,
		Rect,
		Inventory](&kar.WorldECS)
)

// Query Filters
var (
	FilterCollision  = gn.NewFilter3[Rect, ItemID, CollisionCountdown]()
	FilterRect       = gn.NewFilter1[Rect]()
	FilterAnimPlayer = gn.NewFilter1[anim.AnimationPlayer]()
	FilterItem       = gn.NewFilter4[ItemID, Health, Rect, CollisionCountdown]()
	FilterPlayer     = gn.NewFilter5[
		anim.AnimationPlayer,
		Health,
		DrawOptions,
		Rect,
		Inventory]()
)

func init() {
	FilterRect.Register(&kar.WorldECS)
	FilterCollision.Register(&kar.WorldECS)
	FilterAnimPlayer.Register(&kar.WorldECS)
	FilterItem.Register(&kar.WorldECS)
	FilterPlayer.Register(&kar.WorldECS)
}

func SpawnPlayer(x, y float64) ecs.Entity {
	AP := anim.NewAnimationPlayer(res.PlayerAtlas)
	AP.NewAnimationState("idleRight", 0, 0, 16, 16, 1, false, false).FPS = 1
	AP.NewAnimationState("walkRight", 16, 0, 16, 16, 4, false, false)
	AP.NewAnimationState("jump", 16*5, 0, 16, 16, 1, false, false)
	AP.NewAnimationState("skidding", 16*6, 0, 16, 16, 1, false, false)
	AP.NewAnimationState("attack", 16*7, 0, 16, 16, 2, false, false).FPS = 8
	AP.NewAnimationState("attackRight", 16*9, 0, 16, 16, 2, false, false).FPS = 8
	AP.NewAnimationState("attackUp", 16*11, 0, 16, 16, 2, false, false).FPS = 8
	AP.SetState("idleRight")
	return MapPlayer.NewWith(
		AP,
		&Health{100, 100},
		&DrawOptions{Scale: 3},
		&Rect{x, y, 16 * 3, 16 * 3},
		NewInventory(),
	)
}
func SpawnItem(x, y float64, id uint16) ecs.Entity {
	return MapItem.NewWith(
		&ItemID{id},
		&Health{Health: 100, MaxHealth: 100},
		&Rect{x, y, 16 * kar.ItemScale, 16 * kar.ItemScale},
		&CollisionCountdown{60},
	)
}
