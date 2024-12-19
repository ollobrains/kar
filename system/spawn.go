package system

import (
	"kar"
	"kar/arc"
	"kar/items"
	"kar/tilemap"

	"github.com/mlange-42/arche/ecs"
	"github.com/setanarut/tilecollider"
)

var tm = [][]uint16{
	{items.Stone, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, items.Stone},
	{items.Stone, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, items.Stone},
	{items.Stone, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, items.Stone},
	{items.Stone, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, items.Stone},
	{items.Stone, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, items.Stone},
	{items.Stone, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, items.Stone},
	{items.Stone, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, items.Stone},
	{items.Stone, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, items.Stone},
	{items.Stone, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, items.Stone},
	{items.Stone, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, items.Stone},
	{items.Stone, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, items.Stone},
	{items.Stone, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, items.Stone},
	{items.Stone, items.Stone, items.Stone, items.Stone, items.Stone, items.Stone, items.Stone, items.Stone, items.Stone, items.Stone, items.Stone, items.Stone, items.Stone, items.Stone, items.Stone, items.Stone, items.Stone, items.Stone, items.Stone, items.Stone, items.Stone, items.Stone, items.Stone, items.Stone, items.Stone, items.Stone, items.Stone, items.Stone, items.Stone}}

var (
	PlayerEntity     ecs.Entity
	PlayerInventory  *arc.Inventory
	PlayerController = NewController(0, 10, Collider)
	Map              = tilemap.NewTileMap(tm, 48, 48)
	Collider         = tilecollider.NewCollider(Map.Grid, Map.TileW, Map.TileH)
	ToSpawn          = []SpawnData{}
)

// SpawnData is a helper for delaying spawn events
type SpawnData struct {
	X, Y float64
	Id   uint16
}

func AppendToSpawnList(x, y float64, id uint16) {
	ToSpawn = append(ToSpawn, SpawnData{
		X:  x - 8*kar.ItemScale,
		Y:  y - 8*kar.ItemScale,
		Id: id,
	})
}

func init() {
	PlayerController.Collider = Collider
	PlayerController.SetScale(3)
	PlayerController.SkiddingJumpEnabled = true
}

func (s *Spawn) Init() {
	PlayerEntity = arc.SpawnPlayer(512, 400)
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
