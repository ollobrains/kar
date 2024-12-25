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
	Icon8            = make(map[uint16]*ebiten.Image, 0)
	Items20          = make(map[uint16]*ebiten.Image, 0)
	BlockCrackFrames = make(map[uint16][]*ebiten.Image, 0)
	Hotbar           = util.ReadEbImgFS(fs, "assets/img/gui/hotbar.png")
	SelectionBlock   = util.ReadEbImgFS(fs, "assets/img/gui/selection_block.png")
	SelectionBar     = util.ReadEbImgFS(fs, "assets/img/gui/selection_bar.png")
	PlayerAtlas      = util.ReadEbImgFS(fs, "assets/img/player/player.png")
	cracks           = util.ImgFromFS(fs, "assets/img/cracks.png")
	Font             = util.LoadFontFromFS("assets/font/pixelcode.otf", 18, fs)
)

func init() {

	BlockCrackFrames[items.Bedrock] = blockImgs("bedrock.png")
	BlockCrackFrames[items.CoalOre] = blockImgs("coal_ore.png")
	BlockCrackFrames[items.CraftingTable] = blockImgs("crafting_table.png")
	BlockCrackFrames[items.DiamondOre] = blockImgs("diamond_ore.png")
	BlockCrackFrames[items.Dirt] = blockImgs("dirt.png")
	BlockCrackFrames[items.Furnace] = blockImgs("furnace.png")
	BlockCrackFrames[items.FurnaceOn] = blockImgs("furnace_on.png")
	BlockCrackFrames[items.GoldOre] = blockImgs("gold_ore.png")
	BlockCrackFrames[items.GrassBlock] = blockImgs("grass_block.png")
	BlockCrackFrames[items.GrassBlockSnow] = blockImgs("grass_block_snow.png")
	BlockCrackFrames[items.IronOre] = blockImgs("iron_ore.png")
	BlockCrackFrames[items.OakLeaves] = blockImgs("oak_leaves.png")
	BlockCrackFrames[items.OakLog] = blockImgs("oak_log.png")
	BlockCrackFrames[items.OakPlanks] = blockImgs("oak_planks.png")
	BlockCrackFrames[items.OakSapling] = blockImgs("oak_sapling.png")
	BlockCrackFrames[items.Obsidian] = blockImgs("obsidian.png")
	BlockCrackFrames[items.Sand] = blockImgs("sand.png")
	BlockCrackFrames[items.SmoothStone] = blockImgs("smooth_stone.png")
	BlockCrackFrames[items.Snow] = blockImgs("snow.png")
	BlockCrackFrames[items.Stone] = blockImgs("stone.png")
	BlockCrackFrames[items.StoneBricks] = blockImgs("stone_bricks.png")
	BlockCrackFrames[items.Tnt] = blockImgs("tnt.png")
	BlockCrackFrames[items.Torch] = blockImgs("torch.png")

	// blocks
	Icon8[items.Bedrock] = blockIconImg("bedrock.png")
	Icon8[items.CoalOre] = blockIconImg("coal_ore.png")
	Icon8[items.CraftingTable] = blockIconImg("crafting_table.png")
	Icon8[items.DiamondOre] = blockIconImg("diamond_ore.png")
	Icon8[items.Dirt] = blockIconImg("dirt.png")
	Icon8[items.Furnace] = blockIconImg("furnace.png")
	Icon8[items.FurnaceOn] = blockIconImg("furnace_on.png")
	Icon8[items.GoldOre] = blockIconImg("gold_ore.png")
	Icon8[items.GrassBlock] = blockIconImg("grass_block.png")
	Icon8[items.GrassBlockSnow] = blockIconImg("grass_block_snow.png")
	Icon8[items.IronOre] = blockIconImg("iron_ore.png")
	Icon8[items.OakLeaves] = blockIconImg("oak_leaves.png")
	Icon8[items.OakLog] = blockIconImg("oak_log.png")
	Icon8[items.OakPlanks] = blockIconImg("oak_planks.png")
	Icon8[items.OakSapling] = blockIconImg("oak_sapling.png")
	Icon8[items.Obsidian] = blockIconImg("obsidian.png")
	Icon8[items.Sand] = blockIconImg("sand.png")
	Icon8[items.SmoothStone] = blockIconImg("smooth_stone.png")
	Icon8[items.Snow] = blockIconImg("snow.png")
	Icon8[items.Stone] = blockIconImg("stone.png")
	Icon8[items.StoneBricks] = blockIconImg("stone_bricks.png")
	Icon8[items.Tnt] = blockIconImg("tnt.png")
	Icon8[items.Torch] = blockIconImg("torch.png")
	// items
	Icon8[items.Bread] = itemIconImg("bread.png")
	Icon8[items.Bucket] = itemIconImg("bucket.png")
	Icon8[items.Coal] = itemIconImg("coal.png")
	Icon8[items.Diamond] = itemIconImg("diamond.png")
	Icon8[items.DiamondAxe] = itemIconImg("diamond_axe.png")
	Icon8[items.DiamondPickaxe] = itemIconImg("diamond_pickaxe.png")
	Icon8[items.DiamondShovel] = itemIconImg("diamond_shovel.png")
	Icon8[items.GoldIngot] = itemIconImg("gold_ingot.png")
	Icon8[items.IronAxe] = itemIconImg("iron_axe.png")
	Icon8[items.IronIngot] = itemIconImg("iron_ingot.png")
	Icon8[items.IronPickaxe] = itemIconImg("iron_pickaxe.png")
	Icon8[items.IronShovel] = itemIconImg("iron_shovel.png")
	Icon8[items.RawGold] = itemIconImg("raw_gold.png")
	Icon8[items.RawIron] = itemIconImg("raw_iron.png")
	Icon8[items.Snowball] = itemIconImg("snowball.png")
	Icon8[items.Stick] = itemIconImg("stick.png")
	Icon8[items.StoneAxe] = itemIconImg("stone_axe.png")
	Icon8[items.StonePickaxe] = itemIconImg("stone_pickaxe.png")
	Icon8[items.StoneShovel] = itemIconImg("stone_shovel.png")
	Icon8[items.WaterBucket] = itemIconImg("water_bucket.png")
	Icon8[items.WoodenAxe] = itemIconImg("wooden_axe.png")
	Icon8[items.WoodenPickaxe] = itemIconImg("wooden_pickaxe.png")
	Icon8[items.WoodenShovel] = itemIconImg("wooden_shovel.png")

	// blocks

	Items20[items.Bedrock] = blockImg("bedrock.png")
	Items20[items.CoalOre] = blockImg("coal_ore.png")
	Items20[items.CraftingTable] = blockImg("crafting_table.png")
	Items20[items.DiamondOre] = blockImg("diamond_ore.png")
	Items20[items.Dirt] = blockImg("dirt.png")
	Items20[items.Furnace] = blockImg("furnace.png")
	Items20[items.FurnaceOn] = blockImg("furnace_on.png")
	Items20[items.GoldOre] = blockImg("gold_ore.png")
	Items20[items.GrassBlock] = blockImg("grass_block.png")
	Items20[items.GrassBlockSnow] = blockImg("grass_block_snow.png")
	Items20[items.IronOre] = blockImg("iron_ore.png")
	Items20[items.OakLeaves] = blockImg("oak_leaves.png")
	Items20[items.OakLog] = blockImg("oak_log.png")
	Items20[items.OakPlanks] = blockImg("oak_planks.png")
	Items20[items.OakSapling] = blockImg("oak_sapling.png")
	Items20[items.Obsidian] = blockImg("obsidian.png")
	Items20[items.Sand] = blockImg("sand.png")
	Items20[items.SmoothStone] = blockImg("smooth_stone.png")
	Items20[items.Snow] = blockImg("snow.png")
	Items20[items.Stone] = blockImg("stone.png")
	Items20[items.StoneBricks] = blockImg("stone_bricks.png")
	Items20[items.Tnt] = blockImg("tnt.png")
	Items20[items.Torch] = blockImg("torch.png")

	Items20[items.Bread] = itemImg("bread.png")
	Items20[items.Bucket] = itemImg("bucket.png")
	Items20[items.Coal] = itemImg("coal.png")
	Items20[items.Diamond] = itemImg("diamond.png")
	Items20[items.DiamondAxe] = itemImg("diamond_axe.png")
	Items20[items.DiamondPickaxe] = itemImg("diamond_pickaxe.png")
	Items20[items.DiamondShovel] = itemImg("diamond_shovel.png")
	Items20[items.GoldIngot] = itemImg("gold_ingot.png")
	Items20[items.IronAxe] = itemImg("iron_axe.png")
	Items20[items.IronIngot] = itemImg("iron_ingot.png")
	Items20[items.IronPickaxe] = itemImg("iron_pickaxe.png")
	Items20[items.IronShovel] = itemImg("iron_shovel.png")
	Items20[items.RawGold] = itemImg("raw_gold.png")
	Items20[items.RawIron] = itemImg("raw_iron.png")
	Items20[items.Snowball] = itemImg("snowball.png")
	Items20[items.Stick] = itemImg("stick.png")
	Items20[items.StoneAxe] = itemImg("stone_axe.png")
	Items20[items.StonePickaxe] = itemImg("stone_pickaxe.png")
	Items20[items.StoneShovel] = itemImg("stone_shovel.png")
	Items20[items.WaterBucket] = itemImg("water_bucket.png")
	Items20[items.WoodenAxe] = itemImg("wooden_axe.png")
	Items20[items.WoodenPickaxe] = itemImg("wooden_pickaxe.png")
	Items20[items.WoodenShovel] = itemImg("wooden_shovel.png")

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
		x := i * 20
		rec := image.Rect(x, 0, x+20, x+20)
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
