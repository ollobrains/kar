package system

import (
	"kar"
	"kar/arc"
	"kar/tilemap"

	"github.com/mlange-42/arche/ecs"
	"github.com/setanarut/tilecollider"
)

var (
	PlayerEntity     ecs.Entity
	PlayerInventory  *arc.Inventory
	PlayerController = NewController(0, 10, Collider)
	Map              = tilemap.MakeTileMap(512, 512, 40, 40)

	Collider = tilecollider.NewCollider(Map.Grid, Map.TileW, Map.TileH)
	ToSpawn  = []arc.SpawnData{}
)

func AppendToSpawnList(x, y float64, id uint16, durability float64) {
	ToSpawn = append(ToSpawn, arc.SpawnData{
		X:          x - 4*kar.ItemScale,
		Y:          y - 4*kar.ItemScale,
		Id:         id,
		Durability: durability,
	})
}

func (s *Spawn) Init() {
	tilemap.Generate(Map)

	PlayerController.Collider = Collider
	PlayerController.SetScale(2)
	PlayerController.SkiddingJumpEnabled = true
	x, y := Map.FindSpawnPosition()
	p := Map.WorldToTile(x, y)
	px, py := Map.TileToWorld(p)
	kar.Camera.TopLeftX = px - kar.Camera.Width()/2
	kar.Camera.TopLeftY = py - kar.Camera.Height()/2
	PlayerEntity = arc.SpawnPlayer(x, y)
	PlayerInventory = arc.MapInventory.Get(PlayerEntity)
	PlayerInventory.RandomFillAllSlots()
	PlayerInventory.ClearSlot(0)
	PlayerInventory.ClearSlot(1)
	PlayerInventory.ClearSlot(4)
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
