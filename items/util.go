package items

import (
	"kar/engine/mathutil"
	"math/rand/v2"
)

var BlockIDs []uint16

func init() {
	for id := range Property {
		if IsBlock(id) {
			BlockIDs = append(BlockIDs, id)
		}
	}
}

func IsBreakable(id uint16) bool {
	return Property[id].Tags&Unbreakable == 0
}

func IsHarvestable(id uint16) bool {
	return Property[id].Tags&Harvestable != 0
}
func IsBlock(id uint16) bool {
	return Property[id].Tags&Block != 0
}
func IsBestTool(blockID, toolID uint16) bool {
	return Property[blockID].BestToolTag&Property[toolID].Tags != 0
}

func RandomBlock() uint16 {
	return BlockIDs[rand.IntN(len(BlockIDs))]
}
func RandomItem() uint16 {
	return uint16(mathutil.RandRangeInt(1, len(Property)-1))
}
