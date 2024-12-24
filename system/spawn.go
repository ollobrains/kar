package system

import (
	"kar"
	"kar/arc"
	"kar/tilemap"

	"github.com/mlange-42/arche/ecs"
	"github.com/setanarut/kamera/v2"
	"github.com/setanarut/tilecollider"
)

var (
	PlayerEntity     ecs.Entity
	PlayerInventory  *arc.Inventory
	PlayerController = NewController(0, 10, Collider)
	Map              = tilemap.MakeTileMap(512, 512, 32, 32)

	Collider = tilecollider.NewCollider(Map.Grid, Map.TileW, Map.TileH)
	ToSpawn  = []SpawnData{}
)

// SpawnData is a helper for delaying spawn events
type SpawnData struct {
	X, Y float64
	Id   uint16
}

func AppendToSpawnList(x, y float64, id uint16) {
	ToSpawn = append(ToSpawn, SpawnData{
		X:  x - 4*kar.ItemScale,
		Y:  y - 4*kar.ItemScale,
		Id: id,
	})
}

func (s *Spawn) Init() {
	tilemap.Generate(Map)

	PlayerController.Collider = Collider
	PlayerController.SetScale(2)
	PlayerController.SkiddingJumpEnabled = true
	x, y := Map.FindSpawnPosition()
	PlayerEntity = arc.SpawnPlayer(x, y)
	kar.Camera.LookAt(x, y)
	kar.Camera.SmoothType = kamera.None
	PlayerInventory = arc.MapInventory.Get(PlayerEntity)
	PlayerInventory.RandomFillAllSlots()
}
func (s *Spawn) Update() {
	// Spawn item
	for _, spawnData := range ToSpawn {
		arc.SpawnItem(spawnData.X, spawnData.Y, spawnData.Id)
	}
	ToSpawn = ToSpawn[:0]
}
func (s *Spawn) Draw() {
	kar.Screen.Fill(kar.BackgroundColor)
}

type Spawn struct{}
