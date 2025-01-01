package system

import (
	"kar"
	"kar/arc"
	"kar/tilemap"

	"github.com/mlange-42/arche/ecs"
	"github.com/setanarut/tilecollider"
)

var (
	TempFallingY float64
	PlayerEntity ecs.Entity
	CTRL         *Controller
	Map          *tilemap.TileMap
	Collider     *tilecollider.Collider[uint16]
	ToSpawn      = []arc.SpawnData{}
)

func AppendToSpawnList(x, y float64, id uint16, durability int) {
	ToSpawn = append(ToSpawn, arc.SpawnData{
		X:          x - 4*kar.ItemScale,
		Y:          y - 4*kar.ItemScale,
		Id:         id,
		Durability: durability,
	})
}

func (s *Spawn) Init() {
	Map = tilemap.MakeTileMap(512, 512, 40, 40)
	tilemap.Generate(Map)
	Collider = tilecollider.NewCollider(Map.Grid, Map.TileW, Map.TileH)
	CTRL = NewController(0, 10, Collider)
	CTRL.Collider = Collider
	CTRL.SetScale(2)
	CTRL.SkiddingJumpEnabled = true
	SpawnX, SpawnY := Map.FindSpawnPosition()
	p := Map.WorldToTile(SpawnX, SpawnY)
	px, py := Map.TileToWorld(p)
	kar.Camera.TopLeftX = px - kar.Camera.Width()/2
	kar.Camera.TopLeftY = py - kar.Camera.Height()/2

	PlayerEntity := arc.SpawnPlayer(SpawnX, SpawnY)
	anim, hlt, dop, rect, inv := arc.MapPlayer.Get(PlayerEntity)
	CTRL.AnimPlayer = anim
	CTRL.Rect = rect
	CTRL.Health = hlt
	CTRL.DOP = dop
	CTRL.Inventory = inv
	CTRL.EnterFalling()

}
func (s *Spawn) Update() {
	// Spawn item
	for _, spawnData := range ToSpawn {
		arc.SpawnItem(spawnData)
	}
	ToSpawn = ToSpawn[:0]
}
func (s *Spawn) Draw() {
	kar.Screen.Fill(kar.BackgroundColor)
}

type Spawn struct{}
