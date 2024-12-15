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
	All = ^uint(0)
)

type ItemProperty struct {
	DisplayName string
	Drops       uint16
	Stackable   uint8
	MaxHealth   float64
	Category    Category
}

var Property = map[uint16]ItemProperty{
	Air: {
		DisplayName: "",
		Drops:       Air,
		Stackable:   0,
		MaxHealth:   1000,
		Category:    None | Unbreakable,
	},
	Arrow: {
		DisplayName: "Arrow",
		Drops:       Arrow,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Weapon | Item | DropItem,
	},

	Bedrock: {
		DisplayName: "Bedrock",
		Drops:       0,
		Stackable:   0,
		MaxHealth:   1000,
		Category:    Block | Unbreakable,
	},
	BeetrootSeeds: {
		DisplayName: "Beetroot Seeds",
		Drops:       0,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Item | Unbreakable | DropItem,
	},
	Bow: {
		DisplayName: "Bow",
		Drops:       0,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Item | Weapon,
	},

	Bread: {
		DisplayName: "Bread",
		Drops:       0,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Item | Food,
	},
	BrewingStand: {
		DisplayName: "Brewing Stand",
		Drops:       0,
		Stackable:   1,
		MaxHealth:   10,
		Category:    Block,
	},
	Bucket: {
		DisplayName: "Bucket",
		Drops:       0,
		Stackable:   1,
		MaxHealth:   10,
		Category:    Item | Tool,
	},
	CartographyTable: {
		DisplayName: "Cartography Table",
		Drops:       0,
		Stackable:   1,
		MaxHealth:   10,
		Category:    Block,
	},
	Charcoal: {
		DisplayName: "Char Coal",
		Stackable:   64,
		MaxHealth:   1,
		Category:    Item | DropItem | RawOre,
	},
	Clay: {
		DisplayName: "Clay",
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block,
	},
	Coal: {
		DisplayName: "Coal",
		Stackable:   64,
		MaxHealth:   1,
		Category:    Item | DropItem | RawOre,
	},
	CoalBlock: {
		DisplayName: "Coal Block",
		Stackable:   64,
		MaxHealth:   1,
		Category:    Block,
	},
	CoalOre: {
		DisplayName: "Coal Ore",
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block | OreBlock,
	},
	CobbledDeepslate: {
		DisplayName: "Cobbled Deepslate",
		Drops:       CobbledDeepslate,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block,
	},
	Cobblestone: {
		DisplayName: "Cobblestone",
		Drops:       Cobblestone,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block,
	},
	CopperIngot: {
		DisplayName: "Copper Ingot",
		Stackable:   64,
		MaxHealth:   1,
		Category:    Item,
	},
	CopperOre: {
		DisplayName: "Copper Ore",
		Drops:       RawCopper,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block | OreBlock,
	},
	CraftingTable: {
		DisplayName: "Crafting Table",
		Drops:       Air,
		Stackable:   1,
		MaxHealth:   10,
		Category:    Block,
	},

	CrossbowStandby: {
		DisplayName: "CrossbowStandby",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Weapon,
	},
	CryingObsidian: {
		DisplayName: "CryingObsidian",
		Stackable:   64,
		MaxHealth:   20,
		Category:    Block,
	},
	Deepslate: {
		DisplayName: "Deepslate Stone",
		Drops:       CobbledDeepslate,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block,
	},
	DeepslateCoalOre: {
		DisplayName: "Deepslate Coal Ore",
		Drops:       Coal,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block | OreBlock,
	},

	DeepslateCopperOre: {
		DisplayName: "Deepslate Copper Ore",
		Drops:       RawCopper,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block | OreBlock,
	},
	DeepslateDiamondOre: {
		DisplayName: "Deepslate Diamond Ore",
		Drops:       Diamond,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block | OreBlock,
	},
	DeepslateEmeraldOre: {
		DisplayName: "Deepslate Emerald Ore",
		Drops:       Emerald,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block | OreBlock,
	},
	DeepslateGoldOre: {
		DisplayName: "Deepslate Gold Ore",
		Drops:       RawGold,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block | OreBlock,
	},
	DeepslateIronOre: {
		DisplayName: "Deepslate Iron Ore",
		Drops:       RawIron,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block | OreBlock,
	},
	DeepslateLapisOre: {
		DisplayName: "Deepslate Lapis Ore",
		Drops:       LapisLazuli,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block | OreBlock,
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
		Category:    Item | Tool,
	},

	DiamondHoe: {
		DisplayName: "Diamond Hoe",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Tool,
	},
	DiamondOre: {
		DisplayName: "Diamond Ore",
		Drops:       Diamond,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block | OreBlock,
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
	DiamondSword: {
		DisplayName: "Diamond Sword",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Weapon,
	},
	Dirt: {
		DisplayName: "Dirt",
		Drops:       Dirt,
		Stackable:   64,
		MaxHealth:   5,
		Category:    Block,
	},
	Emerald: {
		DisplayName: "Emerald",
		Stackable:   64,
		MaxHealth:   1,
		Category:    Item | DropItem | RawOre,
	},
	EmeraldOre: {
		DisplayName: "Emerald Ore",
		Drops:       Emerald,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block | OreBlock,
	},
	EnchantingTable: {
		DisplayName: "Enchanting Table",
		Drops:       EnchantingTable,
		Stackable:   1,
		MaxHealth:   10,
		Category:    Block,
	},
	EndPortalFrame: {
		DisplayName: "EndPortalFrame",
		Drops:       EndPortalFrame,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block,
	},
	FletchingTable: {
		DisplayName: "Fletching Table",
		Drops:       FletchingTable,
		Stackable:   1,
		MaxHealth:   10,
		Category:    Block,
	},
	Furnace: {
		DisplayName: "Furnace",
		Drops:       Furnace,
		Stackable:   1,
		MaxHealth:   10,
		Category:    Block,
	},
	FurnaceOn: {
		DisplayName: "Furnace On",
		Drops:       Furnace,
		Stackable:   1,
		MaxHealth:   10,
		Category:    Block,
	},
	GoldIngot: {
		DisplayName: "Gold Ingot",
		Stackable:   64,
		MaxHealth:   1,
		Category:    Item,
	},
	GoldOre: {
		DisplayName: "Gold Ore",
		Drops:       RawGold,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block | OreBlock,
	},
	GoldenAxe: {
		DisplayName: "Golden Axe",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Tool,
	},
	GoldenHoe: {
		DisplayName: "Golden Hoe",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Tool,
	},
	GoldenPickaxe: {
		DisplayName: "Golden Pickaxe",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Tool,
	},
	GoldenShovel: {
		DisplayName: "Golden Hoe",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Tool,
	},
	GoldenSword: {
		DisplayName: "Golden Sword",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Weapon,
	},
	GrassBlock: {
		DisplayName: "Grass Block",
		Drops:       Dirt,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block,
	},
	GrassBlockSnow: {
		DisplayName: "Grass Block Snow",
		Drops:       Dirt,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block,
	},

	IronAxe: {
		DisplayName: "Iron Axe",
		Stackable:   1,
		MaxHealth:   1,
	},
	IronHoe: {
		DisplayName: "Iron Hoe",
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
		Drops:       RawIron,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block | OreBlock,
	},
	IronPickaxe: {
		DisplayName: "Iron Pickaxe",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Tool,
	},
	IronShovel: {
		DisplayName: "Iron Hoe",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Tool,
	},
	IronSword: {
		DisplayName: "Iron Sword",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Weapon,
	},
	LapisLazuli: {
		DisplayName: "Lapis Lazuli",
		Stackable:   64,
		MaxHealth:   1,
		Category:    Item | DropItem | RawOre,
	},
	LapisOre: {
		DisplayName: "Lapis Ore",
		Drops:       LapisLazuli,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block | OreBlock,
	},
	LavaBucket: {
		DisplayName: "Lava Bucket",
		Drops:       Air,
		Stackable:   1,
		MaxHealth:   10,
		Category:    Item | Tool,
	},
	MelonSeeds: {
		DisplayName: "Melon Seeds",
		Drops:       Air,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Item | Unbreakable | DropItem,
	},
	MilkBucket: {
		DisplayName: "Milk Bucket",
		Drops:       Air,
		Stackable:   1,
		MaxHealth:   10,
		Category:    Item | Tool,
	},
	NetherBricks: {
		DisplayName: "Nether Bricks",
		Drops:       NetherBricks,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block,
	},
	NetherGoldOre: {
		DisplayName: "Nether Gold Ore",
		Drops:       RawGold,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block | OreBlock,
	},
	NetherQuartzOre: {
		DisplayName: "Nether Quartz Ore",
		Drops:       NetherQuartzOre,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block | OreBlock,
	},
	NetheriteAxe: {
		DisplayName: "Netherite Axe",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Tool,
	},
	NetheriteHoe: {
		DisplayName: "Netherite Hoe",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Tool,
	},
	NetheriteIngot: {
		DisplayName: "Netherite Ingot",
		Stackable:   64,
		MaxHealth:   1,
		Category:    Item,
	},
	NetheritePickaxe: {
		DisplayName: "Netherite Pickaxe",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Tool,
	},
	NetheriteScrap: {
		DisplayName: "Netherite Scrap",
		Stackable:   64,
		MaxHealth:   1,
		Category:    Item | RawOre | DropItem,
	},
	NetheriteShovel: {
		DisplayName: "Netherite Shovel",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Tool,
	},
	NetheriteSword: {
		DisplayName: "Netherite Sword",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Weapon,
	},
	Netherrack: {
		DisplayName: "Netherrack",
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block,
	},
	OakLeaves: {
		DisplayName: "Leaves",
		Drops:       OakLeaves,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block,
	},
	OakLog: {
		DisplayName: "Log",
		Drops:       OakLog,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block,
	},
	OakPlanks: {
		DisplayName: "Tree Plank",
		Drops:       OakPlanks,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block,
	},
	OakSapling: {
		DisplayName: "Sapling",
		Drops:       OakSapling,
		Stackable:   64,
		MaxHealth:   1,
		Category:    Block | Item,
	},
	Obsidian: {
		DisplayName: "Obsidian",
		Drops:       Obsidian,
		Stackable:   64,
		MaxHealth:   20,
		Category:    Block,
	},
	PowderSnowBucket: {
		DisplayName: "Powder Snow Bucket",
		Drops:       Air,
		Stackable:   1,
		MaxHealth:   1,
		Category:    Tool | Item,
	},
	PumpkinSeeds: {
		DisplayName: "Pumpkin Seeds",
		Drops:       Air,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Item | Unbreakable | DropItem,
	},
	RawCopper: {
		DisplayName: "Raw Copper",
		Stackable:   64,
		MaxHealth:   1,
		Category:    Item | DropItem | RawOre,
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

	RedNetherBricks: {
		DisplayName: "Red Nether Bricks",
		Drops:       RedNetherBricks,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block,
	},
	RedSand: {
		DisplayName: "Red Sand",
		Drops:       RedSand,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block,
	},
	RedSandstone: {
		DisplayName: "Red Sandtone",
		Drops:       RedSandstone,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block,
	},

	RedstoneTorch: {
		DisplayName: "Redstone Torch",
		Drops:       RedstoneTorch,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Item,
	},
	RedstoneTorchOff: {
		DisplayName: "Redstone Torch Off",
		Drops:       RedstoneTorchOff,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Item,
	},

	RootedDirt: {
		DisplayName: "Rooted Dirt",
		Drops:       RootedDirt,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block,
	},
	Sand: {
		DisplayName: "Sand",
		Drops:       Sand,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block,
	},
	Sandstone: {
		DisplayName: "Sandstone",
		Drops:       Sandstone,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block,
	},
	SmithingTable: {
		DisplayName: "Snow",
		Drops:       Air,
		Stackable:   1,
		MaxHealth:   10,
		Category:    Block,
	},
	SmoothStone: {
		DisplayName: "Smooth Stone",
		Drops:       SmoothStone,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block,
	},
	Snow: {
		DisplayName: "Snow",
		Drops:       Dirt,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block,
	},
	Snowball: {
		DisplayName: "Snowball",
		Drops:       Snowball,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Item,
	},
	SoulSand: {
		DisplayName: "Soul Sand",
		Drops:       SoulSand,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block,
	},

	SoulSoil: {
		DisplayName: "Soul Soil",
		Drops:       SoulSoil,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block,
	},
	SoulTorch: {
		DisplayName: "Soul Torch",
		Drops:       Air,
		Stackable:   64,
		MaxHealth:   10,
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
		Drops:       Cobblestone,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block,
	},
	StoneAxe: {
		DisplayName: "Stone Axe",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Tool,
	},
	StoneBricks: {
		DisplayName: "Stone Bricks",
		Drops:       StoneBricks,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block,
	},
	StoneHoe: {
		DisplayName: "Stone Hoe",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Tool,
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
	StoneSword: {
		DisplayName: "Stone Sword",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Weapon,
	},
	Tnt: {
		DisplayName: "TNT",
		Drops:       Tnt,
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block,
	},
	Torch: {
		DisplayName: "Torch",
		Stackable:   64,
		MaxHealth:   10,
		Category:    Item,
	},
	TorchflowerSeeds: {
		DisplayName: "Torchflower Seeds",
		Stackable:   64,
		MaxHealth:   10,
		Category:    Item | Unbreakable | DropItem,
	},
	WaterBucket: {
		DisplayName: "Water Bucket",
		Stackable:   1,
		MaxHealth:   10,
		Category:    Item,
	},
	Wheat: {
		DisplayName: "Wheat",
		Stackable:   64,
		MaxHealth:   10,
		Category:    Item,
	},
	WheatCrops: {
		DisplayName: "Wheat Crops",
		Stackable:   64,
		MaxHealth:   10,
		Category:    Block | Unbreakable | Harvestable,
	},
	WheatSeeds: {
		DisplayName: "Wheat Seeds",
		Stackable:   64,
		MaxHealth:   10,
		Category:    Item | Unbreakable | DropItem,
	},
	WoodenAxe: {
		DisplayName: "Wooden Axe",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Tool,
	},
	WoodenHoe: {
		DisplayName: "Wooden Hoe",
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
		DisplayName: "Wooden Hoe",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Tool,
	},
	WoodenSword: {
		DisplayName: "Wooden Sword",
		Stackable:   1,
		MaxHealth:   1,
		Category:    Item | Weapon,
	},
}
