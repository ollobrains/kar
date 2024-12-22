// res is resources
package res

import (
	"embed"
	"image"
	"kar/engine/util"
	"kar/items"

	"github.com/anthonynsimon/bild/blend"
	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/*
var fs embed.FS

var (
	Items8             = make(map[uint16]*ebiten.Image, 0)
	Items16            = make(map[uint16]*ebiten.Image, 0)
	BlockCrackFrames16 = make(map[uint16][]*ebiten.Image, 0)
	Hotbar             = util.ReadEbImgFS(fs, "assets/img/gui/hotbar.png")
	SelectionBlock     = util.ReadEbImgFS(fs, "assets/img/gui/selection_block.png")
	SelectionBar       = util.ReadEbImgFS(fs, "assets/img/gui/selection_bar.png")
	PlayerAtlas        = util.ReadEbImgFS(fs, "assets/img/player/player.png")
	cracks             = util.ImgFromFS(fs, "assets/img/cracks.png")
	Font               = util.LoadFontFromFS("assets/font/pixelcode.otf", 18, fs)
)

func init() {

	BlockCrackFrames16[items.Bedrock] = blockImgs("bedrock.png")
	BlockCrackFrames16[items.CoalOre] = blockImgs("coal_ore.png")
	BlockCrackFrames16[items.Cobblestone] = blockImgs("cobblestone.png")
	BlockCrackFrames16[items.CraftingTable] = blockImgs("crafting_table.png")
	BlockCrackFrames16[items.DiamondOre] = blockImgs("diamond_ore.png")
	BlockCrackFrames16[items.Dirt] = blockImgs("dirt.png")
	BlockCrackFrames16[items.Furnace] = blockImgs("furnace.png")
	BlockCrackFrames16[items.FurnaceOn] = blockImgs("furnace_on.png")
	BlockCrackFrames16[items.GoldOre] = blockImgs("gold_ore.png")
	BlockCrackFrames16[items.GrassBlock] = blockImgs("grass_block.png")
	BlockCrackFrames16[items.GrassBlockSnow] = blockImgs("grass_block_snow.png")
	BlockCrackFrames16[items.IronOre] = blockImgs("iron_ore.png")
	BlockCrackFrames16[items.OakLeaves] = blockImgs("oak_leaves.png")
	BlockCrackFrames16[items.OakLog] = blockImgs("oak_log.png")
	BlockCrackFrames16[items.OakPlanks] = blockImgs("oak_planks.png")
	BlockCrackFrames16[items.OakSapling] = blockImgs("oak_sapling.png")
	BlockCrackFrames16[items.Obsidian] = blockImgs("obsidian.png")
	BlockCrackFrames16[items.Sand] = blockImgs("sand.png")
	BlockCrackFrames16[items.SmoothStone] = blockImgs("smooth_stone.png")
	BlockCrackFrames16[items.Snow] = blockImgs("snow.png")
	BlockCrackFrames16[items.Stone] = blockImgs("stone.png")
	BlockCrackFrames16[items.StoneBricks] = blockImgs("stone_bricks.png")
	BlockCrackFrames16[items.Tnt] = blockImgs("tnt.png")
	BlockCrackFrames16[items.Torch] = blockImgs("torch.png")

	// blocks
	Items8[items.Bedrock] = blockIconImg("bedrock.png")
	Items8[items.CoalOre] = blockIconImg("coal_ore.png")
	Items8[items.Cobblestone] = blockIconImg("cobblestone.png")
	Items8[items.CraftingTable] = blockIconImg("crafting_table.png")
	Items8[items.DiamondOre] = blockIconImg("diamond_ore.png")
	Items8[items.Dirt] = blockIconImg("dirt.png")
	Items8[items.Furnace] = blockIconImg("furnace.png")
	Items8[items.FurnaceOn] = blockIconImg("furnace_on.png")
	Items8[items.GoldOre] = blockIconImg("gold_ore.png")
	Items8[items.GrassBlock] = blockIconImg("grass_block.png")
	Items8[items.GrassBlockSnow] = blockIconImg("grass_block_snow.png")
	Items8[items.IronOre] = blockIconImg("iron_ore.png")
	Items8[items.OakLeaves] = blockIconImg("oak_leaves.png")
	Items8[items.OakLog] = blockIconImg("oak_log.png")
	Items8[items.OakPlanks] = blockIconImg("oak_planks.png")
	Items8[items.OakSapling] = blockIconImg("oak_sapling.png")
	Items8[items.Obsidian] = blockIconImg("obsidian.png")
	Items8[items.Sand] = blockIconImg("sand.png")
	Items8[items.SmoothStone] = blockIconImg("smooth_stone.png")
	Items8[items.Snow] = blockIconImg("snow.png")
	Items8[items.Stone] = blockIconImg("stone.png")
	Items8[items.StoneBricks] = blockIconImg("stone_bricks.png")
	Items8[items.Tnt] = blockIconImg("tnt.png")
	Items8[items.Torch] = blockIconImg("torch.png")
	// items
	Items8[items.Bread] = itemIconImg("bread.png")
	Items8[items.Bucket] = itemIconImg("bucket.png")
	Items8[items.Coal] = itemIconImg("coal.png")
	Items8[items.Diamond] = itemIconImg("diamond.png")
	Items8[items.DiamondAxe] = itemIconImg("diamond_axe.png")
	Items8[items.DiamondPickaxe] = itemIconImg("diamond_pickaxe.png")
	Items8[items.DiamondShovel] = itemIconImg("diamond_shovel.png")
	Items8[items.GoldIngot] = itemIconImg("gold_ingot.png")
	Items8[items.IronAxe] = itemIconImg("iron_axe.png")
	Items8[items.IronIngot] = itemIconImg("iron_ingot.png")
	Items8[items.IronPickaxe] = itemIconImg("iron_pickaxe.png")
	Items8[items.IronShovel] = itemIconImg("iron_shovel.png")
	Items8[items.RawGold] = itemIconImg("raw_gold.png")
	Items8[items.RawIron] = itemIconImg("raw_iron.png")
	Items8[items.Snowball] = itemIconImg("snowball.png")
	Items8[items.Stick] = itemIconImg("stick.png")
	Items8[items.StoneAxe] = itemIconImg("stone_axe.png")
	Items8[items.StonePickaxe] = itemIconImg("stone_pickaxe.png")
	Items8[items.StoneShovel] = itemIconImg("stone_shovel.png")
	Items8[items.WaterBucket] = itemIconImg("water_bucket.png")
	Items8[items.WoodenAxe] = itemIconImg("wooden_axe.png")
	Items8[items.WoodenPickaxe] = itemIconImg("wooden_pickaxe.png")
	Items8[items.WoodenShovel] = itemIconImg("wooden_shovel.png")

	// blocks

	Items16[items.Bedrock] = blockImg("bedrock.png")
	Items16[items.Bedrock] = blockImg("bedrock.png")
	Items16[items.CoalOre] = blockImg("coal_ore.png")
	Items16[items.Cobblestone] = blockImg("cobblestone.png")
	Items16[items.CraftingTable] = blockImg("crafting_table.png")
	Items16[items.DiamondOre] = blockImg("diamond_ore.png")
	Items16[items.Dirt] = blockImg("dirt.png")
	Items16[items.Furnace] = blockImg("furnace.png")
	Items16[items.FurnaceOn] = blockImg("furnace_on.png")
	Items16[items.GoldOre] = blockImg("gold_ore.png")
	Items16[items.GrassBlock] = blockImg("grass_block.png")
	Items16[items.GrassBlockSnow] = blockImg("grass_block_snow.png")
	Items16[items.IronOre] = blockImg("iron_ore.png")
	Items16[items.OakLeaves] = blockImg("oak_leaves.png")
	Items16[items.OakLog] = blockImg("oak_log.png")
	Items16[items.OakPlanks] = blockImg("oak_planks.png")
	Items16[items.OakSapling] = blockImg("oak_sapling.png")
	Items16[items.Obsidian] = blockImg("obsidian.png")
	Items16[items.Sand] = blockImg("sand.png")
	Items16[items.SmoothStone] = blockImg("smooth_stone.png")
	Items16[items.Snow] = blockImg("snow.png")
	Items16[items.Stone] = blockImg("stone.png")
	Items16[items.StoneBricks] = blockImg("stone_bricks.png")
	Items16[items.Tnt] = blockImg("tnt.png")
	Items16[items.Torch] = blockImg("torch.png")

	Items16[items.Bread] = itemImg("bread.png")
	Items16[items.Bucket] = itemImg("bucket.png")
	Items16[items.Coal] = itemImg("coal.png")
	Items16[items.Diamond] = itemImg("diamond.png")
	Items16[items.DiamondAxe] = itemImg("diamond_axe.png")
	Items16[items.DiamondPickaxe] = itemImg("diamond_pickaxe.png")
	Items16[items.DiamondShovel] = itemImg("diamond_shovel.png")
	Items16[items.GoldIngot] = itemImg("gold_ingot.png")
	Items16[items.IronAxe] = itemImg("iron_axe.png")
	Items16[items.IronIngot] = itemImg("iron_ingot.png")
	Items16[items.IronPickaxe] = itemImg("iron_pickaxe.png")
	Items16[items.IronShovel] = itemImg("iron_shovel.png")
	Items16[items.RawGold] = itemImg("raw_gold.png")
	Items16[items.RawIron] = itemImg("raw_iron.png")
	Items16[items.Snowball] = itemImg("snowball.png")
	Items16[items.Stick] = itemImg("stick.png")
	Items16[items.StoneAxe] = itemImg("stone_axe.png")
	Items16[items.StonePickaxe] = itemImg("stone_pickaxe.png")
	Items16[items.StoneShovel] = itemImg("stone_shovel.png")
	Items16[items.WaterBucket] = itemImg("water_bucket.png")
	Items16[items.WoodenAxe] = itemImg("wooden_axe.png")
	Items16[items.WoodenPickaxe] = itemImg("wooden_pickaxe.png")
	Items16[items.WoodenShovel] = itemImg("wooden_shovel.png")

}

func toEbiten(st []image.Image) []*ebiten.Image {
	l := make([]*ebiten.Image, 0)
	for _, v := range st {
		l = append(l, ebiten.NewImageFromImage(v))
	}
	return l
}

func makeStages(block, stages image.Image) []image.Image {
	frames := make([]image.Image, 0)
	frames = append(frames, block)
	for i := range 4 {
		x := i * 16
		rec := image.Rect(x, 0, x+16, x+16)
		si := stages.(*image.NRGBA).SubImage(rec)
		frames = append(frames, blend.Normal(block, si))
	}
	return frames
}
func blockImgs(f string) []*ebiten.Image {
	frames := makeStages(util.ImgFromFS(fs, "assets/img/blocks/"+f), cracks)
	return toEbiten(frames)
}
func itemIconImg(f string) *ebiten.Image {
	return util.ReadEbImgFS(fs, "assets/img/items_icon/"+f)
}
func blockIconImg(f string) *ebiten.Image {
	return util.ReadEbImgFS(fs, "assets/img/blocks_icon/"+f)
}
func blockImg(f string) *ebiten.Image {
	return util.ReadEbImgFS(fs, "assets/img/blocks/"+f)
}

func itemImg(f string) *ebiten.Image {
	return util.ReadEbImgFS(fs, "assets/img/items/"+f)
}
