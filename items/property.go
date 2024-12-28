package items

// Item Bitmask tag
const (
	None  uint = 0
	Block uint = 1 << iota
	OreBlock
	Unbreakable
	Harvestable
	DropItem
	Item
	RawOre
	Tool
	Weapon
	Food
	MaterialWooden
	MaterialGold
	MaterialStone
	MaterialIron
	MaterialDiamond
	ToolHand
	ToolAxe
	ToolPickaxe
	ToolShovel
	ToolBucket
)

const All = ^uint(0)

type ItemProperty struct {
	DisplayName string
	DropID      uint16
	Stackable   uint8
	MaxHealth   float64
	Tags        uint
	BestToolTag uint
}

var Property = map[uint16]ItemProperty{
	Air: {
		DisplayName: "",
		DropID:      Air,
		Stackable:   0,
		MaxHealth:   10000,
		Tags:        None | Unbreakable,
	},

	Bedrock: {
		DisplayName: "Bedrock",
		DropID:      0,
		Stackable:   0,
		MaxHealth:   10000,
		Tags:        Block | Unbreakable,
	},
	Bread: {
		DisplayName: "Bread",
		DropID:      0,
		Stackable:   64,
		MaxHealth:   100,
		Tags:        Item | Food,
	},

	Bucket: {
		DisplayName: "Bucket",
		DropID:      0,
		Stackable:   1,
		MaxHealth:   100,
		Tags:        Item | Tool,
	},

	Coal: {
		DisplayName: "Coal",
		Stackable:   64,
		MaxHealth:   1,
		Tags:        Item | DropItem | RawOre,
	},

	CoalOre: {
		DisplayName: "Coal Ore",
		DropID:      Coal,
		Stackable:   64,
		MaxHealth:   100,
		Tags:        Block | OreBlock,
		BestToolTag: ToolPickaxe,
	},

	CraftingTable: {
		DisplayName: "Crafting Table",
		DropID:      CraftingTable,
		Stackable:   1,
		MaxHealth:   100,
		Tags:        Block,
		BestToolTag: ToolAxe,
	},

	Diamond: {
		DisplayName: "Diamond",
		Stackable:   64,
		MaxHealth:   1,
		Tags:        Item | DropItem | RawOre,
	},
	DiamondAxe: {
		DisplayName: "Diamond Axe",
		Stackable:   1,
		MaxHealth:   1,
		Tags:        Item | Tool | ToolAxe | Weapon | MaterialDiamond,
	},

	DiamondOre: {
		DisplayName: "Diamond Ore",
		DropID:      Diamond,
		Stackable:   64,
		MaxHealth:   100,
		Tags:        Block | OreBlock,
		BestToolTag: ToolPickaxe,
	},
	DiamondPickaxe: {
		DisplayName: "Diamond Pickaxe",
		Stackable:   1,
		MaxHealth:   1,
		Tags:        Item | Tool | ToolPickaxe | MaterialDiamond,
	},
	DiamondShovel: {
		DisplayName: "Diamond Hoe",
		Stackable:   1,
		MaxHealth:   1,
		Tags:        Item | Tool | ToolShovel | MaterialDiamond,
	},

	Dirt: {
		DisplayName: "Dirt",
		DropID:      Dirt,
		Stackable:   64,
		MaxHealth:   100,
		Tags:        Block,
		BestToolTag: ToolShovel,
	},

	Furnace: {
		DisplayName: "Furnace",
		DropID:      Furnace,
		Stackable:   1,
		MaxHealth:   100,
		Tags:        Block,
		BestToolTag: ToolPickaxe,
	},
	FurnaceOn: {
		DisplayName: "Furnace On",
		DropID:      Furnace,
		Stackable:   1,
		MaxHealth:   100,
		BestToolTag: ToolPickaxe,
	},
	GoldIngot: {
		DisplayName: "Gold Ingot",
		Stackable:   64,
		MaxHealth:   1,
		Tags:        Item,
	},
	GoldOre: {
		DisplayName: "Gold Ore",
		DropID:      RawGold,
		Stackable:   64,
		MaxHealth:   100,
		Tags:        Block | OreBlock,
		BestToolTag: ToolPickaxe,
	},

	GrassBlock: {
		DisplayName: "Grass Block",
		DropID:      Dirt,
		Stackable:   64,
		MaxHealth:   100,
		Tags:        Block,
		BestToolTag: ToolShovel,
	},
	GrassBlockSnow: {
		DisplayName: "Grass Block Snow",
		DropID:      Dirt,
		Stackable:   64,
		MaxHealth:   100,
		Tags:        Block,
		BestToolTag: ToolShovel,
	},

	IronAxe: {
		DisplayName: "Iron Axe",
		Stackable:   1,
		MaxHealth:   1,
		Tags:        Item | Tool | ToolAxe | MaterialIron,
	},

	IronIngot: {
		DisplayName: "Iron Ingot",
		Stackable:   64,
		MaxHealth:   1,
		Tags:        Item,
	},
	IronOre: {
		DisplayName: "Iron Ore",
		DropID:      RawIron,
		Stackable:   64,
		MaxHealth:   100,
		Tags:        Block | OreBlock,
		BestToolTag: ToolPickaxe,
	},

	IronPickaxe: {
		DisplayName: "Iron Pickaxe",
		Stackable:   1,
		MaxHealth:   1,
		Tags:        Item | Tool | ToolPickaxe | MaterialIron,
	},

	IronShovel: {
		DisplayName: "Iron Shovel",
		Stackable:   1,
		MaxHealth:   1,
		Tags:        Item | Tool | ToolShovel | MaterialIron,
	},

	OakLeaves: {
		DisplayName: "Oak Leaves",
		DropID:      OakLeaves,
		Stackable:   64,
		MaxHealth:   100,
		Tags:        Block,
		BestToolTag: ToolAxe,
	},
	OakLog: {
		DisplayName: "Oak Log",
		DropID:      OakLog,
		Stackable:   64,
		MaxHealth:   100,
		Tags:        Block,
		BestToolTag: ToolAxe,
	},
	OakPlanks: {
		DisplayName: "Oak Planks",
		DropID:      OakPlanks,
		Stackable:   64,
		MaxHealth:   100,
		Tags:        Block,
		BestToolTag: ToolAxe,
	},
	OakSapling: {
		DisplayName: "Oak Sapling",
		DropID:      OakSapling,
		Stackable:   64,
		MaxHealth:   1,
		Tags:        Block | Item,
		BestToolTag: ToolAxe,
	},
	Obsidian: {
		DisplayName: "Obsidian",
		DropID:      Obsidian,
		Stackable:   64,
		MaxHealth:   20,
		Tags:        Block,
		BestToolTag: ToolPickaxe,
	},

	RawGold: {
		DisplayName: "Raw Gold",
		Stackable:   64,
		MaxHealth:   1,
		Tags:        Item | DropItem | RawOre,
	},
	RawIron: {
		DisplayName: "Raw Iron",
		Stackable:   64,
		MaxHealth:   1,
		Tags:        Item | DropItem | RawOre,
	},

	Sand: {
		DisplayName: "Sand",
		DropID:      Sand,
		Stackable:   64,
		MaxHealth:   100,
		Tags:        Block,
		BestToolTag: ToolShovel,
	},

	SmoothStone: {
		DisplayName: "Smooth Stone",
		DropID:      SmoothStone,
		Stackable:   64,
		MaxHealth:   100,
		Tags:        Block,
	},
	Snow: {
		DisplayName: "Snow",
		DropID:      Dirt,
		Stackable:   64,
		MaxHealth:   5,
		Tags:        Block,
		BestToolTag: ToolShovel,
	},
	Snowball: {
		DisplayName: "Snowball",
		DropID:      Snowball,
		Stackable:   64,
		MaxHealth:   100,
		Tags:        Item,
	},
	Stick: {
		DisplayName: "Stick",
		Stackable:   64,
		MaxHealth:   1,
		Tags:        Item,
	},

	Stone: {
		DisplayName: "Stone",
		DropID:      Stone,
		Stackable:   64,
		MaxHealth:   100,
		Tags:        Block,
		BestToolTag: ToolPickaxe,
	},
	StoneAxe: {
		DisplayName: "Stone Axe",
		Stackable:   1,
		MaxHealth:   1,
		Tags:        Item | Tool | ToolAxe | MaterialStone,
	},
	StoneBricks: {
		DisplayName: "Stone Bricks",
		DropID:      StoneBricks,
		Stackable:   64,
		MaxHealth:   100,
		Tags:        Block,
		BestToolTag: ToolPickaxe,
	},

	StonePickaxe: {
		DisplayName: "Stone Pickaxe",
		Stackable:   1,
		MaxHealth:   1,
		Tags:        Item | Tool | ToolPickaxe | MaterialStone,
	},
	StoneShovel: {
		DisplayName: "Stone Shovel",
		Stackable:   1,
		MaxHealth:   1,
		Tags:        Item | Tool | ToolShovel | MaterialStone,
	},

	Tnt: {
		DisplayName: "TNT",
		DropID:      Tnt,
		Stackable:   64,
		MaxHealth:   100,
		Tags:        Block,
		BestToolTag: All,
	},
	Torch: {
		DisplayName: "Torch",
		Stackable:   64,
		MaxHealth:   100,
		Tags:        Item | Block,
		BestToolTag: All,
	},
	WaterBucket: {
		DisplayName: "Water Bucket",
		Stackable:   1,
		MaxHealth:   100,
		Tags:        Item | Tool | ToolBucket | MaterialIron,
	},

	WoodenAxe: {
		DisplayName: "Wooden Axe",
		Stackable:   1,
		MaxHealth:   1,
		Tags:        Item | Tool,
	},

	WoodenPickaxe: {
		DisplayName: "Wooden Pickaxe",
		Stackable:   1,
		MaxHealth:   1,
		Tags:        Item | Tool,
	},
	WoodenShovel: {
		DisplayName: "Wooden Shovel",
		Stackable:   1,
		MaxHealth:   1,
		Tags:        Item | Tool,
	},
}
