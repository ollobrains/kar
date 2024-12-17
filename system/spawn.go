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

// Reproduction is a helper for delaying reproduction events
type SpawnData struct {
	X, Y float64
	Id   uint16
}

var PlayerEntity ecs.Entity
var PlayerInventory *arc.Inventory
var PlayerController = NewController(0, 10, Collider)
var Map = tilemap.NewTileMap(tm, 48, 48)
var Collider = tilecollider.NewCollider(Map.Grid, Map.TileW, Map.TileH)

var toSpawn = []SpawnData{}

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
	for _, spawnData := range toSpawn {
		arc.SpawnItem(spawnData.X, spawnData.Y, spawnData.Id)
	}
	toSpawn = toSpawn[:0]
}
func (s *Spawn) Draw() {
	kar.Screen.Fill(kar.BackgroundColor)
}

type Spawn struct{}
