package items

type Category uint

// Item Bitmask Category
const (
	None  Category = 0
	Block Category = 1 << iota
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

const All = Category(^uint(0))

type ItemProperty struct {
	DisplayName string
	DropID      uint16
	Stackable   uint8
	MaxHealth   float64
	Category    Category
	BestTool    Category
}

var Property = map[uint16]ItemProperty{
	Air: {
		DisplayName: "",
		DropID:      Air,
		Stackable:   0,
		MaxHealth:   10000,
		Category:    None | Unbreakable,
	},

	Bedrock: {
		DisplayName: "Bedrock",
		DropID:      0,
		Stackable:   0,
		MaxHealth:   10000,
		Category:    Block | Unbreakable,
	},
	Bread: {
		DisplayName: "Bread",
		DropID:      0,
		Stackable:   64,
		MaxHealth:   100,
		Category:    Item | Food,
	},

	Bucket: {
		DisplayName: "Bucket",
		DropID:      0,
		Stackable:   1,
		MaxHealth:   100,
		Category:    Item | Tool,
	},

	Coal: {
		DisplayName: "Coal",
		Stackable:   64,
		MaxHealth:   1,
		Category:    Item | DropItem | RawOre,
	},

	CoalOre: {
		DisplayName: "Coal Ore",
		DropID:      Coal,
		Stackable:   64,
		MaxHealth:   100,
		Category:    Block | OreBlock,
		BestTool:    ToolPickaxe,
	},

	CraftingTable: {
		DisplayName: "Crafting Table",
		DropID:      CraftingTable,
		Stackable:   1,
		MaxHealth:   100,
		Category:    Block,
		BestTool:    ToolAxe,
	},

	Diamond: {
		DisplayName: "Diamond",
		Stackable:   64,
		MaxHealth:   1,
		Category:    Item | DropItem | RawOre,
	},
	DiamondAxe: {
		DisplayName: "Diamond Axe",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Tool | Weapon,
	},

	DiamondOre: {
		DisplayName: "Diamond Ore",
		DropID:      Diamond,
		Stackable:   64,
		MaxHealth:   100,
		Category:    Block | OreBlock,
		BestTool:    ToolPickaxe,
	},
	DiamondPickaxe: {
		DisplayName: "Diamond Pickaxe",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Tool,
	},
	DiamondShovel: {
		DisplayName: "Diamond Hoe",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Tool,
	},

	Dirt: {
		DisplayName: "Dirt",
		DropID:      Dirt,
		Stackable:   64,
		MaxHealth:   100,
		BestTool:    ToolShovel,
	},

	Furnace: {
		DisplayName: "Furnace",
		DropID:      Furnace,
		Stackable:   1,
		MaxHealth:   100,
		Category:    Block,
		BestTool:    ToolPickaxe,
	},
	FurnaceOn: {
		DisplayName: "Furnace On",
		DropID:      Furnace,
		Stackable:   1,
		MaxHealth:   100,
		BestTool:    ToolPickaxe,
	},
	GoldIngot: {
		DisplayName: "Gold Ingot",
		Stackable:   64,
		MaxHealth:   1,
		Category:    Item,
	},
	GoldOre: {
		DisplayName: "Gold Ore",
		DropID:      RawGold,
		Stackable:   64,
		MaxHealth:   100,
		Category:    Block | OreBlock,
		BestTool:    ToolPickaxe,
	},

	GrassBlock: {
		DisplayName: "Grass Block",
		DropID:      Dirt,
		Stackable:   64,
		MaxHealth:   100,
		Category:    Block,
		BestTool:    ToolShovel,
	},
	GrassBlockSnow: {
		DisplayName: "Grass Block Snow",
		DropID:      Dirt,
		Stackable:   64,
		MaxHealth:   100,
		Category:    Block,
		BestTool:    ToolShovel,
	},

	IronAxe: {
		DisplayName: "Iron Axe",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Tool,
	},

	IronIngot: {
		DisplayName: "Iron Ingot",
		Stackable:   64,
		MaxHealth:   1,
		Category:    Item,
	},
	IronOre: {
		DisplayName: "Iron Ore",
		DropID:      RawIron,
		Stackable:   64,
		MaxHealth:   100,
		Category:    Block | OreBlock,
		BestTool:    ToolPickaxe,
	},
	IronPickaxe: {
		DisplayName: "Iron Pickaxe",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Tool,
	},
	IronShovel: {
		DisplayName: "Iron Shovel",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Tool,
	},

	OakLeaves: {
		DisplayName: "Oak Leaves",
		DropID:      OakLeaves,
		Stackable:   64,
		MaxHealth:   100,
		Category:    Block,
		BestTool:    ToolAxe,
	},
	OakLog: {
		DisplayName: "Oak Log",
		DropID:      OakLog,
		Stackable:   64,
		MaxHealth:   100,
		Category:    Block,
		BestTool:    ToolAxe,
	},
	OakPlanks: {
		DisplayName: "Oak Planks",
		DropID:      OakPlanks,
		Stackable:   64,
		MaxHealth:   100,
		Category:    Block,
		BestTool:    ToolAxe,
	},
	OakSapling: {
		DisplayName: "Oak Sapling",
		DropID:      OakSapling,
		Stackable:   64,
		MaxHealth:   1,
		Category:    Block | Item,
		BestTool:    ToolAxe,
	},
	Obsidian: {
		DisplayName: "Obsidian",
		DropID:      Obsidian,
		Stackable:   64,
		MaxHealth:   20,
		Category:    Block,
		BestTool:    ToolPickaxe,
	},

	RawGold: {
		DisplayName: "Raw Gold",
		Stackable:   64,
		MaxHealth:   1,
		Category:    Item | DropItem | RawOre,
	},
	RawIron: {
		DisplayName: "Raw Iron",
		Stackable:   64,
		MaxHealth:   1,
		Category:    Item | DropItem | RawOre,
	},

	Sand: {
		DisplayName: "Sand",
		DropID:      Sand,
		Stackable:   64,
		MaxHealth:   100,
		Category:    Block,
		BestTool:    ToolShovel,
	},

	SmoothStone: {
		DisplayName: "Smooth Stone",
		DropID:      SmoothStone,
		Stackable:   64,
		MaxHealth:   100,
		Category:    Block,
	},
	Snow: {
		DisplayName: "Snow",
		DropID:      Dirt,
		Stackable:   64,
		MaxHealth:   5,
		Category:    Block,
		BestTool:    ToolShovel,
	},
	Snowball: {
		DisplayName: "Snowball",
		DropID:      Snowball,
		Stackable:   64,
		MaxHealth:   100,
		Category:    Item,
	},
	Stick: {
		DisplayName: "Stick",
		Stackable:   64,
		MaxHealth:   1,
		Category:    Item,
	},

	Stone: {
		DisplayName: "Stone",
		DropID:      Stone,
		Stackable:   64,
		MaxHealth:   100,
		Category:    Block,
		BestTool:    ToolPickaxe,
	},
	StoneAxe: {
		DisplayName: "Stone Axe",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Tool,
	},
	StoneBricks: {
		DisplayName: "Stone Bricks",
		DropID:      StoneBricks,
		Stackable:   64,
		MaxHealth:   100,
		Category:    Block,
		BestTool:    ToolPickaxe,
	},

	StonePickaxe: {
		DisplayName: "Stone Pickaxe",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Tool,
	},
	StoneShovel: {
		DisplayName: "Stone Shovel",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Tool,
	},

	Tnt: {
		DisplayName: "TNT",
		DropID:      Tnt,
		Stackable:   64,
		MaxHealth:   100,
		Category:    Block,
		BestTool:    All,
	},
	Torch: {
		DisplayName: "Torch",
		Stackable:   64,
		MaxHealth:   100,
		Category:    Item | Block,
		BestTool:    All,
	},
	WaterBucket: {
		DisplayName: "Water Bucket",
		Stackable:   1,
		MaxHealth:   100,
		Category:    Item,
	},

	WoodenAxe: {
		DisplayName: "Wooden Axe",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Tool,
	},

	WoodenPickaxe: {
		DisplayName: "Wooden Pickaxe",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Tool,
	},
	WoodenShovel: {
		DisplayName: "Wooden Shovel",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Tool,
	},
}
