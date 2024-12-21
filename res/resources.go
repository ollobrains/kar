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
	ItemIcons        = make(map[uint16]*ebiten.Image, 0)
	BlockCrackFrames = make(map[uint16][]*ebiten.Image, 0)
	Hotbar           = util.ReadEbImgFS(fs, "assets/img/gui/hotbar.png")
	SelectionBar     = util.ReadEbImgFS(fs, "assets/img/gui/selection_bar.png")
	SelectionBlock   = util.ReadEbImgFS(fs, "assets/img/gui/selection_block.png")
	Font             = util.LoadFontFromFS("assets/font/pixelcode.otf", 18, fs)
	PlayerAtlas      = util.ReadEbImgFS(fs, "assets/img/player/player.png")
	cracks           = util.ImgFromFS(fs, "assets/img/cracks.png")
)

func init() {
	ItemIcons[items.Air] = util.ReadEbImgFS(fs, "assets/img/air.png")

	BlockCrackFrames[items.Bedrock] = blockImgs("bedrock.png")
	BlockCrackFrames[items.CoalBlock] = blockImgs("coal_block.png")
	BlockCrackFrames[items.CoalOre] = blockImgs("coal_ore.png")
	BlockCrackFrames[items.Cobblestone] = blockImgs("cobblestone.png")
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
	ItemIcons[items.Bedrock] = blockIconImg("bedrock.png")
	ItemIcons[items.CoalBlock] = blockIconImg("coal_block.png")
	ItemIcons[items.CoalOre] = blockIconImg("coal_ore.png")
	ItemIcons[items.Cobblestone] = blockIconImg("cobblestone.png")
	ItemIcons[items.CraftingTable] = blockIconImg("crafting_table.png")
	ItemIcons[items.DiamondOre] = blockIconImg("diamond_ore.png")
	ItemIcons[items.Dirt] = blockIconImg("dirt.png")
	ItemIcons[items.Furnace] = blockIconImg("furnace.png")
	ItemIcons[items.FurnaceOn] = blockIconImg("furnace_on.png")
	ItemIcons[items.GoldOre] = blockIconImg("gold_ore.png")
	ItemIcons[items.GrassBlock] = blockIconImg("grass_block.png")
	ItemIcons[items.GrassBlockSnow] = blockIconImg("grass_block_snow.png")
	ItemIcons[items.IronOre] = blockIconImg("iron_ore.png")
	ItemIcons[items.OakLeaves] = blockIconImg("oak_leaves.png")
	ItemIcons[items.OakLog] = blockIconImg("oak_log.png")
	ItemIcons[items.OakPlanks] = blockIconImg("oak_planks.png")
	ItemIcons[items.OakSapling] = blockIconImg("oak_sapling.png")
	ItemIcons[items.Obsidian] = blockIconImg("obsidian.png")
	ItemIcons[items.Sand] = blockIconImg("sand.png")
	ItemIcons[items.SmoothStone] = blockIconImg("smooth_stone.png")
	ItemIcons[items.Snow] = blockIconImg("snow.png")
	ItemIcons[items.Stone] = blockIconImg("stone.png")
	ItemIcons[items.StoneBricks] = blockIconImg("stone_bricks.png")
	ItemIcons[items.Tnt] = blockIconImg("tnt.png")
	ItemIcons[items.Torch] = blockIconImg("torch.png")

	// items
	ItemIcons[items.Bread] = itemIconImg("bread.png")
	ItemIcons[items.Bucket] = itemIconImg("bucket.png")
	ItemIcons[items.Coal] = itemIconImg("coal.png")
	ItemIcons[items.Diamond] = itemIconImg("diamond.png")
	ItemIcons[items.DiamondAxe] = itemIconImg("diamond_axe.png")
	ItemIcons[items.DiamondPickaxe] = itemIconImg("diamond_pickaxe.png")
	ItemIcons[items.DiamondShovel] = itemIconImg("diamond_shovel.png")
	ItemIcons[items.GoldIngot] = itemIconImg("gold_ingot.png")
	ItemIcons[items.IronAxe] = itemIconImg("iron_axe.png")
	ItemIcons[items.IronIngot] = itemIconImg("iron_ingot.png")
	ItemIcons[items.IronPickaxe] = itemIconImg("iron_pickaxe.png")
	ItemIcons[items.IronShovel] = itemIconImg("iron_shovel.png")
	ItemIcons[items.RawGold] = itemIconImg("raw_gold.png")
	ItemIcons[items.RawIron] = itemIconImg("raw_iron.png")
	ItemIcons[items.Snowball] = itemIconImg("snowball.png")
	ItemIcons[items.Stick] = itemIconImg("stick.png")
	ItemIcons[items.StoneAxe] = itemIconImg("stone_axe.png")
	ItemIcons[items.StonePickaxe] = itemIconImg("stone_pickaxe.png")
	ItemIcons[items.StoneShovel] = itemIconImg("stone_shovel.png")
	ItemIcons[items.WaterBucket] = itemIconImg("water_bucket.png")
	ItemIcons[items.WoodenAxe] = itemIconImg("wooden_axe.png")
	ItemIcons[items.WoodenPickaxe] = itemIconImg("wooden_pickaxe.png")
	ItemIcons[items.WoodenShovel] = itemIconImg("wooden_shovel.png")

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
