package items

import (
	"image/color"
	"kar/engine/util"
)

var ItemColorMap = map[uint16]color.RGBA{
	Air:        util.HexToRGBA("#0099ff"),
	GrassBlock: util.HexToRGBA("#00903f"),
	Dirt:       util.HexToRGBA("#74573E"),
	Sand:       util.HexToRGBA("#fff5cc"),
	Stone:      util.HexToRGBA("#949494"),
	CoalOre:    util.HexToRGBA("#372f2f"),
	GoldOre:    util.HexToRGBA("#ffe100"),
	IronOre:    util.HexToRGBA("#b8947d"),
	DiamondOre: util.HexToRGBA("#40efd4"),
}
