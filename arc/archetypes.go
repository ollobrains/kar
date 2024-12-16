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
	MapItem      = gn.NewMap4[ItemID, Health, DrawOptions, Rect](&kar.WorldECS)
	MapPlayer    = gn.NewMap5[
		Health,
		DrawOptions,
		anim.AnimationPlayer,
		Rect,
		Inventory](&kar.WorldECS)
)

// Query Filters
var (
	FilterAnimPlayer = gn.NewFilter1[anim.AnimationPlayer]()
	FilterItem       = gn.NewFilter4[ItemID, Health, DrawOptions, Rect]()
	FilterPlayer     = gn.NewFilter5[
		Health,
		DrawOptions,
		anim.AnimationPlayer,
		Rect,
		Inventory]()
)

func init() {
	FilterAnimPlayer.Register(&kar.WorldECS)
	FilterItem.Register(&kar.WorldECS)
	FilterPlayer.Register(&kar.WorldECS)
}

func SpawnPlayer(x, y float64) ecs.Entity {
	h := &Health{Health: 100, MaxHealth: 100}
	a := anim.NewAnimationPlayer(res.PlayerAtlas)
	a.NewAnimationState("idleRight", 0, 0, 16, 16, 1, false, false).FPS = 1
	a.NewAnimationState("walkRight", 16, 0, 16, 16, 4, false, false)
	a.NewAnimationState("jump", 16*5, 0, 16, 16, 1, false, false)
	a.NewAnimationState("skidding", 16*6, 0, 16, 16, 1, false, false)
	a.NewAnimationState("attack", 16*7, 0, 16, 16, 2, false, false).FPS = 8
	a.NewAnimationState("attackRight", 16*9, 0, 16, 16, 2, false, false).FPS = 8
	a.NewAnimationState("attackUp", 16*11, 0, 16, 16, 2, false, false).FPS = 8
	a.SetState("idleRight")
	d := &DrawOptions{Scale: 3}
	r := &Rect{X: x, Y: y, W: 16 * 3, H: 16 * 3}
	return MapPlayer.NewWith(h, d, a, r, NewInventory())
}
func SpawnItem(x, y float64, id uint16) ecs.Entity {
	i := &ItemID{ID: id}
	h := &Health{Health: 100, MaxHealth: 100}
	d := &DrawOptions{Scale: 2}
	r := &Rect{X: x, Y: y, W: 16 * 2, H: 16 * 2}
	return MapItem.NewWith(i, h, d, r)
}
